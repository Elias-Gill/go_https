package users

import (
	"net/http"

	"github.com/elias-gill/go_pokemon/servidor"
	"github.com/go-chi/chi/v5"
)

// TODO: anadir las funciones
func UserHandlers(r *chi.Mux) {
	// cada nueva accion
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		user, pasw, ok := r.BasicAuth()
		if ok {
			clave, err := servidor.IniciarSesion(user, pasw)
			if err != nil {
				w.Write([]byte(err.Error()))
			}
			w.Write([]byte(clave))
		}
	})

	r.Get("/perfil", func(w http.ResponseWriter, r *http.Request) {
		credential := r.Header["Authentication"][0] // <-- cuidar el estandar http
		if credential == "123" {
			w.Write([]byte("exitoso"))
		} else {
			w.Write([]byte("Credenciales invalidas"))
			r.Response.Status = "405"
		}
	})

	r.Post("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("crear nueva cuenta"))
	})
}
