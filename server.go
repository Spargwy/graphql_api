package main

import (
	"gql_app/graph/generated"
	"gql_app/graph/resolvers"
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
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}
	db, err := resolvers.DBConnect()
	if err != nil {
		log.Fatal("Cant connect to db: ", err)
	}
	schema := generated.NewExecutableSchema(generated.Config{
		Resolvers: &resolvers.Resolver{
			DB: db,
		},
	})
	srv := handler.NewDefaultServer(schema)

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
