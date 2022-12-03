package routers

import (
	"net/http"
	"github.com/go-chi/chi/v5"
)

func TeamsHandlers(r chi.Router) {
	// iniciar sesion
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
        w.Write([]byte("get de manera exitosa"))
        // TODO: realizar la busqueda y esas cosas
	})

	// anadir un nuevo usuario a la base de datos
	r.Post("/", func(w http.ResponseWriter, r *http.Request) {
        w.Write([]byte("post de manera exitosa"))
	})
}

// struct de creacion de un nuevo usuario
type newTeam struct {
	UserName string `bson:"userName"`
	Password string `bson:"password"`
}
