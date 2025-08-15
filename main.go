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

func setAddrAndPort(argsCommandLine []string) string {
    addr := os.Getenv("BW_ADDR")
    log.Println("setAddrAndPort: ", addr)
    if addr == "" {
        if len(argsCommandLine) == 2 {
            log.Println("command line addr: ", argsCommandLine[1])
            addr = argsCommandLine[1]
            addrModel = addr
        } else {
            addrModel = "127.0.0.1:80"
            log.Println("source code addr: ", addrModel)
            addr = addrModel
        }
    } else {
        log.Println("env variable addr: ", addr)
        addrModel = addr
    }

    return addr
}

// @title Bastille-Web
// @version 1.0
// @description API interface to FreeBSD bastille
// @termsOfService http://swagger.io/terms/
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host addrModel
// @BasePath /
func startHttpServer(addr string) {
	log.Println("startHttpServer")

	mux := http.NewServeMux()
	bastille := &Bastille{}
	handlerTemplates := &HandlersTemplates{bl: *bastille}
	handlersData := &HandlersData{bl: *bastille}
	routes := &Routes{ht: *handlerTemplates, hd: *handlersData}
	routes.staticRoutes(mux)
	routes.swaggerRoutes(mux)
	routes.templatesRoutes(mux)
	routes.dataRoutes(mux)

	log.Printf("Server starting on http://%s", addr)
	if err := http.ListenAndServe(addr, mux); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}

func main() {
	log.Println("main")
    addr := setAddrAndPort(os.Args)
	startHttpServer(addr)
}
