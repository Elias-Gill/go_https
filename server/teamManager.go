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
	// anadir desde la pokeapi
	pokemon, err := getPokemonFromApi(pokName)
	team = append(team, *pokemon)
	// guardar en la db
	filter := bson.D{{Key: "userName", Value: user}}
	update := bson.D{{Key: "$set", Value: bson.D{{Key: "team", Value: team}}}}
	Database.UpdateOne(context.TODO(), filter, update)
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
	Database.UpdateOne(context.TODO(), filter, update)
	return nil
}

// carga el team por defecto que contiene 3 pokemones
func cargarTeamPorDefecto() []pokemon {
	var team []pokemon
	pk, _ := getPokemonFromApi("charizard")
	team = append(team, *pk)
	pk, _ = getPokemonFromApi("bulbasaur")
	team = append(team, *pk)
	pk, _ = getPokemonFromApi("pikachu")
	team = append(team, *pk)
	return team
}
