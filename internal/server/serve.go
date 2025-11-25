package server

import (
	"encoding/json"
	"graphite/internal/graph"
	"log"
	"net/http"
)

func Serve(graph *graph.Graph) error {
	http.Handle("/", http.FileServer(http.Dir("web")))
	http.HandleFunc("/graph.json", func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(graph)
	})

	log.Println("Serving graph at http://localhost:6969")
	return http.ListenAndServe(":6969", nil)
}
