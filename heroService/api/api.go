package api

import (
	"hero/hero"
	"net/http"
)

func NewApp() *http.ServeMux {
	mux := http.NewServeMux()
	store := hero.NewHeroServer()

	mux.HandleFunc("POST /hero/add", store.CreateHeroHandler)
	mux.HandleFunc("PATCH /hero/update/{id}", store.UpdateHeroHandler)
	mux.HandleFunc("GET /hero/getAll", store.GetHeroesHandler)
	mux.HandleFunc("GET /hero/get/{id}", store.GetHeroByIdHandler)
	mux.HandleFunc("GET /hero/winner/{id}/{id2}", store.GetWinnerHandler)
	mux.HandleFunc("GET /hero/winner/all", store.GetGloblaWinnerHandler)
	mux.HandleFunc("DELETE /hero/delete/{id}", store.DeleteHeroHandler)

	return mux
}
