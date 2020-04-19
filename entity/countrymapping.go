package entity

type CountriesMap struct {
	AllCountries map[string]CountryMap
}
type CountryMap struct {
	CountryName string              `json:"name"`
	States      map[string][]string `json:"states"`
}

type CitiesNames struct {
	Cityname []string `json:""`
}
