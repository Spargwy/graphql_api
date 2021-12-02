package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"gql_app/storage"

	"github.com/joho/godotenv"
)

func greet(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World! %s", time.Now())
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("load envs error: ", err)
	}
	storage.DBConnect()
	http.HandleFunc("/", greet)
	http.ListenAndServe(":8080", nil)
}
