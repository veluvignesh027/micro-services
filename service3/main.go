package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/service3", func(w http.ResponseWriter, r *http.Request) {
		log.Println("Received request:", r)
		fmt.Fprint(w, "This response is from service 3")
	}).Methods(http.MethodGet)
	http.ListenAndServe(":3000", r)
}
