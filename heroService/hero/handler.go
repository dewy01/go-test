package hero

import (
	"hero/util"
	"log"
	"net/http"
	"strconv"
)

func (hs *heroSever) CreateHeroHandler(w http.ResponseWriter, req *http.Request) {
	log.Printf("handling hero create at %s\n", req.URL.Path)

	hero, err := util.Decode[ReqHero](req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	hs.server.CreateHero(hero)

	util.Encode(w, req, http.StatusOK, hero)

}

func (hs *heroSever) UpdateHeroHandler(w http.ResponseWriter, req *http.Request) {
	id, err := strconv.Atoi(req.PathValue("id"))
	if err != nil {
		http.Error(w, "Invalid hero id", http.StatusBadRequest)
		return
	}

	hero, err := util.Decode[ReqHero](req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := hs.server.UpdateHero(id, hero); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	util.Encode(w, req, http.StatusOK, hero)
}

func (hs *heroSever) GetHeroesHandler(w http.ResponseWriter, req *http.Request) {
	heroes, err := hs.server.GetAllHeroes()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	util.Encode(w, req, http.StatusOK, heroes)
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
	util.Encode(w, req, http.StatusOK, hero)
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
	util.Encode(w, req, http.StatusOK, id)
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
	util.Encode(w, req, http.StatusOK, winner)
}

func (hs *heroSever) GetGloblaWinnerHandler(w http.ResponseWriter, req *http.Request) {

	winner, err3 := hs.server.GetGlobalWinner()
	if err3 != nil {
		http.Error(w, err3.Error(), http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	util.Encode(w, req, http.StatusOK, winner)
}
