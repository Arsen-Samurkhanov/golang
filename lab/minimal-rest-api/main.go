package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type HealthResponse struct {
	Status string `json:"string"`
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(HealthResponse{Status: "ok"})
}

func main() {
	http.HandleFunc("/health", healthHandler)
	log.Println("Listening.....")
	http.ListenAndServe(":8081", nil)
}
