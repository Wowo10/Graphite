package main

import (
	"flag"
	"fmt"
	"graphite/internal/graph"
	"graphite/internal/server"
	"log"
	"os"
	"path/filepath"
)

var (
	layoutFlag string
)

func main() {
	flag.StringVar(&layoutFlag, "layout", "cose", "graph layout (dagre, klay, cose, cola)")
	flag.Parse()

	allowed := map[string]bool{
		"dagre": true,
		"klay":  true,
		"cose":  true,
		"cola":  true,
	}

	if !allowed[layoutFlag] {
		fmt.Fprintf(os.Stderr, "Error: unsupported layout '%s'\n", layoutFlag)
		fmt.Fprintf(os.Stderr, "Valid layouts: dagre, klay, cose, cola\n")
		os.Exit(1)
	}

	path := "."
	if flag.NArg() > 0 {
		path = flag.Arg(0)
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

	// TODO: Make this configurable
	g.Layout = layoutFlag

	server.Serve(g)
}
