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

var Todos []*Todo

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
	id := vars["id"]

	//Read Body
	var t Todo
	decoded := json.NewDecoder(r.Body)
	decoded.Decode(&t)

	i, _ := strconv.Atoi(id)

	Todos[i] = &t

	w.Write([]byte("success update"))

}

func DeleteTodo(w http.ResponseWriter, r *http.Request) {
	//Get Path param
	vars := mux.Vars(r)
	id := vars["id"]

	i, _ := strconv.Atoi(id)

	//Search from todos
	for idx, t := range Todos {
		id := t.ID
		if id == i {
			Todos = append(Todos[:idx], Todos[idx+1:]...)
		}
	}

	w.Write([]byte("success delete"))
}
