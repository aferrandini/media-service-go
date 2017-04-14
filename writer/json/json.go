package json

import (
	"medias/model"
	"os"
	"fmt"
	"io/ioutil"
	"encoding/json"
)

func Write(medias model.MediasType) {
	filename := os.Getenv("MEDIAS_JSON_FILE")

	if "" == filename {
		fmt.Print("Please consider declaring MEDIAS_JSON_FILE environment var. Default value './var/medias.json' will be used.\n")
		filename = "./var/medias.json"
	}

	data, _ := json.Marshal(medias)
	error := ioutil.WriteFile(filename, data, 0644)

	if error != nil {
		panic(error)
	}
}
