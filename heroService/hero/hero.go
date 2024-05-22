package hero

import (
	"database/sql"
	"hero/database"
	"sync"
)

type heroSever struct {
	server *HeroStore
}

func NewHeroServer() *heroSever {
	store := New()
	return &heroSever{server: store}
}

type ReqHero struct {
	Name   string `json:"name"`
	Damage int    `json:"damage"`
	Health int    `json:"health"`
	Gender bool   `json:"gender"`
	Class  Class  `json:"class"`
}

type HeroStore struct {
	sync.Mutex
	heroes *sql.DB
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
	hs.heroes = database.InitDB("./heroes-db.db")
	return hs
}
