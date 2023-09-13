package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

var (
	Movies []Movie
)

func main() {
	// setting up the routes
	router := mux.NewRouter()
	router.HandleFunc("/movies", GetMovies).Methods("GET")
	router.HandleFunc("/movies/{id}", GetMovie).Methods("GET")
	router.HandleFunc("/movies", postMovie).Methods("POST") // Updated endpoint path
	router.HandleFunc("/movies/{id}", DeleteMovie).Methods("DELETE")
	router.HandleFunc("/movies/{id}", UpdateMovie).Methods("PUT")

	
	// setting up the port number and starting the server at that port number
	listenAddr := ":3000"
	fmt.Printf("Starting the new server at %v....\n", listenAddr)
	log.Fatal(http.ListenAndServe(listenAddr, router))
}
