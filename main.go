package main

import (
	"flag"
	"fmt"
	"github.com/prometheus/common/log"
	"github.com/rfornea/invasion/filehandling"
	"github.com/rfornea/invasion/maps"
	"strconv"
)

const defaultAliensNum = 6

func main() {
	aliensPtr := flag.Int("aliens", defaultAliensNum, "number of aliens to start game with, specify with "+
		"'-aliens=N' where N is the number of aliens")
	fptr := flag.String("fpath", "testMap.csv", "file path to read from, specify with '-fpath=X' "+
		"where X is the file path")

	flag.Parse()

	if *aliensPtr < 2 {
		log.Fatal("You need at least 2 aliens for an invasion!")
	}

	fmt.Println("numAliens:", *aliensPtr)
	fmt.Println("fpath:", *fptr)
	fmt.Println()

	err := filehandling.ReadFile(*fptr)
	if err != nil {
		log.Fatal(err)
	}

	allCitiesDestroyed := maps.InitializeAliens(*aliensPtr)

	maxMoves := 10000
	var actualMoves int
	allDead := false
	allTrapped := false

	for actualMoves = 0; actualMoves < maxMoves; actualMoves++ {
		if allCitiesDestroyed || allDead || allTrapped {
			break
		}
		allDead, allTrapped = maps.MoveAllAliens()
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
	if allCitiesDestroyed {
		msg = "Invasion over because all cities were destroyed (you may have created too many aliens for the number of cities in your map)."
	}

	fmt.Println("_________________________________")
	fmt.Println(msg)

	fmt.Println("_________________________________")

	fmt.Println("Remaining world:")
	fmt.Println()
	maps.PrintInvasionResult()

	fmt.Println("_________________________________")
}
