package webservice

import (
	"encoding/json"
	"net/http"
	mediasReaderJson "medias/reader/json"
)

func MediaIdHandler(w http.ResponseWriter, r *http.Request) {
	medias := mediasReaderJson.Read()

	for _, media := range medias.Medias {
		if media.Id == r.URL.Query().Get(":id") {
			body, _ := json.Marshal(MediaResponse{media})
			makeResponse(w, r, body)
			return
		}
	}
}
