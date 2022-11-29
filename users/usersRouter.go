package users

import (
	"net/http"

	servidor "github.com/elias-gill/go_pokemon/server"
	"github.com/go-chi/chi/v5"
)

// TODO: anadir las funciones
func UserHandlers(r chi.Router) {
	// iniciar sesion
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		user, pasw, ok := r.BasicAuth() // get user credentials
		if ok {
			clave, err := servidor.IniciarSesion(user, pasw)
			if err != nil {
				w.Write([]byte(err.Error()))
			}
			w.Write([]byte(clave))
		}
	})

	// crear usuario nuevo
	r.Post("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("crear nueva cuenta"))
	})
}
