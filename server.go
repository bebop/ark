package main

import (
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/TimothyStiles/allbase/graph"
	"github.com/TimothyStiles/allbase/parameters"
	"github.com/TimothyStiles/allbase/pkg/retsynth"
)

// RunServerChecks runs all the checks to make sure the server is ready to run.
func RunServerChecks() {
	// Check if the retsynth database connection is valid
	var _, err = retsynth.ConnectDB()
	if err != nil {
		panic(err)
	}
}

// StartGraphQLServer starts the GraphQL server on the port specified in the environment variable PORT or 8080 if not specified.
func StartGraphQLServer() {

	RunServerChecks()

	port, ok := os.LookupEnv("ALLBASE_PORT")
	if !ok {
		port = parameters.DefaultServerPort
	}

	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{}}))

	http.Handle("/", playground.Handler("AllBase GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))

}
