# Microservice to handle medias in Golang

This is a POC of a microservice to handle medias in Golang

You have to set two environment vars to run the service.

- `MEDIAS_SERVER_ADDRESS` should contain the ip and port you want to run the service `127.0.0.1:8080`. The defaut value is `:8080`.
- `MEDIAS_JSON_FILENAME` should contain the path to the JSON file which contains the medias. The default value is `./var/medias.json`.

