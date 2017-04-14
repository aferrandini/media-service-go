package main

import (
	"net/http"
	"log"
	"medias/webservice"
	"fmt"
	"os"
)

func main() {
	jsonFilename := os.Getenv("MEDIAS_JSON_FILE")
	if "" == jsonFilename {
		fmt.Print("Please consider declaring MEDIAS_JSON_FILE environment var. Default value './var/medias.json' will be used.\n")
	}

	serverAddress := os.Getenv("MEDIAS_SERVER_ADDRESS")
	if "" == serverAddress {
		serverAddress = ":8080"
	}
	fmt.Println("Starting server at " + serverAddress)


	http.Handle("/", webservice.NewRouter())
	err := http.ListenAndServe(serverAddress, nil)

	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
