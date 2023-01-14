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
	r.HandleFunc("/todos", CreateTodo).Methods(http.MethodPost)
	r.HandleFunc("/todos/{id}", UpdateTodo).Methods(http.MethodPut)
	r.HandleFunc("/todos/{id}", DeleteTodo).Methods(http.MethodDelete)

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
	isFound := false
	for idx, t := range Todos {
		if t.ID == id {
			Todos[idx] = &inputTodo
			isFound = true
			break
		}
	}
	if isFound {
		w.Write([]byte("success update"))
	} else {
		w.WriteHeader(404)
		w.Write([]byte("data not found"))
	}

}

func DeleteTodo(w http.ResponseWriter, r *http.Request) {
	//Get Path param
	vars := mux.Vars(r)

	id, _ := strconv.Atoi(vars["id"])

	//Search from todos
	isFound := false
	for idx, t := range Todos {
		if t.ID == id {
			Todos = append(Todos[:idx], Todos[idx+1:]...)
			isFound = true
			break
		}
	}

	if isFound {
		w.Write([]byte("success delete"))
	} else {
		w.WriteHeader(404)
		w.Write([]byte("data not found"))
	}

}
