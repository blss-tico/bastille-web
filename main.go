package main

import (
	"encoding/json"
	"html/template"
	"log"
	"net/http"
	"os"
	"time"
)

var templates *template.Template
var bastille bastilleModel

func init() {
	log.Println("init")

	bastilleFile, err := os.ReadFile("bastille.json")
	if err != nil {
		log.Fatalf("Error reading file: %v", err)
	}

	err = json.Unmarshal(bastilleFile, &bastille)
	if err != nil {
		log.Fatalf("Error unmarshaling JSON: %v", err)
	}
}

func main() {
	log.Println("main")
	mux := http.NewServeMux()

	// data handlers
	mux.HandleFunc("POST /bootstrap", LoggingMiddleware(bootstrapHandler))

	port, ok := os.LookupEnv("BWU_PORT")
	if !ok || port == "" {
		port = ":8088"
	}

	log.Printf("Server starting on http://localhost%s", port)
	if err := http.ListenAndServe(port, mux); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
