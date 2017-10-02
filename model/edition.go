package model

type EditionType struct {
	Id       string 	`json:"id"`
	Name     string 	`json:"name"`
	Path     string 	`json:"path"`
	Default  bool   	`json:"default"`
}