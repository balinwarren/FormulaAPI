package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gofor-little/env"
	"go.mongodb.org/mongo-driver/bson"

	//"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

/*type Driver struct {
	_id         primitive.ObjectID `json:"_id"`
	firstName   string	`json:"first_name"`
	lastName    string	`json:"last_name"`
	dateOfBirth string	`json:"date_of_birth"`
	dateofDeath string	`json:"date_of_death"`
	nationality string
	teamsRaced  []string
	yearsActive []int
	raceStarts  int
	wdc         int
	wins        int
	podiums     int
	points      int
	uniqueGpWon []string
	active      bool
	carNumber   string
}*/

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

	//homePage endpoint
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Welcome to Formula API!")
		fmt.Println("Endpoint hit: homePage")
	})

	//all drivers endpoint
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
			jsonBytes, err := json.MarshalIndent(driver, "", "   ")
			if err != nil {
				panic(err)
			}

			fmt.Fprintf(w, "%s\n\n", string(jsonBytes))
		}
	})

	//all drivers by year endpoint
	mux.HandleFunc("GET /{year}/drivers", func(w http.ResponseWriter, r *http.Request) {
		year, err := strconv.Atoi(r.PathValue("year"))
		if err != nil {
			panic(err)
		}

		cursor, err := driverCollection.Find(context.TODO(), bson.M{"yearsActive": year})
		if err != nil {
			panic(err)
		}
		var drivers []bson.M
		if err = cursor.All(context.TODO(), &drivers); err != nil {
			panic(err)
		}
		fmt.Fprintf(w, "Return all drivers active in year %v\n", year)
		fmt.Printf("Endpoint hit: all drivers in year %v\n", year)
		for _, driver := range drivers {
			jsonBytes, err := json.MarshalIndent(driver, "", "   ")
			if err != nil {
				panic(err)
			}

			fmt.Fprintf(w, "%s\n\n", string(jsonBytes))
		}
	})

	if err := http.ListenAndServe(":10000", mux); err != nil {
		fmt.Println(err.Error())
	}
}
