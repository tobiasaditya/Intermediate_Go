package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Todo struct {
	Name string `json:"name"`
}

var Todos []*Todo

const baseURL string = "localhost:8080"

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/todos", GetTodos).Methods(http.MethodGet)
	r.HandleFunc("/todos", CreateTodo).Methods(http.MethodPost)
	r.HandleFunc("/todos/{idx}", UpdateTodo).Methods(http.MethodPut)

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
	idx := vars["idx"]
	fmt.Println("index =" + idx)

	//Read Body
	var t Todo
	decoded := json.NewDecoder(r.Body)
	decoded.Decode(&t)

	i, _ := strconv.Atoi(idx)

	Todos[i] = &t

	w.Write([]byte("success update"))

}
