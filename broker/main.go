package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

const resp =  `
{
	"response": "This is a response from broker"
}
`

func main() {

	r := mux.NewRouter()

	r.HandleFunc("/submit", func(w http.ResponseWriter, r *http.Request) {
		log.Println("Received request:", r)
		w.Header().Set("Content-type", "application/json; charset=UTF-8")
		w.Write([]byte(resp))
	}).Methods("POST")

	r.HandleFunc("/page", func(w http.ResponseWriter, r *http.Request) {
		log.Println("Received request:", r)
		http.ServeFile(w, r, "index.html")
	}).Methods("GET")

	fmt.Println("Server listening on port 3000")
	if err := http.ListenAndServe(":3000", r); err != nil {
		log.Fatal("Could not start server")
	}
}
