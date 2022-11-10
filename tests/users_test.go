package test

import (
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/elias-gill/go_pokemon/teams"
	users "github.com/elias-gill/go_pokemon/users"
	"github.com/go-chi/chi/v5"
)

// funcion para incializar un nuevo server de pruebas
func nuevoServerPruebas() *httptest.Server {
	r := chi.NewRouter()
    r.Route("/user", users.UserHandlers)
    r.Route("/teams", teams.TeamsHandlers)
	return httptest.NewServer(r)
}

// testear la llamada autenticacion de usuario
func TestUserAuthentication(t *testing.T) {
	// nuevo server de pruebas
	ts := nuevoServerPruebas()
	defer ts.Close()

	// iniciar sesion con el nuevo usuario
	req, _ := http.NewRequest("GET", ts.URL+"/user/", nil)
	req.URL.User = url.UserPassword("elias", "123")
	res, err := ts.Client().Do(req)
	if err != nil {
		t.Error(err)
		return
	}
	x, _ := io.ReadAll(res.Body)
	clave := string(x)
	if clave != "123" {
		t.Error("Credenciales mal retornadas: " + string(x))
	}
}

	/* // iniciar sesion con el nuevo usuario
	req, _ = http.NewRequest("GET", ts.URL+"/user/", nil)
	req.Header.Set("Authentication", clave)
	res, err = ts.Client().Do(req)
	if err != nil {
		t.Error(err)
		return
	}
	x, _ = io.ReadAll(res.Body)
	if string(x) != "exitoso" {
		t.Error("Credenciales mal retornadas: " + string(x))
	} */
