package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	err := CreateIndex("wikidump")
	if err != nil {
		panic(err)
	}
	router := mux.NewRouter()
	SetEndpoints(router)
	log.Fatal(http.ListenAndServe(":8000", router))
}
