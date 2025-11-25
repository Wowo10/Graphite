package web

import (
	"embed"
	"net/http"
)

//go:embed *
var Files embed.FS

func Handler() http.Handler {
	return http.FileServer(http.FS(Files))
}
