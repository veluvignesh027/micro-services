package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/service2", func(w http.ResponseWriter, r *http.Request) {
		log.Println("Received request:", r)
		fmt.Fprint(w, "This response is from service 2")
	}).Methods(http.MethodGet)
	log.Println("Server Listening at port 3002")

	http.ListenAndServe(":3002", r)
}
