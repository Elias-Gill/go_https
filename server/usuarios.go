package server

import (
	"errors"
	"net/mail"
	// "golang.org/x/crypto/bcrypt"

	"context"
	"encoding/json"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func IniciarSesion(usuario string, contrasena string) (string, error) {
	// TODO: conectar con mongoDB
	if usuario == "elias" && contrasena == "123" {
		return "123", nil
	}
	return "", errors.New("Usuario o contrasena invalido")
}

// funcion para anadir nuevo usuario a la base de datos
func NewUser(nombre string, contrasena string, correo mail.Address) {

}

// funcion para cerrar mongo tranquilamente
func (c ServerMongo) closeMongo() {
	if err := c.conn.Disconnect(context.TODO()); err != nil {
		panic(err)
	}
}

// servidor conectado
type ServerMongo struct {
	conn *mongo.Client
}

// funcion para conectar a la base de datos
func connectToMongo() *ServerMongo {
	// TODO: add uri token to a enviroment variable
	uri := "mongodb+srv://elias:0105Elias@cluster0.zrvrq.mongodb.net/?retryWrites=true&w=majority"
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
	}
	return &ServerMongo{conn: client}
}

// funcion para buscar un valor dentrto de la base de datos
func SearchMongo(c *ServerMongo) []byte{
	// db
	database := c.conn.Database("sample_mflix")
	// coleccion
	coll := database.Collection("movies")
	// cosa a buscar
	title := "Back to the Future"

	var result bson.M
    err := coll.FindOne(context.TODO(), bson.D{{"title", title}}).Decode(&result)
	if err == mongo.ErrNoDocuments {
		fmt.Printf("No document was found with the title %s\n", title)
		return nil
	}
	if err != nil {
		panic(err)
	}

	jsonData, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		panic(err)
	}
    return jsonData
}
