package main

import (
	"graphite/internal/graph"
	"graphite/internal/server"
	"log"
	"os"
	"path/filepath"
)

func main() {
	path := "."
	if len(os.Args) > 1 {
		path = os.Args[1]
	}
	absPath, err := filepath.Abs(path)
	if err != nil {
		log.Fatalf("invalid path: %v", err)
	}
	if err := os.Chdir(absPath); err != nil {
		log.Fatalf("failed to change directory to %s: %v", absPath, err)
	}

	log.Printf("Analyzing Go project at %s\n", absPath)

	g, err := graph.BuildGraph()
	if err != nil {
		log.Fatal(err)
	}

	server.Serve(g)
}
