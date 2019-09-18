package maps

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func deleteAllMapEntries() {
	for k := range CityMap {
		delete(CityMap, k)
	}

	for k := range alienMap {
		delete(alienMap, k)
	}
}

func Test_AddCityToMap(t *testing.T) {
	deleteAllMapEntries()

	testStr := "Foo north=Bar"

	AddCityToMap(testStr)

	foo := "Foo"
	north := "north"
	bar := "Bar"
	south := "south"

	errMsg := "city map not built as expected"

	if _, ok := CityMap[foo]; !ok {
		t.Fatal(errMsg)
	}

	if _, ok := CityMap[bar]; !ok {
		t.Fatal(errMsg)
	}

	if _, ok := CityMap[foo].Links[north]; !ok {
		t.Fatal(errMsg)
	}

	if _, ok := CityMap[bar].Links[south]; !ok {
		t.Fatal(errMsg)
	}

	if CityMap[foo].Links[north].Name != bar {
		t.Fatal(errMsg)
	}

	if CityMap[bar].Links[south].Name != foo {
		t.Fatal(errMsg)
	}
}

func Test_formatPrintStatement(t *testing.T) {
	deleteAllMapEntries()

	testStr := "Foo north=Bar"

	AddCityToMap(testStr)

	result := formatPrintStatement(CityMap["Foo"])

	assert.Equal(t, testStr, result)
}

func Test_createLinkString(t *testing.T) {
	deleteAllMapEntries()

	testStr := "Foo north=Bar west=Baz"
	expectedResult1 := "north=Bar west=Baz"
	expectedResult2 := "west=Baz north=Bar"

	AddCityToMap(testStr)

	result := createLinkString(CityMap["Foo"].Links)

	assert.True(t, result == expectedResult1 || result == expectedResult2)
}

func Test_randCity(t *testing.T) {
	deleteAllMapEntries()

	testStr := "Foo north=Bar"

	AddCityToMap(testStr)

	result := randCity(CityMap)

	assert.True(t, result.Name == CityMap["Foo"].Name || result.Name == CityMap["Bar"].Name)
}

func Test_randLink(t *testing.T) {
	deleteAllMapEntries()

	testStr := "Foo north=Bar"

	AddCityToMap(testStr)

	city := CityMap["Foo"]

	result := city.randLink()

	assert.True(t, result.Name == CityMap["Bar"].Name)
}

func Test_addAlien(t *testing.T) {
	deleteAllMapEntries()

	testStr := "Foo north=Bar"

	AddCityToMap(testStr)

	newAlien1 := alien{
		number:      1,
		currentCity: nil,
	}
	newAlien2 := alien{
		number:      2,
		currentCity: nil,
	}

	foo := CityMap["Foo"]
	fooPtr := &foo

	fooPtr.addAlien(&newAlien1)

	// expect that we have successfully added 1 alien to Foo
	assert.Equal(t, 1, len(CityMap["Foo"].aliens))

	fooPtr.addAlien(&newAlien2)

	// expect that Foo is gone
	if _, ok := CityMap["Foo"]; ok {
		t.Fatal("Foo should be gone")
	}

	// expect that alien 1 is gone
	if _, ok := alienMap[1]; ok {
		t.Fatal("Alien 1 should be gone")
	}

	// expect that alien 2 is gone
	if _, ok := alienMap[2]; ok {
		t.Fatal("Alien 2 should be gone")
	}

	// should be only 1 city left in city map
	assert.Equal(t, 1, len(CityMap))

	// Bar should still exist
	if _, ok := CityMap["Bar"]; !ok {
		t.Fatal("Bar should still exist")
	}

	// Bar should have no links left
	assert.Equal(t, 0, len(CityMap["Bar"].Links))
}

func Test_removeAlien(t *testing.T) {
	deleteAllMapEntries()

	testStr := "Foo north=Bar"

	AddCityToMap(testStr)

	newAlien1 := alien{
		number:      1,
		currentCity: nil,
	}

	foo := CityMap["Foo"]
	fooPtr := &foo

	fooPtr.addAlien(&newAlien1)

	// 1 alien in Foo
	assert.Equal(t, 1, len(CityMap["Foo"].aliens))

	fooPtr.removeAlien(&newAlien1)

	// no aliens in Foo
	assert.Equal(t, 0, len(CityMap["Foo"].aliens))
}

func Test_removeZeroWidthCharacters(t *testing.T) {
	deleteAllMapEntries()

	startStr := "Foo \uFEFFnorth=Bar"
	endStr := "Foo north=Bar"

	newStr := removeZeroWidthCharacters(startStr)
	assert.Equal(t, endStr, newStr)
}

func Test_recombineCitiesWithSpaces(t *testing.T) {
	deleteAllMapEntries()

	cityStr := "Old York Town east=New China"

	cityData := strings.Fields(cityStr)

	assert.Equal(t, 5, len(cityData))

	cityData = recombineCitiesWithSpaces(cityData)

	assert.Equal(t, 2, len(cityData))
	assert.Equal(t, "Old York Town", cityData[0])
	assert.Equal(t, "east=New China", cityData[1])
}
