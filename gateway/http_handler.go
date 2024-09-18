package main

import (
	"log"
	"net/http"
)

type handler struct {
	//gateway
}

func NewHandler() *handler {
	return &handler{}
}

func (h *handler) registerRoutes(mux *http.ServeMux) {
	// Register the handler function for a specific route
	mux.HandleFunc("/api/customers/", h.HandleCreateOrder)
}

func (h *handler) HandleCreateOrder(w http.ResponseWriter, r *http.Request) {
	// Ensure the request is a POST method
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	//Log for demonstration
	log.Println("Hello")

	w.Write([]byte("Order created"))
}
