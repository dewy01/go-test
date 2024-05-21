package herostore

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
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

func (hs *heroSever) CreateHeroHandler(w http.ResponseWriter, req *http.Request) {
	log.Printf("handling hero create at %s\n", req.URL.Path)
	var reqHero ReqHero

	err := json.NewDecoder(req.Body).Decode(&reqHero)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	hs.server.CreateHero(reqHero)
	json.NewEncoder(w).Encode(http.StatusOK)

}

func (hs *heroSever) GetHeroesHandler(w http.ResponseWriter, req *http.Request) {
	heroes, err := hs.server.GetAllHeroes()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(heroes)
}

func (hs *heroSever) GetHeroByIdHandler(w http.ResponseWriter, req *http.Request) {
	id, err := strconv.Atoi(req.PathValue("id"))
	if err != nil {
		http.Error(w, "Invalid hero id", http.StatusBadRequest)
		return
	}

	hero, err := hs.server.GetHero(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(hero)
}

func (hs *heroSever) DeleteHeroHandler(w http.ResponseWriter, req *http.Request) {
	id, err := strconv.Atoi(req.PathValue("id"))
	if err != nil {
		http.Error(w, "Invalid hero id", http.StatusBadRequest)
		return
	}

	err2 := hs.server.DeleteHero(id)
	if err2 != nil {
		http.Error(w, err2.Error(), http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(http.StatusOK)
}

func (hs *heroSever) GetWinnerHandler(w http.ResponseWriter, req *http.Request) {
	id, err1 := strconv.Atoi(req.PathValue("id"))
	if err1 != nil {
		http.Error(w, "Invalid hero id", http.StatusBadRequest)
		return
	}

	id2, err2 := strconv.Atoi(req.PathValue("id2"))
	if err2 != nil {
		http.Error(w, "Invalid hero id", http.StatusBadRequest)
		return
	}

	winner, err3 := hs.server.GetWinner(id, id2)
	if err3 != nil {
		http.Error(w, err3.Error(), http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(winner)
}
