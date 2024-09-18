package main

import (
	"log"
	"net/http"

	common "github.com/Stanley3970/commons"
	"github.com/joho/godotenv"
)

var (
	httpAddr string
)

func init() {
	// Load .env file manually
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found, proceeding with default environment variables.")
	}

	// Set the HTTP address from environment variables or use default
	httpAddr = common.EnvString("HTTP_ADDR", ":8080")
}

func main() {
	mux := http.NewServeMux()
	handler := NewHandler()
	handler.registerRoutes(mux)

	log.Printf("Stating HTTP server as %s", httpAddr)

	if err := http.ListenAndServe(httpAddr, mux); err != nil {
		log.Fatal("Failed to start http server")

	}

}
