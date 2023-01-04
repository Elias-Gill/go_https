package test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/elias-gill/go_pokemon/routers"
	"github.com/elias-gill/go_pokemon/server"
	"github.com/elias-gill/go_pokemon/tools"
	"github.com/go-chi/chi/v5"
)

// struct para extraer el token de la request
type tokens struct {
	Token string `json:"jwt"`
}

// funcion para incializar un nuevo server de pruebas
func nuevoServerPruebas() *httptest.Server {
	anadirUsuarioDefecto()
	r := chi.NewRouter()
	r.Use(tools.JwtMidleware)
	r.Route("/user", routers.UserHandlers)
	r.Route("/teams", routers.TeamsHandlers)
	return httptest.NewServer(r)
}

// testear la funcion de busqueda de usuarios
func TestSearchUser(t *testing.T) {
	_, err := server.SearchUser("Elias")
	if err != nil {
		println(err.Error())
	}
}

// testear la llamada autenticacion de usuario
func TestUserAuthentication(t *testing.T) {
	// nuevo server de pruebas
	ts := nuevoServerPruebas()
	defer ts.Close()

	// iniciar sesion con el nuevo usuario
	req, _ := http.NewRequest("GET", ts.URL+"/user/", nil)
	req.URL.User = url.UserPassword("Elias", "123")
	res, err := ts.Client().Do(req)
	if err != nil {
		t.Error(err)
		return
	}
	var to tokens
	json.NewDecoder(res.Body).Decode(&to)
	// comprobar que el token mandado es valido
	user, err := tools.ComprobarJWT(to.Token)
	if err != nil || user != "Elias" {
		t.Error("token invalido: " + err.Error())
		return
	}
}

// establece un usuario por defecto con un nuevo team por defecto {"Elias": {charizard, bulbasaur y pikachu}}
func anadirUsuarioDefecto() {
	err := server.DeleteUser("Elias")
	if err != nil && err.Error() != "El usuario ya existe" {
		println(err.Error())
	}
}
