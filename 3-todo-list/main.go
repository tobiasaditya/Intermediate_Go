package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Todo struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

var Todos []*Todo = []*Todo{}

const baseURL string = "localhost:8080"

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/todos", GetTodos).Methods(http.MethodGet)
	r.HandleFunc("/todo/{id}", GetByID).Methods(http.MethodGet)
	r.HandleFunc("/todo", CreateTodo).Methods(http.MethodPost)
	r.HandleFunc("/todo/{id}", UpdateTodo).Methods(http.MethodPut)
	r.HandleFunc("/todo/{id}", DeleteTodo).Methods(http.MethodDelete)

	log.Println("Listening on " + baseURL)
	http.ListenAndServe(baseURL, r)
}

func GetTodos(w http.ResponseWriter, r *http.Request) {
	data, _ := json.Marshal(Todos)
	w.Header().Set("Content-Type", "application/json")
	w.Write(data)
}

func CreateTodo(w http.ResponseWriter, r *http.Request) {
	var t Todo
	decoded := json.NewDecoder(r.Body)
	decoded.Decode(&t)

	//Append ke todos
	Todos = append(Todos, &t)
	w.Write([]byte("success add todo"))
}

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
	Todos = append(Todos[:idx], Todos[idx+1:]...)

	w.Write([]byte("success delete"))

}

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
