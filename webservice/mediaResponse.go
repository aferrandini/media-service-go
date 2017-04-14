package webservice

import "medias/model"

type MediaResponse struct {
	Media model.MediaType `json:"media"`
}
