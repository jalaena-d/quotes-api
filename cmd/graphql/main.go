package main

import (
	"log"
	"net/http"
	"quotes-api/cmd/graphql/graph"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
)

func main() {
	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{
		Resolvers: graph.NewResolver(),
	}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Println("GraphQL server running at http://localhost:8081")
	log.Fatal(http.ListenAndServe(":8081", nil))
}
