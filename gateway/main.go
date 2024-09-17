package main

import (
	"log"
	"net/http"
)

const (
	httpAddr = ":8080"
)

func main() {
	mux := http.NewServeMux()
	handler := NewHandler()
	handler.registerRoutes(mux)

	log.Printf("Stating HTTP server as %s", httpAddr)

	if err := http.ListenAndServe(httpAddr, mux); err != nil {
		log.Fatal("Failed to start http server")

	}

}
