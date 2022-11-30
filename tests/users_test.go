package test

import (
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/elias-gill/go_pokemon/server"
	"github.com/elias-gill/go_pokemon/teams"
	users "github.com/elias-gill/go_pokemon/users"
	"github.com/go-chi/chi/v5"
	"github.com/golang-jwt/jwt/v4"
)

// funcion para incializar un nuevo server de pruebas
func nuevoServerPruebas() *httptest.Server {
	anadirUsuarioDefecto()
	r := chi.NewRouter()
	r.Route("/user", users.UserHandlers)
	r.Route("/teams", teams.TeamsHandlers)
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
	x, _ := io.ReadAll(res.Body)
	clave := string(x)
	if clave == "" {
		t.Error("Credenciales mal retornadas: " + clave)
	}
}

func anadirUsuarioDefecto() {
	err := server.NewUser("Elias", "123")
	if err != nil {
		print(err.Error())
		print("\n\n")
	}
}
