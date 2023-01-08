package test

import (
	"encoding/json"
	"net/http"
	"net/url"
	"testing"

	"github.com/elias-gill/go_pokemon/authentication"
)

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
	user, err := authentication.ComprobarJWT(to.Token)
	if err != nil || user != "Elias" {
		t.Error("token invalido: " + err.Error())
		return
	}
}
