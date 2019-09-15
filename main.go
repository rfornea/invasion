package main

import (
	"flag"
	"fmt"
	"github.com/rfornea/invasion/filehandling"
)

const defaultAliensNum = 42

func main() {
	aliensPtr := flag.Int("aliens", defaultAliensNum, "number of aliens to start game with, specify with " +
		"'-aliens=N' where N is the number of aliens")
	fptr := flag.String("fpath", "cityMap.txt", "file path to read from")

	flag.Parse()

	fmt.Println("numAliens:", *aliensPtr)
	fmt.Println("fpath:", *fptr)
	fmt.Println()

	filehandling.ReadFile(*fptr)
}
