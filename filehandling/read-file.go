package filehandling

import (
	"bufio"
	"log"
	"os"
	"github.com/rfornea/invasion/models"
)

func ReadFile(fileName string) {
	f, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err = f.Close(); err != nil {
			log.Fatal(err)
		}
	}()
	s := bufio.NewScanner(f)
	for s.Scan() {
		models.AddCityToMap(s.Text())
	}
	err = s.Err()
	if err != nil {
		log.Fatal(err)
	}
}
