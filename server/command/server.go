package main

import (
	"log"
	"net/http"
	"time"

	"template.github.com/server/repo"

	"template.github.com/server/config"

	"github.com/gorilla/mux"
	"template.github.com/server/api1"
)

func main() {
	dbConfig := &config.DatabaseConfig{}
	dbConfig.Driver = "postgres"
	dbConfig.JdbcUrl = `host=localhost port=5432 sslmode=disable 
	dbname=testgo user=dbapplication_user password=dbapplication_user`
	config.GetConfig().DatabaseConfig = dbConfig

	//initialize repo
	repo.Init()

	rootRoute := mux.NewRouter()
	api1.Init(rootRoute)
	srv := &http.Server{
		Handler: rootRoute,
		Addr:    ":8001",
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	log.Fatal(srv.ListenAndServe())
}
