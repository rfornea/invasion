package models

import (
	"strings"
)

const (
	argDelimiter = " "

	linkDelimiter = "="

	East  = "east"
	West  = "west"
	North = "north"
	South = "south"
)

type City struct {
	Name string
	Links map[string]string
}

var CityMap map[string]City

func init() {
	CityMap = make(map[string]City)
}

func AddCityToMap(cityStr string) {
	cityData := strings.Split(cityStr, argDelimiter)

	if _, ok := CityMap[cityData[0]]; !ok {
		CityMap[cityData[0]] = City {
			Name: cityData[0],
			Links: make(map[string]string),
		}
	}

	for i := 1; i < len(cityData); i++ {
		linkData := strings.Split(cityData[i], linkDelimiter)
		CityMap[cityData[0]].Links[linkData[0]] = linkData[1]
	}
}
