package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	_ "3-todo-list/docs"

	"github.com/gorilla/mux"
	httpSwagger "github.com/swaggo/http-swagger" // http-swagger middleware
)

type Todo struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

var Todos []*Todo = []*Todo{}

const baseURL string = "localhost:8080"

// @title Todo Application
// @version 1.0
// @description This is a todo list test management application
// @host localhost:8080
// @BasePath /
func main() {
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
	data, _ := json.Marshal(Todos)
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
	Todos = append(Todos, &t)
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
	idx := getIndexByID(id)

	if idx == -1 {
		w.Write([]byte("data not found"))
	}

	//Update todo at found index
	Todos[idx] = &inputTodo

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
	idx := getIndexByID(id)

	if idx == -1 {
		w.Write([]byte("data not found"))
		return
	}

	//Delete data at found index
	Todos = append(Todos[:idx], Todos[idx+1:]...) //...ellipsis similar to **args in python

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
	idx := getIndexByID(id)

	if idx == -1 {
		w.Write([]byte("data not found"))
		return
	}

	//Get data and marshal
	foundData := Todos[idx]
	data, _ := json.Marshal(foundData)

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
