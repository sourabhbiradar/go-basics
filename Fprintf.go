package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func gh(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hey")
}

func main() {
	fmt.Println("Hello World")
	r := mux.NewRouter()

	r.HandleFunc("/", gh).Methods("GET")

	log.Fatal(http.ListenAndServe(":8081", r))

}
