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
		fmt.Print("Please consider declaring MEDIAS_JSON_FILE environment var. Default value './var/medias.json' will be used.\n")
		filename = "./var/medias.json"
	}

	file, e := ioutil.ReadFile(filename)

	if e != nil {
		fmt.Printf("Read JSON file error '%s': %v\n", filename, e)
		file = []byte("{\"medias\":[]}")
	}

	var medias model.MediasType
	json.Unmarshal(file, &medias)

	fmt.Println(medias)
	return medias
}
