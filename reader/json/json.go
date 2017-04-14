package json

import (
	"io/ioutil"
	"fmt"
	"os"
	"encoding/json"
	"medias/model"
)

func Read() model.MediasType {
	filename := os.Getenv("MEDIAS_JSON_FILENAME")

	if "" == filename {
		filename = "./var/medias.json"
	}

	file, err := ioutil.ReadFile(filename)

	if err != nil {
		fmt.Printf("Read JSON file error '%s': %v\n", filename, err)
		file = []byte("{\"medias\":[]}")
	}

	var medias model.MediasType
	json.Unmarshal(file, &medias)

	return medias
}
