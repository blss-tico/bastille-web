package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
)

var bastille bastilleModel

func init() {
	log.Println("init")

	bastilleFile, err := os.ReadFile("bastille.json")
	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	err = json.Unmarshal(bastilleFile, &bastille)
	if err != nil {
		log.Fatalf("Error unmarshaling JSON: %v", err)
	}
}

func main() {
	log.Println("main")

	mux := http.NewServeMux()

	handlerTemplates := &HandlersTemplates{}
	handlersData := &HandlersData{}
	routes := &Routes{ht: *handlerTemplates, hd: *handlersData}
	routes.staticRoutes(mux)
	routes.templatesRoutes(mux)
	routes.dataRoutes(mux)

	port, ok := os.LookupEnv("BW_PORT")
	if !ok || port == "" {
		port = IpAddrModel
	}

	log.Printf("Server starting on http://%s", port)
	if err := http.ListenAndServe(port, mux); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
