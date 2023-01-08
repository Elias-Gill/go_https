package test

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"

	"github.com/elias-gill/go_pokemon/authentication"
	"github.com/elias-gill/go_pokemon/routers"
	"github.com/elias-gill/go_pokemon/server"
	"github.com/go-chi/chi/v5"
)

var (
	ts    = nuevoServerPruebas() // sever de pruebas
	token = iniciarSesion()      // jwt de usuario de pruebas
)

// struct para extraer el token de la request
type tokens struct {
	Token string `json:"jwt"`
}

// funcion para incializar un nuevo server de pruebas
func nuevoServerPruebas() *httptest.Server {
	anadirUsuarioDefecto()
	r := chi.NewRouter()
	r.Use(authentication.JwtMidleware)
	r.Route("/user", routers.UserHandlers)
	r.Route("/teams", routers.TeamsHandlers)
	return httptest.NewServer(r)
}

// funcion para iniciar sesion en el usuario de pruebas. Consigue el jwt y lo guarda
// en la variable global toke
func iniciarSesion() string {
	// iniciar sesion con el nuevo usuario
	req, _ := http.NewRequest("GET", ts.URL+"/user/", nil)
	req.URL.User = url.UserPassword("Elias", "123")
	res, err := ts.Client().Do(req)
	if err != nil || res.StatusCode != 200 {
		println(io.ReadAll(res.Body))
		panic("Error en el inicio de sesion: status " + res.Status)
	}
	// extraer y guardar el token
	var t tokens
	err = json.NewDecoder(res.Body).Decode(&t)
	if err != nil || t.Token == "" {
		panic("Cannot parse token")
	}

	// comprobar que el token sea valido
	if user, err := authentication.ComprobarJWT(t.Token); err != nil || user != "Elias" {
		panic("Invalid token in iniciarSesion()")
	}
	return t.Token
}

// establece un usuario por defecto con un nuevo team por defecto {"Elias": {charizard, bulbasaur y pikachu}}
func anadirUsuarioDefecto() {
	server.DeleteUser("Elias")
	err := server.NewUser("Elias", "123")
	if err != nil && err.Error() != "El usuario ya existe" {
		println(err.Error())
	}
}
