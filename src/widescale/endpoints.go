package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

type jsonResponse struct {
	Status string      `json:"status"`
	Result *[]Document `json:"result"`
}

// SetEndpoints attaches all the handler functions to router.
func SetEndpoints(router *mux.Router) {
	router.HandleFunc("/index", returnEntireIndex).Methods("GET")
	router.HandleFunc("/search", searchWord).Methods("GET")
}

func returnEntireIndex(w http.ResponseWriter, r *http.Request) {
	var queryResult []Document
	for i := range Documents {
		queryResult = append(queryResult, Documents[i])
	}
	response, err := json.MarshalIndent(jsonResponse{"success", &queryResult}, "", "  ")
	if err != nil {
		log.Panic(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}

func searchWord(w http.ResponseWriter, r *http.Request) {
	q, ok := r.URL.Query()["q"]
	if !ok || len(q[0]) <= 0 {
		log.Println("`q`(query) parameter missing.")
		return
	}
	query := Normalize(q[0])
	queryWords := strings.Split(query, " ")
	result := make(map[int]bool) // using it as a set.
	for _, word := range queryWords {
		if InvertedIndex[word] != nil {
			for _, doc := range InvertedIndex[word] {
				if !result[doc] {
					result[doc] = true
				}
			}
		}
	}
	var queryResult []Document
	for i := range result {
		queryResult = append(queryResult, Documents[i])
	}
	response, err := json.MarshalIndent(jsonResponse{"success", &queryResult}, "", "  ")
	if err != nil {
		log.Panic(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}
