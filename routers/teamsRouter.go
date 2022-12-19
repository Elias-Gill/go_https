package routers

import (
	"encoding/json"
	"net/http"

	"github.com/elias-gill/go_pokemon/server"
	"github.com/go-chi/chi/v5"
)

// TODO: conectar con la db
func TeamsHandlers(r chi.Router) {
	// iniciar sesion
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		w.Header().Set("Content-type", "application/json")
		// aux := teamsBody{Team: []pokemon{{Name: "charizard", Power: 12, Type: "fuego"}}}
        user, _, _ := r.BasicAuth()
        aux, err := server.GetTeamFromUser(user)
        if err != nil {
            w.Write([]byte("Ocurrio un error inesperado"))
            return
        }
		json.NewEncoder(w).Encode(&aux)
	})

	// anadir un nuevo usuario a la base de datos
	r.Post("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("post de manera exitosa"))
	})
}
