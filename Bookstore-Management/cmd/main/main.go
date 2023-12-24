package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/aman-av/go-bookstore/pkg/routes"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterBookStoreRoutes(r)
	http.Handle("/", r)
	fmt.Printf("Listening on port :9010")
	log.Fatal(http.ListenAndServe("localhost:9010", r))
}
