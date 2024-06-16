package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gofor-little/env"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	//load env
	if err := env.Load("src/.env"); err != nil {
		panic(err)
	}

	dbUri := env.Get("DB_URI", "")

	//mongo connection
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(dbUri).SetServerAPIOptions(serverAPI)

	client, err := mongo.Connect(context.TODO(), opts)

	if err != nil {
		panic(err)
	}
	defer func() {
		if err = client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()

	var result bson.M
	if err := client.Database("admin").RunCommand(context.TODO(), bson.D{{"ping", 1}}).Decode(&result); err != nil {
		panic(err)
	}
	fmt.Println("Pinged your deployment. You successfully connected to MongoDB!")

	driverCollection := client.Database("formulaone").Collection("drivers")

	//start api
	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Welcome to Formula API!")
		fmt.Println("Endpoint hit: homePage")
	})

	mux.HandleFunc("GET /driver", func(w http.ResponseWriter, r *http.Request) {
		cursor, err := driverCollection.Find(context.TODO(), bson.M{})
		if err != nil {
			panic(err)
		}
		var drivers []bson.M
		if err = cursor.All(context.TODO(), &drivers); err != nil {
			panic(err)
		}
		fmt.Fprintf(w, "Return all drivers\n")
		fmt.Println("Endpoint hit: all drivers")
		for _, driver := range drivers {
			fmt.Fprintf(w, "%v\n\n", driver)
		}
	})

	if err := http.ListenAndServe(":10000", mux); err != nil {
		fmt.Println(err.Error())
	}
}
