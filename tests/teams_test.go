package test

import (
	"testing"
)

// testear que se este recibiendo el team indicado
func TestGetPokemonTeam(t *testing.T) {
	user, err := nuevaRequestTeams("GET", nil)
	if err != nil {
		t.Errorf("Error en la llamada get: %s", err.Error())
		return
	}
	if len(user.Team) != 3 {
		t.Errorf("Error en el largo del team: \n Esperado: 3 \tEncontrado: %d", len(user.Team))
		return
	}
	team := user.Team
	if team[0].Name != "charizard" || team[1].Name != "bulbasaur" || team[2].Name != "pikachu" {
		t.Errorf("Pokemones en el team incorrectos")
		return
	}
}

// testear el anadir un nuevo pokemon a la lista
func TestAddNewPokemon(t *testing.T) {
	body, err := nuevaRequestTeams("POST", &responseBody{Pokemon: "charmander"})
	if err != nil {
		t.Errorf("NO se pudo realizar la request: %s", err.Error())
		return
	}
	// volver a pedir el team actualizado
	body, err = nuevaRequestTeams("GET", nil)
	if err != nil {
		t.Errorf("NO se pudo realizar la request: %s", err.Error())
		return
	}
	if len(body.Team) != 4 {
		t.Errorf("NO se anadio el pokemon correcto (charmander): %d", len(body.Team))
		return
	}
}

// testear el eliminar un nuevo pokemon a la lista
func TestDeletePokemon(t *testing.T) {
	// eliminar el pokemon
	_, err := nuevaRequestTeams("DELETE", &responseBody{Pokemon: "charmander"})
	if err != nil {
		t.Errorf("NO se pudo realizar la request de DELETE: %s", err.Error())
		return
	}
	// comprobar que se elimino
	body, err := nuevaRequestTeams("GET", &responseBody{Pokemon: "charmander"})
	if err != nil {
		t.Errorf("NO se pudo realizar la request de GET: %s", err.Error())
		return
	}
	team := body.Team
	if team[len(team)-1].Name != "pikachu" { // ultimo debe de ser ahora pikachu
		t.Errorf("NO se ELIMINO el pokemon correcto (charmander)")
		return
	}
}
