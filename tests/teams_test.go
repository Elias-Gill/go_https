package test

import (
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"testing"
)

var (
	ts    = nuevoServerPruebas() // sever de pruebas
	token = iniciarSesion()      // jwt de usuario de pruebas
)

// struct para extraer el token de la request
type tokens struct {
	Token string `json:"jwt"`
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
	return t.Token
}

// test para agregar un nuevo pokemon al equipo
func TestGetPokemonTeam(t *testing.T) {
	// requerir el team del usuario
	req, _ := http.NewRequest("GET", ts.URL+"/teams/", nil)
	req.Header.Set("Authorization", "Bearer "+token)
	res, err := ts.Client().Do(req)
	if err != nil {
		t.Errorf("No se pudo get teams pokemon: %s", err)
		return
	}
	if res.StatusCode != 200 {
		t.Errorf("Status not ok de get teams")
        x, _ := io.ReadAll(res.Body)
		println("que ? " + string(x))
		return
	}
}
