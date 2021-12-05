package main

import (
	"gql_app/graph"
	"gql_app/graph/generated"
	"gql_app/storage"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/joho/godotenv"
)

const defaultPort = "8080"

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Cant load envs: ", err)
	}
	if err = storage.DBConnect(); err != nil {
		log.Fatal("Cant connect to DB: ", err)
	}

}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
