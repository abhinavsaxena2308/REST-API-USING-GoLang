package main

import (
	"log"

	"github.com/abhinavsaxena2308/REST-API-USING-GoLang/cmd/api"
)

func main() {
	//initializr db

	// start api server
	apiServer := api.NewAPIServer(":8080")
	if err := apiServer.Run(); err!= nil{
		log.Fatal("error in running the server")
	}
}