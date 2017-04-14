package webservice

import (
	"encoding/json"
	"net/http"
	"io/ioutil"
	"medias/model"
	mediasReaderJson "medias/reader/json"
	mediasWriterJson "medias/writer/json"
	"github.com/satori/go.uuid"
)

func CreateMediaHandler(w http.ResponseWriter, r *http.Request) {
	medias := mediasReaderJson.Read()

	var media model.MediaType
	media.Id = uuid.NewV4().String()

	bodyContent, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(bodyContent, &media)

	medias.Medias = append(medias.Medias, media)

	mediasWriterJson.Write(medias)

	body, _ := json.Marshal(MediaResponse{media})
	w.WriteHeader(http.StatusCreated)
	makeResponse(w, r, body)
}
