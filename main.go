// Copyright 2015 Comcast Cable Communications Management, LLC

// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at

// http://www.apache.org/licenses/LICENSE-2.0

// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	// "./api"
	"./auth"
	"./db"
	routes "./routes"

	"encoding/gob"
	"encoding/json"
	"fmt"
	"github.com/gorilla/handlers"
	"log"
	"net/http"
	"os"
)

type Config struct {
	DbTypeName       string `json:"dbTypeName"`
	DbName           string `json:"dbName"`
	DbUser           string `json:"dbUser"`
	DbPassword       string `json:"dbPassword"`
	ListenerType     string `json:"listenerType,omitempty"`
	ListenerPort     string `json:"listenerPort"`
	ListenerCertFile string `json:"listenerCertFile",omitempty"`
	ListenerKeyFile  string `json:"listenerKeyFile",omitempty"`
}

func main() {
	file, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Println("Error opening config file:", err.Error())
	}
	decoder := json.NewDecoder(file)
	config := Config{}
	err = decoder.Decode(&config)
	if err != nil {
		fmt.Println("Error reading config file:", err.Error())
	}

	gob.Register(auth.SessionUser{}) // this is needed to pass the SessionUser struct around in the gorilla session.

	db.InitializeDatabase(config.DbTypeName, config.DbUser, config.DbPassword, config.DbName)

	var Logger = log.New(os.Stdout, " ", log.Ldate|log.Ltime|log.Lshortfile)
	Logger.Printf("Starting " + config.ListenerType + " server on port " + config.ListenerPort + "...")

	if config.ListenerType == "https" {
		// for https. Make sure you have the server.pem and server.key file. To gen self signed:
		// openssl genrsa -out server.key 2048
		// openssl req -new -x509 -key server.key -out server.pem -days 3650
		err = http.ListenAndServeTLS(":"+config.ListenerPort, config.ListenerCertFile, config.ListenerKeyFile,
			handlers.CombinedLoggingHandler(os.Stdout, routes.CreateRouter()))
	} else {
		err = http.ListenAndServe(":"+config.ListenerPort, handlers.CombinedLoggingHandler(os.Stdout, routes.CreateRouter()))
	}

	if err != nil {
		fmt.Println(err.Error())
	}
}
