package routers

import (
	"encoding/json"
	"net/http"

	"github.com/elias-gill/go_pokemon/server"
	"github.com/go-chi/chi/v5"
)

// struct de nueva request
type body struct {
	Pokemon string `json:"pokemon"`
}

// handler sobre la ruta "/teams"
func TeamsHandlers(r chi.Router) {
    // retornar el team del equipo
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-type", "application/json")
		user, _, _ := r.BasicAuth()
		team, err := server.SearchUserInfo(user)
		if err != nil {
			w.WriteHeader(404)
			w.Write([]byte("No se pudo obtener el team del usuario"))
			return
		}
		w.WriteHeader(200)
		json.NewEncoder(w).Encode(&team)
	})

    // anadir un pokemon al equipo
	r.Post("/", func(w http.ResponseWriter, r *http.Request) {
		// obtener el usuario de la request
		user, _, _ := r.BasicAuth()
		// obtener datos de la request
		var body body
		json.NewDecoder(r.Body).Decode(&body)
		err := server.AddNewPokemon(user, body.Pokemon)
		if err != nil {
			print(err.Error())
			w.WriteHeader(400)
			w.Write([]byte("Error al parsear request o pokemon no encontrado"))
			return
		}
		w.WriteHeader(200)
		w.Write([]byte("Insercion exitosa del nuevo pokemon"))
		return
	})

    // eliminar un pokemon en cierta posicion o la primera ocurrencia del nombre
	r.Post("/", func(w http.ResponseWriter, r *http.Request) {
		// obtener el usuario de la request
		user, _, _ := r.BasicAuth()
		// obtener datos de la request
		var body body
		json.NewDecoder(r.Body).Decode(&body)
		err := server.AddNewPokemon(user, body.Pokemon)
		if err != nil {
			print(err.Error())
			w.WriteHeader(400)
			w.Write([]byte("Error al parsear request o pokemon no encontrado"))
			return
		}
		w.WriteHeader(200)
		w.Write([]byte("Insercion exitosa del nuevo pokemon"))
		return
	})
}
