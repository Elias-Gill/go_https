package routers

import (
	"encoding/json"
	"net/http"

	"github.com/elias-gill/go_pokemon/server"
	"github.com/go-chi/chi/v5"
)

// struct de nueva request
type body struct {
	Pokemon  string `json:"pokemon"`
	Position int    `json:"position"`
}

// handler for /teams path
func TeamsHandlers(r chi.Router) {
	// INFO: esta estructurado asi para poder autogenerar la documentacion de swagger
	r.Delete("/", delTeams)
	r.Get("/", getTeams)
	r.Post("/", postTeams)
}

//	@Summary		add new pokemon to the team
//	@Description	adds a new pokemon
//	@Tags			teams
//	@Accept			json
//	@Produce		json
//	@Param			q	query	string	false	"name search by q"	Format(email)
//	@Failure		400
//	@Router			/teams [post]
func postTeams(w http.ResponseWriter, r *http.Request) {
	// obtener el usuario de la request
	user, _, _ := r.BasicAuth()
	// obtener datos de la request
	var body body
	json.NewDecoder(r.Body).Decode(&body)
	err := server.AddNewPokemon(user, body.Pokemon)
	if err != nil {
		print(err.Error())
		err := httpError{Error: "Error al parsear request o pokemon no encontrado"}
		w.WriteHeader(400)
		json.NewEncoder(w).Encode(err)
		return
	}
	w.WriteHeader(200)
	return
}

//	@Summary		teams actions
//	@Description	teams actions
//	@Tags			teams
//	@Accept			json
//	@Produce		json
//	@Param			q	query	string	false	"name search by q"	Format(email)
//	@Failure		400
//	@Router			/teams [get]
func getTeams(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	user, _, _ := r.BasicAuth()
	team, err := server.SearchUserInfo(user)
	if err != nil {
		err := httpError{Error: "No se pudo obtener el team del usuario"}
		w.WriteHeader(404)
		json.NewEncoder(w).Encode(err)
		return
	}
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(&team)
}

//	@Summary		delete pokemon
//	@Description	deletes the first occurence of a pokemon or deletes the given position
//	@Tags			teams
//	@Accept			json
//	@Produce		json
//	@Param			q	query	string	false	"name search by q"	Format(email)
//	@Failure		400
//	@Router			/teams [delete]
func delTeams(w http.ResponseWriter, r *http.Request) {
	// obtener el usuario de la request
	user, _, _ := r.BasicAuth()
	// obtener datos de la request
	var body body
	json.NewDecoder(r.Body).Decode(&body)
	err := server.DeletePokemon(user, body.Pokemon)
	if err != nil {
		print(err.Error())
		err := httpError{Error: "Error al parsear request o pokemon no encontrado"}
		w.WriteHeader(400)
		json.NewEncoder(w).Encode(err)
		return
	}
	w.WriteHeader(200)
	return
}
