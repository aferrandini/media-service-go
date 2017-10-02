package model

type MediaType struct {
	Id        string 	 		`json:"id"`
	Name      string 	 		`json:"name"`
	Enabled   bool 		 		`json:"enabled"`
	External  bool 		 		`json:"external"`
	Cdns      []DomainType   	`json:"cdns"`
	Domains   []DomainType   	`json:"domains"`
	Editions  []EditionType  	`json:"editions"`
	Provinces []ProvinceType 	`json:"provinces"`
}
