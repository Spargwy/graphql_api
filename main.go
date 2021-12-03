package main

import (
	"log"
	"net/http"

	"gql_app/api"
	"gql_app/storage"

	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("load envs error: ", err)
	}
	err = storage.DBConnect()
	if err != nil {
		log.Fatal("db connect error: ", err)
	}
}

func main() {
	api.SetupRoutes()
	log.Fatal(http.ListenAndServe(":8080", nil))
}
