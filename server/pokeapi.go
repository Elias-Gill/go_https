package server

import (
	"encoding/json"
	"net/http"
)

const pokeapi = "https://pokeapi.co/api/v2/pokemon/"

type genericInfo struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

type move struct {
	Move genericInfo `json:"move"`
}

type pktype struct {
	Type genericInfo `json:"type"`
}

type sprites struct {
	Other other `json:"other"`
}

type other struct {
	Official_artwork official_artwork `json:"official-artwork"`
}

type official_artwork struct {
	Image string `json:"front_default"`
}

type pokemon struct {
	Base_experience int      `json:"base_experience"`
	Sprites         sprites  `json:"sprites"`
	Height          int      `json:"height"`
	Weight          int      `json:"weight"`
	Id              int      `json:"id"`
	Moves           []move   `json:"moves"`
	Name            string   `json:"name"`
	Order           int      `json:"order"`
	Types           []pktype `json:"types"`
}

func getPokemonFromApi(pkName string) (*pokemon, error) {
	// realizar peticion
	query := pokeapi + pkName
	res, err := http.Get(query)
	if err != nil {
		return nil, err
	}
	// leer body y pasar el string
	var pokemon pokemon
	err = json.NewDecoder(res.Body).Decode(&pokemon)
	if err != nil {
		return nil, err
	}
	return &pokemon, nil
}
