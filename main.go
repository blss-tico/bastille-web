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

}
