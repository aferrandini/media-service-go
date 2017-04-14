package webservice

import (
	"net/http"
	"github.com/satori/go.uuid"
)

func makeResponse(w http.ResponseWriter, r *http.Request, body []byte) {
	// Set headers
	w.Header().Add("content-type", "application/json; charset=utf-8")
	w.Header().Add("request-id", uuid.NewV4().String())

	// Set response body
	w.Write(body)
}
