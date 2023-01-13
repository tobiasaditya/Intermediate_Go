package main

import (
	"encoding/json"
	"log"
	"net/http"

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
	log.Println("Listening on " + baseURL)
	http.ListenAndServe(baseURL, r)
}

func GetTodos(w http.ResponseWriter, r *http.Request) {
	data, _ := json.Marshal(Todos)
	w.Header().Set("Content-Type", "application/json")
	w.Write(data)
}
