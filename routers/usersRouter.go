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
            w.Header().Set("Content-type", "application/json")
			token, err := servidor.IniciarSesion(user, pasw)
			if err != nil {
				w.WriteHeader(405) // error de autenticacion
				w.Write([]byte(err.Error()))
				return
			}
			// mandar el jwt con json
			jwt := jwtResponse{Jwt: token}
			json.NewEncoder(w).Encode(jwt)
			return
		}
		w.WriteHeader(405) // error de autenticacion
        w.Write([]byte("Usuario o contrasena invalidos"))
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

// Gill, sos medio boludo y te soles olvidar de que las variables
// para poder hacer el encoding tenes que poner en mayusculas(exportar)
type jwtResponse struct {
	Jwt string `bson:"jwt"`
}
