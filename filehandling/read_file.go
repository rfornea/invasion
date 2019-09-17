package filehandling

import (
	"bufio"
	"github.com/rfornea/invasion/models"
	"log"
	"os"
)

func ReadFile(fileName string) error {
	f, err := os.Open(fileName)
	if err != nil {
		return err
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
		return err
	}
	return nil
}
