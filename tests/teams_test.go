package test

import (
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"testing"

	"github.com/elias-gill/go_pokemon/tools"
)

var (
	ts    = nuevoServerPruebas() // sever de pruebas
	token = iniciarSesion()      // jwt de usuario de pruebas
)

// struct para un pokemon individual
type pokemon struct {
	Name  string `json:"name"`
	Type  string `json:"type"`
	Power int    `json:"power"`
}

// struct para decodificacion del team pokemon del json
type teamsBody struct {
	Team []pokemon `json:"team"`
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
	if user, err := tools.ComprobarJWT(t.Token); err != nil || user != "Elias" {
		panic("Invalid token in iniciarSesion()")
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
		t.Errorf("Error de autenticacion: %s", err.Error())
		return
	}

	// parsear el equipo de la response de la api
	var userTeam teamsBody
	json.NewDecoder(res.Body).Decode(&userTeam)

	// comprobar el largo del equipo
	if len(userTeam.Team) != 1 {
		t.Errorf("Largo de team invalido: %d \n Esperado: 1\n", len(userTeam.Team))
		return
	}

	// comprobar el pokemon inicial del usuario
	if userTeam.Team[0].Name != "charizard" {
		t.Errorf("pokemon inicial invalido: \n %s", userTeam.Team[0].Name)
		return
	}
}

// test para probar la creacion del team por defecto
func TestNewPokemonTeam(t *testing.T) {
}
