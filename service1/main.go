package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/service1", func(w http.ResponseWriter, r *http.Request) {
		log.Println("Received request:", r)
		fmt.Fprint(w, "This response is from service 1")
	}).Methods(http.MethodGet)
	log.Println("Server Listening at port 3001")

	if err := http.ListenAndServe(":3001", r); err != nil {
		log.Println(err)
	}
}
