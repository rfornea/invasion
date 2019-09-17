package models

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
)

const (
	argDelimiter = " "

	linkDelimiter = "="

	east  = "east"
	west  = "west"
	north = "north"
	south = "south"

	// ZWSP represents zero-width space.
	ZWSP = '\u200B'

	// ZWNBSP represents zero-width no-break space.
	ZWNBSP = '\uFEFF'

	// ZWJ represents zero-width joiner.
	ZWJ = '\u200D'

	// ZWNJ represents zero-width non-joiner.
	ZWNJ = '\u200C'

	empty = ""
)

var replacer = strings.NewReplacer(string(ZWSP), empty,
	string(ZWNBSP), empty,
	string(ZWJ), empty,
	string(ZWNJ), empty)

var oppositeDirections map[string]string

type City struct {
	Name   string
	Links  map[string]*City
	Aliens map[int]Alien
}

var CityMap map[string]City

func init() {
	CityMap = make(map[string]City)
	oppositeDirections = make(map[string]string)

	oppositeDirections[east] = west
	oppositeDirections[west] = east
	oppositeDirections[north] = south
	oppositeDirections[south] = north
}

func AddCityToMap(cityStr string) {
	cityStr = RemoveZeroWidthCharacters(cityStr)
	cityData := strings.Fields(cityStr)
	cityData = recombineCitiesWithSpaces(cityData)

	if _, ok := CityMap[cityData[0]]; !ok {
		CityMap[cityData[0]] = City{
			Name:   cityData[0],
			Links:  make(map[string]*City),
			Aliens: make(map[int]Alien),
		}
	}

	for i := 1; i < len(cityData); i++ {
		linkData := strings.Split(cityData[i], linkDelimiter)

		if _, ok := CityMap[linkData[1]]; !ok {
			CityMap[linkData[1]] = City{
				Name:   linkData[1],
				Links:  make(map[string]*City),
				Aliens: make(map[int]Alien),
			}
		}

		city := CityMap[cityData[0]]
		linkCity := CityMap[linkData[1]]

		CityMap[linkData[1]].Links[oppositeDirections[linkData[0]]] = &city

		CityMap[cityData[0]].Links[linkData[0]] = &linkCity
	}
}

func PrintInvasionResult() {
	for _, v := range CityMap {
		links := createLinkString(v.Links)
		fmt.Println(v.Name + " " + links)
	}
}

func createLinkString(linksMap map[string]*City) string {
	var links []string
	for k, v := range linksMap {
		links = append(links, k+linkDelimiter+v.Name)
	}
	return strings.Join(links, " ")
}

func randCity(m map[string]City) *City {
	i := rand.Intn(len(m))
	for k := range m {
		if i == 0 {
			city := m[k]
			return &city
		}
		i--
	}
	panic("never")
}

func (c *City) randLink() *City {
	i := rand.Intn(len(c.Links))
	for k := range c.Links {
		if i == 0 {
			return c.Links[k]
		}
		i--
	}
	panic("never")
}

func (c *City) addAlien(alien Alien) {
	c.Aliens[alien.Number] = alien
	if len(c.Aliens) == 2 {
		var deadAliens []int
		for k := range c.Aliens {
			deadAliens = append(deadAliens, k)
			delete(AlienMap, k)
		}
		msg := c.Name + " has been destroyed by alien " + strconv.Itoa(deadAliens[0]) +
			" and alien " + strconv.Itoa(deadAliens[1])
		for k, v := range c.Links {
			delete(v.Links, oppositeDirections[k])
		}

		delete(CityMap, c.Name)

		fmt.Println(msg)
	}
}

func (c *City) removeAlien(alien Alien) {
	delete(c.Aliens, alien.Number)
}

// RemoveZeroWidthCharacters removes all zero-width characters from string s.
func RemoveZeroWidthCharacters(s string) string {
	return replacer.Replace(s)
}

func recombineCitiesWithSpaces(cityData []string) []string {
	var fixedData []string

	fixedData = append(fixedData, cityData[0])

	for i := 1; i < len(cityData); i++ {
		if !strings.HasPrefix(cityData[i], east+linkDelimiter) &&
			!strings.HasPrefix(cityData[i], west+linkDelimiter) &&
			!strings.HasPrefix(cityData[i], south+linkDelimiter) &&
			!strings.HasPrefix(cityData[i], north+linkDelimiter) {
			fixedData[len(fixedData)-1] = fixedData[len(fixedData)-1] + argDelimiter + cityData[i]
		} else {
			fixedData = append(fixedData, cityData[i])
		}
	}
	return fixedData
}
