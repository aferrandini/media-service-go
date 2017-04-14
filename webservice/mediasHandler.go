package webservice

import (
	"encoding/json"
	"net/http"
	mediasReaderJson "medias/reader/json"
)

func MediasHandler(w http.ResponseWriter, r *http.Request) {
	medias := mediasReaderJson.Read()

	body, _ := json.Marshal(medias)
	makeResponse(w, r, body)
}
