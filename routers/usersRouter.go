package routers

import (
	"encoding/json"
	"net/http"

	servidor "github.com/elias-gill/go_pokemon/server"
	"github.com/go-chi/chi/v5"
)

// handlers de users/
func UserHandlers(r chi.Router) {
	// iniciar sesion
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		user, pasw, ok := r.BasicAuth() // get user credentials
		if ok {
			jwt, err := servidor.IniciarSesion(user, pasw)
			if err != nil {
				w.Write([]byte(err.Error()))
			}
			w.Header().Add("Content-type", "application/json")
			w.Write([]byte("JWT: " + jwt))
		}
	})

	// anadir un nuevo usuario a la base de datos
	r.Post("/", func(w http.ResponseWriter, r *http.Request) {
		var data newUser
		err := json.NewDecoder(r.Body).Decode(&data)
		if err != nil {
			w.Write([]byte("Body request invalido"))
			w.Write([]byte(err.Error()))
			w.Write([]byte("\n\n" + data.UserName))
			return
		}
		// crear el nuevo usuario
		servidor.NewUser(data.UserName, data.Password)
		w.Write([]byte("cuenta creada satisfactoriamente"))
	})
}

// struct de creacion de un nuevo usuario
type newUser struct {
	UserName string `bson:"userName"`
	Password string `bson:"password"`
}
