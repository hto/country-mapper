package main

import (
	"fmt"

	country_mapper "github.com/hto/country-mapper"
)

const (
	defaultFile = "files/country_info.csv"
)

var countryClient *country_mapper.CountryInfoClient

func init() {
	client, err := country_mapper.Load(defaultFile)
	if err != nil {
		panic(err)
	}

	countryClient = client
}

func main() {
	data := countryClient.MapByAlpha2("GB")

	fmt.Println(data)
	fmt.Println(data.Region)
	fmt.Println(data.Subregion)
}
