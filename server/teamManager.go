package server

type teamModel struct {
	Team []string `bson:"team"`
}

// adds the chossen pokemon to the given username team
func AddNewPokemon(user string, pokName string) error {
    return nil
}
