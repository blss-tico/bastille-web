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

func LoggingMiddleware(f http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		sessionId := r.Header.Get("X-Request-ID")
		start := time.Now()
		log.Printf("%s %s %s %v %s", r.Method, r.URL.Path, r.RemoteAddr, time.Since(start), sessionId)
		f(w, r)
	}
}

func main() {
	log.Println("main")
	mux := http.NewServeMux()

	// template handlers
	mux.HandleFunc("/", homeHandlerTplt)
	mux.HandleFunc("/help", helpHandlerTplt)
	mux.HandleFunc("/contact", contactHandlerTplt)
	mux.HandleFunc("/bootstrap", bootstrapHandlerTplt)
	mux.HandleFunc("/clone", cloneHandlerTplt)
	mux.HandleFunc("/cmd", cmdHandlerTplt)
	mux.HandleFunc("/config", configHandlerTplt)
	mux.HandleFunc("/convert", convertHandlerTplt)
	mux.HandleFunc("/cp", cpHandlerTplt)
	mux.HandleFunc("/create", createHandlerTplt)
	mux.HandleFunc("/destroy", destroyHandlerTplt)
	mux.HandleFunc("/etcupdate", etcupdateHandlerTplt)
	mux.HandleFunc("/export", exportHandlerTplt)
	mux.HandleFunc("/htop", htopHandlerTplt)
	mux.HandleFunc("/import", importHandlerTplt)
	mux.HandleFunc("/jcp", jcpHandlerTplt)

	// data handlers
	mux.HandleFunc("POST /bootstrap", LoggingMiddleware(bootstrapHandler))
	mux.HandleFunc("POST /clone", LoggingMiddleware(cloneHandler))
	mux.HandleFunc("POST /cmd", LoggingMiddleware(cmdHandler))
	mux.HandleFunc("POST /config", LoggingMiddleware(configHandler))
	mux.HandleFunc("POST /convert", LoggingMiddleware(convertHandler))
	mux.HandleFunc("POST /cp", LoggingMiddleware(cpHandler))
	mux.HandleFunc("POST /create", LoggingMiddleware(createHandler))
	mux.HandleFunc("POST /destroy", LoggingMiddleware(destroyHandler))
	mux.HandleFunc("POST /etcupdate", LoggingMiddleware(etcupdateHandler))
	mux.HandleFunc("POST /export", LoggingMiddleware(exportHandler))
	mux.HandleFunc("POST /htop", LoggingMiddleware(htopHandler))
	mux.HandleFunc("POST /import", LoggingMiddleware(importHandler))
	mux.HandleFunc("POST /jcp", LoggingMiddleware(jcpHandler))

	port, ok := os.LookupEnv("BWU_PORT")
	if !ok || port == "" {
		port = ":8088"
	}

	log.Printf("Server starting on http://localhost%s", port)
	if err := http.ListenAndServe(port, mux); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
