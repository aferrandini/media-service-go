package model

type MediaType struct {
	Id        string 	 `json:"id"`
	Name      string 	 `json:"name"`
	Enabled   bool 		 `json:"enabled"`
	Cdns      []DomainType   `json:"cdns"`
	Domains   []DomainType   `json:"domains"`
	Provinces []ProvinceType `json:"provinces"`
}
