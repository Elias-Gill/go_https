package server

import (
	"io"
	"net/http"
)

const pokeapi = "https://pokeapi.co/api/v2/pokemon/"

func getPokemonFromApi(name string) (error, *string) {
	query := pokeapi + name
	req, _ := http.NewRequest("get", query, nil)
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return err, nil
	}
	body, _ := io.ReadAll(res.Body)
    r := string(body)
	return nil, &r
}
