package server

import (
	"io"
	"net/http"
)

const pokeapi = "https://pokeapi.co/api/v2/pokemon/"

func getPokemonFromApi(name string) (*string, error) {
    // realizar peticion
	query := pokeapi + name
	res, err := http.Get(query)
	if err != nil {
		return nil, err
	}
    // leer body y pasar el string
	r, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	body := string(r)
	return &body, nil
}
