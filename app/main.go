package main

import (
	"job/search/config"
	"job/search/job"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	// separate the code from the 'main' function.
	// all code that available in main function were not testable
	Run()
}

func Run() {
	// prepare gin
	gin.SetMode(gin.ReleaseMode)

	// gin setup
	r := gin.New()
	r.Use(gin.Recovery())

	// prepare postgresql database
	dbPool, _, err := config.NewDBPool(config.DatabaseConfig{
		Username: "user_new",
		Password: "userpass",
		Hostname: "localhost",
		Port:     "5432",
		DBName:   "job",
	})

	// log for error if error occur while connecting to the database
	if err != nil {
		log.Fatalf("unexpected error while tried to connect to database: %v\n", err)
	}

	// prepare meilisearch database
	dbSearch, err := config.NewDBSearchIndex(config.DatabaseSearchConfig{
		Host:   "http://127.0.0.1:7700",
		APIKey: "LQ2DYe1qqAWh9Kvihp7Lar86TTf6GLXuc-fpWbhVzGA",
	})

	// log for error if error occur while connecting to the database
	if err != nil {
		log.Fatalf("unexpected error while tried to connect to database: %v\n", err)
	}

	defer dbPool.Close()

	// setup job api
	database := job.NewDatabase(dbPool)
	databaseSearch := job.NewDatabaseSearch(dbSearch)
	service := job.NewJobService(database, databaseSearch)
	controller := job.NewJobController(service)

  job.Routes(r, controller)

	// run the server
	log.Fatalf("%v", r.Run(":8000"))
}
