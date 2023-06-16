package server

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
)

// adds the chossen pokemon to the given username team
func AddNewPokemon(user string, pokName string) error {
	u, err := SearchUserInfo(user)
	if err != nil {
		return err
	}

	// buscar el team y comprobar que no se encuentre lleno
	team := u.Team
	if len(team) == 5 {
		return fmt.Errorf("El team del usuario se encuentra lleno")
	}

	// anadir nuevo pokemon desde la pokeapi
	pokemon, err := getPokemonFromApi(pokName)
	if err != nil {
		return err
	}
	team = append(team, *pokemon)

	// guardar en la db
	filter := bson.D{{Key: "userName", Value: user}}
	update := bson.D{{Key: "$set", Value: bson.D{{Key: "team", Value: team}}}}
	_, err = database.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return err
	}
	return nil
}

// eliminar la primera aparicion de un pokemon en el team
func DeletePokemon(user string, pokName string) error {
	u, err := SearchUserInfo(user)
	if err != nil {
		return err
	}

	// buscar el team y comprobar que no se encuentre lleno
	team := u.Team
	if len(team) == 5 {
		return fmt.Errorf("El team del usuario se encuentra lleno")
	}

	// eliminar la primera ocurrencia
	var aux []pokemon
	eliminado := false
	for _, pk := range team {
		if pk.Name == pokName && !eliminado {
			eliminado = true
			continue
		}
		aux = append(team, pk)
	}

	// guardar en la db
	filter := bson.D{{Key: "userName", Value: user}}
	update := bson.D{{Key: "$set", Value: bson.D{{Key: "team", Value: aux}}}}
	_, err = database.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return err
	}
	return nil
}

// carga el team por defecto que contiene 3 pokemones
func newDefaultTeam() []pokemon {
	var team []pokemon
	pokemons := []string{"charizard", "bulbasaur", "pikachu"}

    // buscar desde la pokeapi
	for _, p := range pokemons {
		pk, _ := getPokemonFromApi(p)
		team = append(team, *pk)
	}
	return team
}
