package main

import (
	"flag"
	"fmt"
	"github.com/prometheus/common/log"
	"github.com/rfornea/invasion/filehandling"
	"github.com/rfornea/invasion/models"
	"strconv"
)

const defaultAliensNum = 6

func main() {
	aliensPtr := flag.Int("aliens", defaultAliensNum, "number of aliens to start game with, specify with "+
		"'-aliens=N' where N is the number of aliens")
	fptr := flag.String("fpath", "testMap.csv", "file path to read from, specify with '-fpath=X' "+
		"where X is the file path")

	flag.Parse()

	fmt.Println("numAliens:", *aliensPtr)
	fmt.Println("fpath:", *fptr)
	fmt.Println()

	err := filehandling.ReadFile(*fptr)
	if err != nil {
		log.Fatal(err)
	}

	models.InitializeAliens(*aliensPtr)

	maxMoves := 10000
	var actualMoves int
	allDead := false
	allTrapped := false

	for actualMoves = 0; actualMoves < maxMoves; actualMoves++ {
		allDead, allTrapped = models.MoveAllAliens()
		if allDead || allTrapped {
			break
		}
	}

	var msg string
	if actualMoves == maxMoves {
		msg = "Invasion over after " + strconv.Itoa(maxMoves) + " turns."
	}
	if allTrapped {
		msg = "Invasion over because all remaining aliens are trapped."
	}
	if allDead {
		msg = "Invasion over because all aliens are dead."
	}
	fmt.Println("_________________________________")
	fmt.Println(msg)

	fmt.Println("_________________________________")

	fmt.Println("Remaining world:")
	fmt.Println()
	models.PrintInvasionResult()

	fmt.Println("_________________________________")
}
