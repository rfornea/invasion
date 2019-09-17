package models

type Alien struct {
	Number      int
	CurrentCity *City
}

var AlienMap map[int]Alien

func init() {
	AlienMap = make(map[int]Alien)
}

func CreateAliens(numAliens int) {
	for i := 1; i <= numAliens; i++ {
		alien := Alien{
			Number:      i,
			CurrentCity: nil,
		}
		AlienMap[i] = alien
	}
}

func InitializeAliens(numAliens int) {
	CreateAliens(numAliens)
	AssignAliensToCities()
}

func AssignAliensToCities() {
	for k := range AlienMap {
		alien := AlienMap[k]
		city := randCity(CityMap)
		alien.CurrentCity = city
		AlienMap[k] = alien
		city.addAlien(AlienMap[k])
	}
}

func MoveAllAliens() (allDead, allTrapped bool) {
	allDead = len(AlienMap) == 0
	allTrapped = true
	for k, v := range AlienMap {
		alien := v
		if len((*alien.CurrentCity).Links) == 0 {
			// alien is trapped, skip
			continue
		}
		allTrapped = false
		city := alien.CurrentCity.randLink()
		alien.CurrentCity.removeAlien(alien)
		alien.CurrentCity = city
		AlienMap[k] = alien
		city.addAlien(alien)
	}
	return
}
