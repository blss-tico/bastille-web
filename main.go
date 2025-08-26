/*
*  SPDX-License-Identifier: BSD-3-Clause
*
*  Copyright (c) 2025, Bruno Leonardo Tico) <bruno.ccutp@gmail.com>
*  All rights reserved.
*
*  Redistribution and use in source and binary forms, with or without
*  modification, are permitted provided that the following conditions are met:
*
*  * Redistributions of source code must retain the above copyright notice, this
*    list of conditions and the following disclaimer.
*
*  * Redistributions in binary form must reproduce the above copyright notice,
*    this list of conditions and the following disclaimer in the documentation
*    and/or other materials provided with the distribution.
*
*  * Neither the name of the copyright holder nor the names of its
*    contributors may be used to endorse or promote products derived from
*    this software without specific prior written permission.
*
*  THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS"
*  AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE
*  IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE
*  DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT HOLDER OR CONTRIBUTORS BE LIABLE
*  FOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL
*  DAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR
*  SERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER
*  CAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY,
*  OR TORT (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE
*  OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.
 */

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

// @title Bastille-Web
// @version 1.0
// @description API interface to FreeBSD bastille
// @termsOfService http://swagger.io/terms/
// @license.name BSD-3-Clause
// @license.url https://opensource.org/license/bsd-3-clause
// @host addrModel
// @BasePath /
func main() {
	log.Println("main")
	addr := setAddrAndPort(os.Args)
	startHttpServer(addr)
}
