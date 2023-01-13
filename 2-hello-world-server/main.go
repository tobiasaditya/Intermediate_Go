package main

import (
	"log"
	"net/http"
)

const baseUrl string = "localhost:8000"

func main() {
	http.HandleFunc("/hello", PrintHello)
	log.Println("Listening on " + baseUrl)
	err := http.ListenAndServe(baseUrl, nil)
	if err != nil {
		log.Fatal(err)
	}
}

func PrintHello(w http.ResponseWriter, r *http.Request) {
	//Set method request
	if r.Method == http.MethodGet {
		_, err := w.Write([]byte("hello world"))
		if err != nil {
			log.Fatal()
		}
	}

}
