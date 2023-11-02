package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type RequestFmt struct {
	Service string `json:"service"`
}
type ResponseFmt struct {
	Response string `json:"response"`
}

func init() {
	log.SetFlags(log.Lshortfile)
}
func main() {

	r := mux.NewRouter()

	r.HandleFunc("/submit", SubmitHandler).Methods("POST")

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Println("Received request:", r)
		http.ServeFile(w, r, "index.html")
	}).Methods("GET")

	fmt.Println("Server listening on port 3000")
	if err := http.ListenAndServe(":3000", r); err != nil {
		log.Fatal("Could not start server")
	}
}

func SubmitHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Received request:", r)
	w.Header().Set("Content-type", "application/json; charset=UTF-8")
	w.Header().Add("Access-Control-Allow-Origin", "*")
	w.Header().Add("Access-Control-Allow-Credentials", "true")

	var temp RequestFmt
	nByte, err := io.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
	}

	err = json.Unmarshal(nByte, &temp)
	if err != nil {
		log.Println(err)
	}

	var resp *http.Response
	switch temp.Service {
	case "service1":
		resp, err = http.Get("http://localhost:3001/service1")
		if err != nil {
			log.Println(err)
		}

	case "service2":
		resp, err = http.Get("http://localhost:3002/service2")
		if err != nil {
			log.Println(err)
		}

	case "service3":
		resp, err = http.Get("http://localhost:3003/service3")
		if err != nil {
			log.Println(err)
		}

	default:
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"response":"Error : Status Bad Request"}`))
		log.Println("Bad Request for ", temp.Service)
		return
	}

	var tempResponse ResponseFmt
	readByte, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
	}
	tempResponse.Response = string(readByte)

	jsonByte, err := json.Marshal(tempResponse)
	if err != nil {
		log.Println(err)
	}

	n, err := w.Write(jsonByte)
	if err != nil {
		log.Println(err)
	}

	log.Println(n, " Bytes read from ", temp.Service, "response.")
}
