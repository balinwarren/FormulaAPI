package data

import (
	"context"
	"os"

	//"github.com/gofor-little/env"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func getClient() (*mongo.Client, error) {
	//load env
	//if err := env.Load(".env"); err != nil {
	//	panic(err)
	//}

	//dbUri := env.Get("DB_URI", "")
	dbUri := os.Getenv("DB_URI")

	//mongo connection
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(dbUri).SetServerAPIOptions(serverAPI)

	client, err := mongo.Connect(context.TODO(), opts)

	if err != nil {
		panic(err)
	}

	var result bson.M
	if err := client.Database("admin").RunCommand(context.TODO(), bson.D{{Key: "ping", Value: 1}}).Decode(&result); err != nil {
		panic(err)
	}

	return client, err
}

func GetCollection(collectionName string) (*mongo.Collection, *mongo.Client, error) {
	client, err := getClient()
	return client.Database("formulaone").Collection(collectionName), client, err
}

func CloseConnection(client *mongo.Client, err error) {
	if err != nil {
		panic(err)
	}

	if err = client.Disconnect(context.TODO()); err != nil {
		panic(err)
	}
}
