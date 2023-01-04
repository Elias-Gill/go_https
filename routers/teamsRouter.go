package routers

import (
	"encoding/json"
	"net/http"

	"github.com/elias-gill/go_pokemon/server"
	"github.com/go-chi/chi/v5"
)

func TeamsHandlers(r chi.Router) {
	// iniciar sesion
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-type", "application/json")
		// aux := teamsBody{Team: []pokemon{{Name: "charizard", Power: 12, Type: "fuego"}}}
		user, _, _ := r.BasicAuth()
		team, err := server.GetTeamFromUser(user)
		if err != nil {
			w.WriteHeader(404)
			w.Write([]byte("No se pudo obtener el team del usuario"))
			return
		}
		w.WriteHeader(200)
		json.NewEncoder(w).Encode(&team)
	})

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		// anadir un nuevo usuario a la base de datos
		user, _, _ := r.BasicAuth()
        var body body
        json.NewDecoder(r.Body).Decode(&body)
		err := server.AddNewPokemon(user, body.Pokemon)
        if err != nil {
            w.WriteHeader(400)
            w.Write([]byte("Error al parsear request o pokemon no encontrado"))
        } else {
            w.WriteHeader(200)
            w.Write([]byte("Insercion exitosa del nuevo pokemon"))
        }
	})
}

type body struct {
    Pokemon string `json:"pokemon"`
}
