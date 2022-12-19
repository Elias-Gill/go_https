package server

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type pokemonModel struct {
	Name  string `bson:"name"`
	Type  string `bson:"type"`
	Power int    `bson:"power"`
}
type teamModel struct {
	Team []pokemonModel `bson:"team"`
}

func GetTeamFromUser(user string) (*teamModel, error) {
	c := connectToMongo()
	defer c.closeMongo()
	database := c.conn.Database("myFirstDatabase").Collection("teammodels")

	// busqueda
	var team teamModel
	err := database.FindOne(context.TODO(), bson.D{{"userName", user}}).Decode(&team)
	if err == mongo.ErrNoDocuments {
		fmt.Println("No user found")
		return nil, err
	}
    return &team, nil
}

// adds the chossen pokemon to the given username team
func AddNewPokemon(user string, pokName string) error {
	// conectar con mongo
	c := connectToMongo()
	defer c.closeMongo()
	// db y coleccion
	database := c.conn.Database("myFirstDatabase").Collection("teammodels")

    team, err := GetTeamFromUser(user)
	// comprobar que no puedas tener un equipo mayor que 5
	if len(team.Team) == 5 {
		return fmt.Errorf("Se alcanzo el maximo de pokemons en el equipo")
	}

    // actualizar el team y mandarlo a la db
	newPok := pokemonModel{Name: pokName, Type: "agua", Power: 5}
	team.Team = append(team.Team, newPok)
	matches, err := database.UpdateOne(context.TODO(), bson.D{{"userName", user}}, bson.E{"team", team.Team})
	if err != nil || matches.MatchedCount == 0 {
		return err
	}
	return nil
}
