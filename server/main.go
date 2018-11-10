package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"template.github.com/server/api1"
)

func main() {
	rootRoute := mux.NewRouter()
	api1.Init(rootRoute)
	srv := &http.Server{
		Handler: rootRoute,
		Addr:    ":8000",
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	log.Fatal(srv.ListenAndServe())
}
