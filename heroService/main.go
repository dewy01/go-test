package main

import (
	"hero/herostore"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	server := herostore.NewHeroServer()

	mux.HandleFunc("POST /hero/add", server.CreateHeroHandler)
	mux.HandleFunc("GET /hero/getAll", server.GetHeroesHandler)
	mux.HandleFunc("GET /hero/get/{id}", server.GetHeroByIdHandler)
	mux.HandleFunc("GET /hero/winner/{id}/{id2}", server.GetWinnerHandler)
	mux.HandleFunc("DELETE /hero/delete/{id}", server.DeleteHeroHandler)

	http.ListenAndServe("localhost:8080", mux)
}
