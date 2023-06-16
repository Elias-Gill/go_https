package server

import (
	"context"
	"fmt"

	"github.com/elias-gill/go_pokemon/authentication"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

var defaultTeam = newDefaultTeam()

/* struct de modelo de usuario */
type UserModel struct {
	UserName string    `bson:"userName"`
	Id       string    `bson:"_id"`
	Password string    `bson:"password"`
	Team     []pokemon `bson:"team"` // array de pokemones
}

// iniciar sesion en el server y enviar un webtoken de autenticacion
func IniciarSesion(nombre string, contrasena string) (string, error) {
	user, err := SearchUserInfo(nombre)
	if err != nil {
		return "", err
	}
	// comparar la contrasena guardada con la contrasena proporcionada por el usuario
	if err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(contrasena)); err != nil {
		return "", err
	}
	// generar y retornar el token con expiracion de 10 mins
	token, err := authentication.GenerateJWT(nombre)
	if err != nil {
		return "", err
	}
	return token, nil
}

// funcion para anadir nuevo usuario a la base de datos
func NewUser(nombre string, password string) error {
	// comprobar datos
	if nombre == "" || password == "" {
		return fmt.Errorf("Datos proporcionados invalidos")
	}

	// si el usuario ya existe entoces retornar un error
	if u, _ := SearchUserInfo(nombre); u != nil {
		return fmt.Errorf("El usuario ya existe")
	}

	// encriptar contrasena
	encriptedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	if err != nil {
		println(err.Error())
		return err
	}

	// insertion "query"
	doc := bson.D{
		{Key: "userName", Value: nombre},
		{Key: "password", Value: encriptedPassword},
		{Key: "team", Value: defaultTeam},
	}
	result, err := database.InsertOne(context.TODO(), doc)
	if err != nil {
		println(err.Error())
		return err
	}

	fmt.Printf("Inserted user with _id: %v\n", result.InsertedID) // debuging
	return nil
}

// funcion para eliminar un usuario de la base de datos
func DeleteUser(user string) error {
	filter := bson.D{{Key: "userName", Value: user}}
	_, err := database.DeleteOne(context.TODO(), filter)
	return err
}

// funcion para buscar un usuario dentro de la base de datos
func SearchUserInfo(user string) (*UserModel, error) {
	// busqueda
	var result UserModel
	filter := bson.D{{Key: "userName", Value: user}}
	err := database.FindOne(context.TODO(), filter).Decode(&result)
	if err == mongo.ErrNoDocuments {
		fmt.Println("No user found")
		return nil, err
	}
	return &result, nil
}
