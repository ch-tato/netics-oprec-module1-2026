package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"
)

var startTime time.Time

type HealthResponse struct {
	Nama      string `json:"nama"`
	NRP       string `json:"nrp"`
	Status    string `json:"status"`
	Timestamp string `json:"timestamp"`
	Uptime    string `json:"uptime"`
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	response := HealthResponse{
		Nama:      "Jokowi",
		NRP:       "5025241036",
		Status:    "UP",
		Timestamp: time.Now().Format(time.RFC3339),
		Uptime:    time.Since(startTime).String(),
	}

	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func main() {
	startTime = time.Now()

	http.HandleFunc("/health", healthHandler)

	port := ":8080"
	log.Printf("API Server is running on port %s...", port)

	if err := http.ListenAndServe(port, nil); err != nil {
		log.Fatal("Failed running the server: ", err)
	}
}
