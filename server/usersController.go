package server

import (
	"context"
	"fmt"

	"github.com/elias-gill/go_pokemon/tools"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/crypto/bcrypt"
)

// iniciar sesion en el server y enviar un webtoken de autenticacion
func IniciarSesion(nombre string, contrasena string) (string, error) {
	user, err := SearchUser(nombre)
	if err != nil {
		return "", err
	}
	// comparar la contrasena guardada con la contrasena proporcionada por el usuario
	if err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(contrasena)); err != nil {
		return "", err
	}
	// generar y retornar el token con expiracion de 10 mins
	token, err := tools.GenerateJWT(nombre)
	if err != nil {
		return "", err
	}
	return token, nil
}

// funcion para anadir nuevo usuario a la base de datos
func NewUser(nombre string, password string) error {
	// conectar con mongo
	c := connectToMongo()
	defer c.closeMongo()
	// comprobar datos
	if nombre == "" || password == "" {
		return fmt.Errorf("Datos proporcionados invalidos")
	}
	// si el usuario ya existe entoces retornar un error
	if u, _ := SearchUser(nombre); u != nil {
		return fmt.Errorf("El usuario ya existe")
	}

	// encriptar contrasena
	encriptedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	if err != nil {
		println(err.Error())
		return err
	}
	// insertion "query"
	coll := c.conn.Database("myFirstDatabase").Collection("usermodels")
	doc := bson.D{{"userName", nombre}, {"password", encriptedPassword}}
	result, err := coll.InsertOne(context.TODO(), doc)
	if err != nil {
		println(err.Error())
		return err
	}

	fmt.Printf("Inserted user with _id: %v\n", result.InsertedID) // debuging
	return nil
}

/* struct de modelo de usuario */
type userModel struct {
	UserName string `bson:"userName"`
	Id       string `bson:"_id"`
	Password string `bson:"password"`
}

// funcion para buscar un usuario dentro de la base de datos
func SearchUser(user string) (*userModel, error) {
	// conectar con mongo
	c := connectToMongo()
	defer c.closeMongo()
	// db y coleccion
	database := c.conn.Database("myFirstDatabase").Collection("usermodels")

	// busqueda
	var result userModel
	err := database.FindOne(context.TODO(), bson.D{{"userName", user}}).Decode(&result)
	if err == mongo.ErrNoDocuments {
		fmt.Println("No user found")
		return nil, err
	}
	return &result, nil
}

// funcion para cerrar mongo tranquilamente
func (c serverMongo) closeMongo() {
	if err := c.conn.Disconnect(context.TODO()); err != nil {
		panic(err)
	}
}

// servidor conectado
type serverMongo struct {
	conn *mongo.Client
}

// funcion para conectar a la base de datos
func connectToMongo() *serverMongo {
	// TODO: add uri token to a enviroment variable
	uri := "mongodb+srv://elias:elias@cluster0.zrvrq.mongodb.net/myFirstDatabase?retryWrites=true&w=majority"
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
	}
	return &serverMongo{conn: client}
}
