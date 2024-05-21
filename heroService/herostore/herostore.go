package herostore

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"sync"

	_ "github.com/mattn/go-sqlite3"
)

type HeroStore struct {
	sync.Mutex
	heroes *sql.DB
}

func initDB(filepath string) *sql.DB {
	if _, err := os.Stat(filepath); os.IsNotExist(err) {
		if _, err := os.Create(filepath); err != nil {
			return nil
		}
	}

	db, err := sql.Open("sqlite3", filepath)
	if err != nil {
		log.Fatal(err)
	}

	createTableSQL := `CREATE TABLE IF NOT EXISTS heroes (
        "id" INTEGER PRIMARY KEY AUTOINCREMENT,
        "name" TEXT,
        "damage" INTEGER,
        "health" INTEGER,
        "gender" BOOLEAN,
        "class" INTEGER
    );`

	statement, err := db.Prepare(createTableSQL)
	if err != nil {
		log.Fatal(err)
	}
	statement.Exec()
	log.Println("Heroes table created successfully")
	return db
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
	hs.heroes = initDB("./heroes-db.db")
	return hs
}

func (hs *HeroStore) CreateHero(newHero ReqHero) error {
	hs.Lock()
	defer hs.Unlock()

	insertHeroSQL := `INSERT INTO heroes(name, damage, health, gender, class) VALUES (?, ?, ?, ?, ?)`
	statement, err := hs.heroes.Prepare(insertHeroSQL)
	if err != nil {
		log.Fatal(err)
	}
	_, err = statement.Exec(newHero.Name, newHero.Damage, newHero.Health, newHero.Gender, newHero.Class)
	if err != nil {
		log.Fatal(err)
	}
	return nil
}

func (hs *HeroStore) UpdateHero(id int, newHero ReqHero) error {
	hs.Lock()
	defer hs.Unlock()

	updateHeroSQL := `UPDATE heroes SET name = ?, damage = ?, health = ?, gender = ?, class = ? WHERE id = ?`
	statement, err := hs.heroes.Prepare(updateHeroSQL)
	if err != nil {
		log.Fatal(err)
		return err
	}
	_, err = statement.Exec(newHero.Name, newHero.Damage, newHero.Health, newHero.Gender, newHero.Class, id)
	if err != nil {
		log.Fatal(err)
		return err
	}

	return nil
}

func (hs *HeroStore) GetHero(id int) (Hero, error) {
	hs.Lock()
	defer hs.Unlock()

	hero := Hero{}
	query := `SELECT id, name, damage, health, gender, class FROM heroes WHERE id = ?`
	row := hs.heroes.QueryRow(query, id)
	err := row.Scan(&hero.Id, &hero.Name, &hero.Damage, &hero.Health, &hero.Gender, &hero.Class)
	if err != nil {
		return hero, err
	}
	return hero, nil
}

func (hs *HeroStore) GetAllHeroes() ([]Hero, error) {
	hs.Lock()
	defer hs.Unlock()

	heroes := []Hero{}
	query := `SELECT id, name, damage, health, gender, class FROM heroes`
	rows, err := hs.heroes.Query(query)
	if err != nil {
		log.Fatal(err)
	}

	for rows.Next() {
		var hero Hero
		err := rows.Scan(&hero.Id, &hero.Name, &hero.Damage, &hero.Health, &hero.Gender, &hero.Class)
		if err != nil {
			return nil, err
		}
		heroes = append(heroes, hero)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return heroes, nil
}

func (hs *HeroStore) DeleteHero(id int) error {
	hs.Lock()
	defer hs.Unlock()

	result, err := hs.heroes.Exec("DELETE FROM heroes HERE id = ?", id)
	if err != nil {
		return err
	}

	rowsAffected, err2 := result.RowsAffected()
	if err2 != nil {
		return err2
	}

	if rowsAffected == 0 {
		return fmt.Errorf("Hero with id %d not found", id)
	}

	return nil
}

func (hs *HeroStore) GetWinner(id int, id2 int) (Hero, error) {
	hs.Lock()
	defer hs.Unlock()

	hero1 := Hero{}
	query := `SELECT id, name, damage, health, gender, class FROM heroes WHERE id = ?`
	row := hs.heroes.QueryRow(query, id)
	err := row.Scan(&hero1.Id, &hero1.Name, &hero1.Damage, &hero1.Health, &hero1.Gender, &hero1.Class)
	if err != nil {
		return Hero{}, err
	}

	hero2 := Hero{}
	query2 := `SELECT id, name, damage, health, gender, class FROM heroes WHERE id = ?`
	row2 := hs.heroes.QueryRow(query2, id2)
	err2 := row2.Scan(&hero2.Id, &hero2.Name, &hero2.Damage, &hero2.Health, &hero2.Gender, &hero2.Class)
	if err2 != nil {
		return Hero{}, err2
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


func (hs *HeroStore) GetGlobalWinner() (Hero, error) {
	heroes, err := hs.GetAllHeroes()
	if err != nil {
		return Hero{}, err
	}

	var lostMap = make(map[int]bool)
	var wg sync.WaitGroup

	for i := 0; i < len(heroes); i++ {
		for j := i + 1; j < len(heroes); j++ {
			if _, lost := lostMap[heroes[i].Id]; lost {
				continue
			}
			if _, lost := lostMap[heroes[j].Id]; lost {
				continue
			}
			wg.Add(1)
			go CompareHeroes(heroes[i], heroes[j], &wg, lostMap)
		}
	}

	wg.Wait()

	for _, hero := range heroes {
		if _, lost := lostMap[hero.Id]; !lost {
			return hero, nil
		}
	}

	return Hero{}, fmt.Errorf("no single best hero found")
}

func CompareHeroes(hero1, hero2 Hero, wg *sync.WaitGroup, lostMap map[int]bool) {
	defer wg.Done()
	winner := CalculateWinner(hero1, hero2)
	if winner.Id != 0 {
		if winner.Id == hero1.Id {
			lostMap[hero2.Id] = true
		} else if winner.Id == hero2.Id {
			lostMap[hero1.Id] = true
		}
	}
}

func CalculateWinner(hero1, hero2 Hero) Hero {
	resultHero1 := float64(hero1.Health) / float64(hero2.Damage)
	resultHero2 := float64(hero2.Health) / float64(hero1.Damage)

	if resultHero1 > resultHero2 {
		return hero1
	} else if resultHero1 < resultHero2 {
		return hero2
	}
	return Hero{}
}