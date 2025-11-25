package main

import (
	"fmt"
	"graphite/internal/graph"
	"graphite/internal/server"
	"log"
)

func main() {
	g, err := graph.BuildGraph()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%#+v\n", g)

	server.Serve(g)
}
