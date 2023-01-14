package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"strconv"

	_ "3-todo-list/docs"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	httpSwagger "github.com/swaggo/http-swagger" // http-swagger middleware
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Todo struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

var Todos []*Todo = []*Todo{}
var coll *mongo.Collection

const baseURL string = "localhost:8080"

// @title Todo Application
// @version 1.0
// @description This is a todo list test management application
// @host localhost:8080
// @BasePath /
func main() {

	//Set up database
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}
	uri := os.Getenv("MONGODB_URI")
	if uri == "" {
		log.Fatal("You must set your 'MONGODB_URI' environmental variable. See\n\t https://www.mongodb.com/docs/drivers/go/current/usage-examples/#environment-variable")
	}
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
	}
	defer func() {
		if err := client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()
	log.Println("Success connected to MongoDB")
	coll = client.Database("go_intermediate").Collection("todo")

	r := mux.NewRouter()

	r.HandleFunc("/todos", GetTodos).Methods(http.MethodGet)
	r.HandleFunc("/todo/{id}", GetByID).Methods(http.MethodGet)
	r.HandleFunc("/todo", CreateTodo).Methods(http.MethodPost)
	r.HandleFunc("/todo/{id}", UpdateTodo).Methods(http.MethodPut)
	r.HandleFunc("/todo/{id}", DeleteTodo).Methods(http.MethodDelete)
	r.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)
	log.Println("Listening on " + baseURL)
	http.ListenAndServe(baseURL, r)
}

// GetTodos is a handler for get todos data
// @Summary Get all todos
// @Description get all todos
// @Tags Todo
// @Accept  json
// @Produce  json
// @Success 200 {array} string
// @Router /todos [get]
func GetTodos(w http.ResponseWriter, r *http.Request) {
	//Get to database
	var todos []Todo
	cursor, err := coll.Find(context.TODO(), bson.D{})
	if err != nil {
		panic(err)
	}
	err = cursor.All(context.TODO(), &todos)
	if err != nil {
		panic(err)
	}

	data, _ := json.Marshal(todos)
	w.Header().Set("Content-Type", "application/json")
	w.Write(data)
}

// Create is a handler for create todos API
// @Summary Create new todos
// @Description Create New Todo by inserting ID and name
// @Tags Todo
// @Param todo    body  Todo  true  "Insert new todo"
// @Accept  json
// @Produce  json
// @Success 200 {array} string
// @Router /todo [post]
func CreateTodo(w http.ResponseWriter, r *http.Request) {
	var t Todo
	decoded := json.NewDecoder(r.Body)
	decoded.Decode(&t)

	//Append ke todos
	//Save to DB
	result, err := coll.InsertOne(context.TODO(), t)
	if err != nil {
		panic(err)
	}
	log.Println(result)

	// Todos = append(Todos, &t)
	w.Write([]byte("success add todo"))
}

// UpdateTodo is a handler for updateing todo
// @Summary Update Todo
// @Description Update Todo
// @Tags Todo
// @Param todo    body  Todo  true  "Update new todo"
// @Param id    path  string  true  "todo id"
// @Accept  json
// @Produce  json
// @Success 200 {array} string
// @Router /todo/{id} [put]
func UpdateTodo(w http.ResponseWriter, r *http.Request) {

	//Get Path param
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	//Read Body
	var inputTodo Todo
	decoded := json.NewDecoder(r.Body)
	decoded.Decode(&inputTodo)

	//Search from todos
	updateByID(id, inputTodo)

	w.Write([]byte("success update"))

}

// DeleteTodo is a handler for deleting todo
// @Summary Delete Todo
// @Description Delete Todo by ID
// @Tags Todo
// @Param id  path  string  true  "todo id"
// @Accept  json
// @Produce  json
// @Success 200 {array} string
// @Router /todo/{id} [delete]
func DeleteTodo(w http.ResponseWriter, r *http.Request) {
	//Get Path param
	vars := mux.Vars(r)

	id, _ := strconv.Atoi(vars["id"])

	//Search from todos
	deleteByID(id)
	w.Write([]byte("success delete"))

}

// GetByID is a handler for getting a single todo by ID
// @Summary Get Todo By ID
// @Description Get Todo by ID
// @Tags Todo
// @Param id  path  string  true  "todo id"
// @Accept  json
// @Produce  json
// @Success 200 {array} string
// @Router /todo/{id} [get]
func GetByID(w http.ResponseWriter, r *http.Request) {
	//Get Path param
	vars := mux.Vars(r)

	id, _ := strconv.Atoi(vars["id"])

	//Search from todos
	foundTodo, err := getByID(id)

	if err != nil {
		w.Write([]byte("data not found"))
		return
	}

	//Get data and marshal
	data, _ := json.Marshal(foundTodo)

	w.Header().Set("Content-Type", "application/json")
	w.Write(data)
}

func getIndexByID(id int) int {
	for idx, t := range Todos {
		if t.ID == id {
			return idx
		}
	}
	return -1
}

func getByID(id int) (Todo, error) {
	var t Todo
	filter := bson.D{{Key: "id", Value: id}}
	err := coll.FindOne(context.TODO(), filter).Decode(&t)
	if err != nil {
		return t, err
	}

	return t, nil
}

func updateByID(id int, todo Todo) {

	filter := bson.D{{Key: "id", Value: id}}
	update := bson.D{{Key: "$set", Value: bson.D{{Key: "name", Value: todo.Name}, {Key: "id", Value: todo.ID}}}}
	_, err := coll.UpdateOne(context.TODO(), filter, update)

	if err != nil {
		log.Fatalln(err)
	}

}

func deleteByID(id int) {
	filter := bson.D{{Key: "id", Value: id}}
	_, err := coll.DeleteOne(context.TODO(), filter)

	if err != nil {
		log.Fatalln(err)
	}

}
