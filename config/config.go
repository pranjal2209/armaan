package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/pranjal2209/armaan/entity"
)

var Config = make(map[string]entity.CountryMap)

func LoadConfig() {
	file, _ := ioutil.ReadFile("countries_state_cit.json")

	//var data entity.CountriesMap

	err := json.Unmarshal(file, &Config)
	fmt.Println(err)
	//fmt.Println(data)

	fmt.Println(err)
}
