package json

import (
	"medias/model"
	"os"
	"io/ioutil"
	"encoding/json"
)

func Write(medias model.MediasType) {
	filename := os.Getenv("MEDIAS_JSON_FILE")

	if "" == filename {
		filename = "./var/medias.json"
	}

	data, _ := json.Marshal(medias)
	err := ioutil.WriteFile(filename, data, 0644)

	if err != nil {
		panic(err)
	}
}
