package main

import (
	"fmt"
	"log"

	"github.com/lectio/lectiod/server"
	observe "github.com/shah/observe-go"
)

func main() {
	observatory := observe.MakeObservatoryFromEnv()
	defer observatory.Close()

	graphQLHTTPServer := server.CreateGraphQLOverHTTPServer(observatory)

	fmt.Printf("Listening on %s", graphQLHTTPServer.Addr)
	log.Fatal(graphQLHTTPServer.ListenAndServe())
}
