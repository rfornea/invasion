package maps

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

	// zWSP represents zero-width space.
	zWSP = '\u200B'

	// zWNBSP represents zero-width no-break space.
	zWNBSP = '\uFEFF'

	// zWJ represents zero-width joiner.
	zWJ = '\u200D'

	// zWNJ represents zero-width non-joiner.
	zWNJ = '\u200C'

	empty = ""
)

var replacer = strings.NewReplacer(string(zWSP), empty,
	string(zWNBSP), empty,
	string(zWJ), empty,
	string(zWNJ), empty)

var oppositeDirections map[string]string

/*City is what an individual entry in CityMap looks like*/
type City struct {
	Name   string
	Links  map[string]*City
	aliens map[int]alien
}

/*CityMap is a map of city names to city data*/
var CityMap map[string]City

func init() {
	CityMap = make(map[string]City)
	oppositeDirections = make(map[string]string)

	oppositeDirections[east] = west
	oppositeDirections[west] = east
	oppositeDirections[north] = south
	oppositeDirections[south] = north
}

/*AddCityToMap accepts a string of city data and adds it to the CityMap.
For example, a string like this:

Foo east=Bar

Would be transformed into something like this:

Foo:{Foo map[east:ptrToBar] map[]}

*/
func AddCityToMap(cityStr string) {
	cityStr = removeZeroWidthCharacters(cityStr)
	cityData := strings.Fields(cityStr)
	cityData = recombineCitiesWithSpaces(cityData)
	cityName := strings.Title(cityData[0])

	if _, ok := CityMap[cityName]; !ok {
		CityMap[cityName] = City{
			Name:   cityName,
			Links:  make(map[string]*City),
			aliens: make(map[int]alien),
		}
	}

	for i := 1; i < len(cityData); i++ {
		linkData := strings.Split(cityData[i], linkDelimiter)
		linkName := strings.Title(linkData[1])

		if _, ok := CityMap[linkName]; !ok {
			CityMap[linkName] = City{
				Name:   linkName,
				Links:  make(map[string]*City),
				aliens: make(map[int]alien),
			}
		}

		direction := strings.ToLower(linkData[0])

		city := CityMap[cityName]
		linkCity := CityMap[linkName]

		CityMap[linkName].Links[oppositeDirections[direction]] = &city

		CityMap[cityName].Links[direction] = &linkCity
	}
}

/*PrintInvasionResult will print the remaining world when the simulation is complete*/
func PrintInvasionResult() {
	for _, v := range CityMap {
		fmt.Println(formatPrintStatement(v))
	}
}

func formatPrintStatement(city City) string {
	links := createLinkString(city.Links)
	return city.Name + " " + links
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
	panic("should never get here")
}

// removeZeroWidthCharacters removes all zero-width characters from string s.
func removeZeroWidthCharacters(s string) string {
	return replacer.Replace(s)
}

func recombineCitiesWithSpaces(cityData []string) []string {
	var fixedData []string

	fixedData = append(fixedData, cityData[0])

	for i := 1; i < len(cityData); i++ {
		if !strings.HasPrefix(strings.ToLower(cityData[i]), east+linkDelimiter) &&
			!strings.HasPrefix(strings.ToLower(cityData[i]), west+linkDelimiter) &&
			!strings.HasPrefix(strings.ToLower(cityData[i]), south+linkDelimiter) &&
			!strings.HasPrefix(strings.ToLower(cityData[i]), north+linkDelimiter) {
			fixedData[len(fixedData)-1] = fixedData[len(fixedData)-1] + argDelimiter + cityData[i]
		} else {
			fixedData = append(fixedData, cityData[i])
		}
	}
	return fixedData
}

func (c *City) randLink() *City {
	i := rand.Intn(len(c.Links))
	for k := range c.Links {
		if i == 0 {
			return c.Links[k]
		}
		i--
	}
	panic("should never get here")
}

func (c *City) addAlien(newAlien *alien) {
	c.aliens[newAlien.number] = *newAlien
	if len(c.aliens) == 2 {
		var deadAliens []int
		for k := range c.aliens {
			deadAliens = append(deadAliens, k)
			delete(alienMap, k)
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

func (c *City) removeAlien(rmAlien *alien) {
	delete(c.aliens, rmAlien.number)
}
