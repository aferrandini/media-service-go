package webservice

import (
	"net/http"
	mediasReaderJson "medias/reader/json"
	"encoding/json"
)

func MediaDomainHandler(w http.ResponseWriter, r *http.Request) {
	medias := mediasReaderJson.Read()

	for _, media := range medias.Medias {
		for _, domain := range media.Domains {
			if domain.Domain == r.URL.Query().Get(":domain") {
				body, _ := json.MarshalIndent(MediaResponse{media}, "", "  ")
				makeResponse(w, r, body)
				return
			}
		}
	}
}
