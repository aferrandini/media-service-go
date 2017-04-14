package main

import (
	"net/http"
	"log"
	"medias/webservice"
	"fmt"
	"os"
)

func main() {
	address := os.Getenv("MEDIAS_SERVER_ADDRESS")
	if "" == address {
		address = ":8080"
	}

	http.Handle("/", webservice.NewRouter())
	fmt.Println("Starting server at " + address)
	err := http.ListenAndServe(address, nil)

	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
