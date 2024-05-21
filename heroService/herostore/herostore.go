package herostore

import (
	"fmt"
	"sync"

	_ "github.com/mattn/go-sqlite3"
)

type HeroStore struct {
	sync.Mutex
	heroes map[int]Hero
	nextId int
}

type Hero struct {
	Id     int    `json:"id"`
	Name   string `json:"name"`
	Damage int    `json:"damage"`
	Health int    `json:"health"`
	Gender bool   `json:"gender"`
	Class  Class  `json:"class"`
}

type Class int64

const (
	Warrior Class = 0
	Hunter  Class = 1
	Mage    Class = 2
	Priest  Class = 3
)

func New() *HeroStore {
	hs := &HeroStore{}
	hs.heroes = make(map[int]Hero)
	hs.nextId = 0
	return hs
}

func (hs *HeroStore) CreateHero(newHero ReqHero) error {
	hs.Lock()
	defer hs.Unlock()

	hero := Hero{
		Id:     hs.nextId,
		Name:   newHero.Name,
		Damage: newHero.Damage,
		Health: newHero.Health,
		Gender: newHero.Gender,
		Class:  newHero.Class,
	}

	hs.heroes[hs.nextId] = hero
	hs.nextId++
	return nil
}

func (hs *HeroStore) GetHero(id int) (Hero, error) {

	hs.Lock()
	defer hs.Unlock()

	hero, ok := hs.heroes[id]
	if !ok {
		return Hero{}, fmt.Errorf("Hero with id %d not found", id)
	}

	return hero, nil
}

func (hs *HeroStore) GetAllHeroes() ([]Hero, error) {

	hs.Lock()
	defer hs.Unlock()

	heroes := make([]Hero, 0, len(hs.heroes))
	for _, hero := range hs.heroes {
		heroes = append(heroes, hero)
	}

	return heroes, nil
}

func (hs *HeroStore) DeleteHero(id int) error {

	hs.Lock()
	defer hs.Unlock()

	// not intrested in deleted hero so _
	_, ok := hs.heroes[id]
	if !ok {
		return fmt.Errorf("Hero with id %d not found", id)
	}

	delete(hs.heroes, id)
	return nil
}

func (hs *HeroStore) GetWinner(id int, id2 int) (Hero, error) {

	hs.Lock()
	defer hs.Unlock()

	hero1, ok := hs.heroes[id]
	if !ok {
		return Hero{}, fmt.Errorf("Hero with id %d not found", id)
	}

	hero2, ok := hs.heroes[id2]
	if !ok {
		return Hero{}, fmt.Errorf("Hero with id %d not found", id)
	}

	var winner Hero
	resultHero1 := float64(hero1.Health) / float64(hero2.Damage)
	resultHero2 := float64(hero2.Health) / float64(hero1.Damage)

	if resultHero1 > resultHero2 {
		winner = hero1
	} else if resultHero1 == resultHero2 {
		winner = Hero{}
	} else {
		winner = hero2
	}
	fmt.Println(resultHero1, resultHero2)

	return winner, nil
}
