package maps

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_createAliens(t *testing.T) {
	deleteAllMapEntries()

	numAliens := 10

	createAliens(numAliens)

	assert.Equal(t, numAliens, len(alienMap))
}

func Test_assignAliensToCities(t *testing.T) {
	deleteAllMapEntries()

	numAliens := 3

	createAliens(numAliens)

	AddCityToMap("Foo north=Bar west=Baz south=Qu-ux")
	AddCityToMap("Bar south=Foo")
	AddCityToMap("Baz east=Foo")
	AddCityToMap("Qu-ux north=Foo")

	assignAliensToCities()

	for _, v := range alienMap {
		assert.NotNil(t, v.currentCity)
		assert.NotEqual(t, "", v.currentCity.Name)
		assert.NotEqual(t, 0, len(v.currentCity.aliens))
	}
}

func Test_MoveAllAliens(t *testing.T) {
	deleteAllMapEntries()

	numAliens := 3

	createAliens(numAliens)

	AddCityToMap("Bar south=Foo")
	AddCityToMap("Foo north=Bar south=Qu-ux")
	AddCityToMap("Qu-ux north=Foo south=Bee")
	AddCityToMap("Bee north=Qu-ux south=York")
	AddCityToMap("York north=Bee south=St. James")
	AddCityToMap("St. James north=York south=Baz")
	AddCityToMap("Baz north=St. James south=Burg")
	AddCityToMap("Burg north=Baz south=Crystal Shores")
	AddCityToMap("Crystal Shores north=Burg")

	tmpAlien := alienMap[1]
	city1 := CityMap["Bar"]
	tmpAlien.currentCity = &city1
	alienMap[1] = tmpAlien
	cityPtr1 := &city1
	cityPtr1.addAlien(&tmpAlien)

	tmpAlien = alienMap[2]
	city2 := CityMap["York"]
	tmpAlien.currentCity = &city2
	alienMap[2] = tmpAlien
	cityPtr2 := &city2
	cityPtr2.addAlien(&tmpAlien)

	tmpAlien = alienMap[3]
	city3 := CityMap["Crystal Shores"]
	tmpAlien.currentCity = &city3
	alienMap[3] = tmpAlien
	cityPtr3 := &city3
	cityPtr3.addAlien(&tmpAlien)

	copyAlienMap := make(map[int]alien)
	for k, v := range alienMap {
		copyAlienMap[k] = v
	}

	MoveAllAliens()

	for k, v := range alienMap {
		assert.NotEqual(t, copyAlienMap[k].currentCity.Name, v.currentCity.Name)
	}
}

func Test_move(t *testing.T) {
	deleteAllMapEntries()

	numAliens := 1

	createAliens(numAliens)

	AddCityToMap("Bar south=Foo")
	AddCityToMap("Foo north=Bar")

	assignAliensToCities()

	startCityName := alienMap[1].currentCity.Name

	tmpAlien := alienMap[1]
	tmpPtr := &tmpAlien

	tmpPtr.move()

	endCityName := alienMap[1].currentCity.Name

	assert.NotEqual(t, startCityName, endCityName)
	assert.Equal(t, 1, len(CityMap[endCityName].aliens))
	assert.Equal(t, 0, len(CityMap[startCityName].aliens))
}
