package main

import (
	"gql_app/auth"
	"gql_app/graph/generated"
	"gql_app/graph/resolvers"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/go-chi/chi"
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
	resolver := &resolvers.Resolver{}

	err := resolver.DBConnect()
	if err != nil {
		log.Fatal("Cant connect to db: ", err)
	}
	router := chi.NewRouter()

	router.Use(auth.Middleware(resolver.Psql))

	schema := generated.NewExecutableSchema(generated.Config{
		Resolvers: resolver,
	})
	srv := handler.NewDefaultServer(schema)

	router.Handle("/", playground.Handler("GraphQL playground", "/query"))
	router.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
