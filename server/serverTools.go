package server

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// conectar con mongo
var (
	// conexion a mongoDB
	C = ConnectToMongo()
	// db y coleccion
	Database = C.conn.Database("myFirstDatabase").Collection("users")
)

// servidor conectado
type serverMongo struct {
	conn *mongo.Client
}

// funcion para conectar a la base de datos
func ConnectToMongo() *serverMongo {
	// TODO: add uri token to a enviroment variable
	uri := "mongodb+srv://elias:elias@cluster0.zrvrq.mongodb.net/myFirstDatabase?retryWrites=true&w=majority"
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
	}
	return &serverMongo{conn: client}
}

// funcion para cerrar la conexion con mongo tranquilamente
func (c serverMongo) CloseMongo() {
	if err := c.conn.Disconnect(context.TODO()); err != nil {
		panic(err)
	}
}
