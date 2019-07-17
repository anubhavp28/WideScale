package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Insufficient or invalid argument(s) provided.")
		return
	}

	// build the inverted index.
	err := CreateIndex(os.Args[1])
	if err != nil {
		panic(err)
	}

	// setup server
	router := mux.NewRouter()
	SetEndpoints(router)
	log.Println("Starting the widescale server.")
	log.Fatal(http.ListenAndServe(":8000", router))
}
