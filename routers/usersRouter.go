package routers

import (
	"encoding/json"
	"net/http"

	servidor "github.com/elias-gill/go_pokemon/server"
	"github.com/go-chi/chi/v5"
)

type httpError struct {
	Error string `json:"error"`
}

// handlers de users/
func UserHandlers(r chi.Router) {
	// iniciar sesion
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		user, pasw, ok := r.BasicAuth() // get user credentials
		if ok {
			w.Header().Set("Content-type", "application/json")
			token, err := servidor.IniciarSesion(user, pasw)
			if err != nil {
                // error de parseo 
                println("error al parsear credenciales")
				err := httpError{Error: err.Error()}
				w.WriteHeader(405)
				json.NewEncoder(w).Encode(err)
				return
			}
			// mandar el jwt con json
			jwt := jwtResponse{Jwt: token}
			json.NewEncoder(w).Encode(jwt)
			return
		}
		// error de autenticacion
        println("Error de credenciales " + user + " " + pasw)
		err := httpError{Error: "Usuario o contrasena invalidos"}
		w.WriteHeader(405)
		json.NewEncoder(w).Encode(err)
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
		w.WriteHeader(200)
	})
}

// struct de creacion de un nuevo usuario
type newUser struct {
	UserName string `bson:"userName"`
	Password string `bson:"password"`
}

type jwtResponse struct {
	Jwt string `bson:"jwt"`
}
