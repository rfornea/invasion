package filehandling

import (
	"github.com/rfornea/invasion/maps"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func Test_ReadFile(t *testing.T) {
	filePath := "/tmp/unitTestMap.txt"

	f, err := os.Create(filePath)
	assert.Nil(t, err)
	defer f.Close()

	foo := "Foo"
	north := "north"
	bar := "Bar"
	south := "south"

	_, err = f.WriteString(foo + " " + north + "=" + bar + "\n" + bar + " " + south + "=" + foo)
	assert.Nil(t, err)

	f.Sync()

	err = ReadFile(filePath)
	assert.Nil(t, err)

	errMsg := "city map not built as expected"

	if _, ok := maps.CityMap[foo]; !ok {
		t.Fatal(errMsg)
	}

	if _, ok := maps.CityMap[bar]; !ok {
		t.Fatal(errMsg)
	}

	if _, ok := maps.CityMap[foo].Links[north]; !ok {
		t.Fatal(errMsg)
	}

	if _, ok := maps.CityMap[bar].Links[south]; !ok {
		t.Fatal(errMsg)
	}

	if maps.CityMap[foo].Links[north].Name != bar {
		t.Fatal(errMsg)
	}

	if maps.CityMap[bar].Links[south].Name != foo {
		t.Fatal(errMsg)
	}

	os.Remove(filePath)
}
