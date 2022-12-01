package test

import (
	"encoding/json"
	"net/http"
	"net/url"
	"testing"
)

// nuevo server de pruebas
var ts = nuevoServerPruebas()

func iniciarSesion() (*http.Response, error) {
	// iniciar sesion con el nuevo usuario
	req, _ := http.NewRequest("GET", ts.URL+"/user/", nil)
	req.URL.User = url.UserPassword("Elias", "123")
	res, err := ts.Client().Do(req)
	if err != nil {
		return nil, err
	}
	return res, nil
}

type tokens struct {
	token string `bson:"JWT"`
}

// TODO: continuar con los tests
// test para agregar un nuevo pokemon al equipo
func TestNewPokemon(t *testing.T) {
	res, err := iniciarSesion()
	if err != nil {
		t.Error(err)
		return
	}
	// extraer y guardar el token
	var data tokens
	err = json.NewDecoder(res.Body).Decode(&data)
	if err != nil {
		t.Error(err)
		return
	}

	// requerir el team del usuario
	req, _ := http.NewRequest("GET", ts.URL+"/teams/", nil)
	req.Header.Add("Bearer", data.token)
	res, err = ts.Client().Do(req)
	if err != nil {
        t.Error(err)
		return
	}
	return
}
