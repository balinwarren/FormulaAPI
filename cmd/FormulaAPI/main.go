package main

import (
	"context"
	"fmt"
	"net/http"
	"strconv"

	"github.com/balinwarren/FormulaAPI/internal/json"
	"go.mongodb.org/mongo-driver/bson"
)

func main() {
	//start api
	mux := http.NewServeMux()

	//homePage endpoint
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Welcome to Formula API!")
		fmt.Println("Endpoint hit: homePage")
	})

	//all drivers endpoint
	mux.HandleFunc("GET /drivers", func(w http.ResponseWriter, r *http.Request) {
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
			fmt.Fprintf(w, "%s\n\n", json.ConvertJSON(driver))
		}
	})

	//all drivers by year endpoint
	mux.HandleFunc("GET /drivers/year/{year}", func(w http.ResponseWriter, r *http.Request) {
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
			fmt.Fprintf(w, "%s\n\n", json.ConvertJSON(driver))
		}
	})

	//individual diver information endpoint
	mux.HandleFunc("GET /drivers/name/{lastName}/{firstName}", func(w http.ResponseWriter, r *http.Request) {
		firstName := r.PathValue("firstName")
		lastName := r.PathValue("lastName")

		cursor, err := driverCollection.Find(context.TODO(), bson.M{"firstName": firstName, "lastName": lastName})
		if err != nil {
			panic(err)
		}
		var drivers []bson.M
		if err = cursor.All(context.TODO(), &drivers); err != nil {
			panic(err)
		}
		fmt.Fprintf(w, "Return all drivers with name %v %v\n", firstName, lastName)
		fmt.Printf("Endpoint hit: all drivers with name %v %v\n", firstName, lastName)
		for _, driver := range drivers {
			fmt.Fprintf(w, "%s\n\n", json.ConvertJSON(driver))
		}
	})

	//all driver by last name endpoint
	mux.HandleFunc("GET /drivers/name/{lastName}", func(w http.ResponseWriter, r *http.Request) {
		lastName := r.PathValue("lastName")

		cursor, err := driverCollection.Find(context.TODO(), bson.M{"lastName": lastName})
		if err != nil {
			panic(err)
		}
		var drivers []bson.M
		if err = cursor.All(context.TODO(), &drivers); err != nil {
			panic(err)
		}
		fmt.Fprintf(w, "Return all drivers with last name %v\n", lastName)
		fmt.Printf("Endpoint hit: all drivers with last name %v\n", lastName)
		for _, driver := range drivers {
			fmt.Fprintf(w, "%s\n\n", json.ConvertJSON(driver))
		}
	})

	if err := http.ListenAndServe(":10000", mux); err != nil {
		fmt.Println(err.Error())
	}
}
