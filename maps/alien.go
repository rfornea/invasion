package maps

import "fmt"

type alien struct {
	number      int
	currentCity *City
}

var alienMap map[int]alien

func init() {
	alienMap = make(map[int]alien)
}

func createAliens(numAliens int) {
	for i := 1; i <= numAliens; i++ {
		newAlien := alien{
			number:      i,
			currentCity: nil,
		}
		alienMap[i] = newAlien
	}
}

/*InitializeAliens is called by main at program start to create the aliens and assign them their initial cities*/
func InitializeAliens(numAliens int) bool {
	createAliens(numAliens)
	return assignAliensToCities()
}

func assignAliensToCities() bool {
	for k := range alienMap {
		currentAlien := alienMap[k]
		if len(CityMap) == 0 {
			return true
		}
		city := randCity(CityMap)
		currentAlien.currentCity = city
		alienMap[k] = currentAlien
		city.addAlien(&currentAlien)
	}
	return false
}

/*MoveAllAliens is called by main for each "turn" to move all the aliens to new cities*/
func MoveAllAliens() (allDead, allTrapped bool) {
	fmt.Println(CityMap)
	allDead = len(alienMap) == 0
	allTrapped = true
	for _, v := range alienMap {
		currentAlien := v
		if len((*currentAlien.currentCity).Links) == 0 {
			// alien is trapped, skip
			continue
		}
		allTrapped = false
		currentAlien.move()
	}
	return
}

func (a *alien) move() {
	city := a.currentCity.randLink()
	a.currentCity.removeAlien(a)
	a.currentCity = city
	alienMap[a.number] = *a
	a.currentCity.addAlien(a)
}
