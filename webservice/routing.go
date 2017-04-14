package webservice

import (
	"github.com/bmizerany/pat"
	"net/http"
)

func NewRouter() *pat.PatternServeMux {
	router := pat.New()
	router.Get("/medias", http.HandlerFunc(MediasHandler))
	router.Get("/medias/:id", http.HandlerFunc(MediaIdHandler))
	router.Get("/medias/domain/:domain", http.HandlerFunc(MediaDomainHandler))

	router.Post("/medias", http.HandlerFunc(CreateMediaHandler))
	return router
}
