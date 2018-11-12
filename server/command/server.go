package main

import (
	"log"
	"net/http"
	"time"

	"template.github.com/server/config"

	"github.com/gorilla/mux"
	"template.github.com/server/api1"
)

func main() {
	dbConfig := &config.DatabaseConfig{}
	dbConfig.Driver = "sqllite3"
	dbConfig.JdbcUrl = "./template.db"
	config.GetConfig().DatabaseConfig = dbConfig

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
