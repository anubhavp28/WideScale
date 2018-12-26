package main

import "net/http"
import "github.com/gorilla/mux"

func SetEndpoints(router *mux.Router) {
	router.HandleFunc("/index", returnEntireIndex).Methods("GET")
	router.HandleFunc("/search/{word}", searchWord).Methods("GET")
}

func returnEntireIndex(w http.ResponseWriter, r *http.Request) {}

func searchWord(w http.ResponseWriter, r *http.Request) {}
