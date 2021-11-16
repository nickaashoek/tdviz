package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

const PORT = 8000

func HandleFormula(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Got a request for handling a formula")
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/formula", HandleFormula).
		Methods("POST")

	go http.ListenAndServe(fmt.Sprintf(":%v", PORT), r)

	fmt.Printf("[TD-BACK] Running server on localhost:%d\n", PORT)

	select {}
}
