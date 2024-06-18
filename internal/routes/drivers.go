package routes

import (
	"context"
	"fmt"
	"net/http"
	"strconv"

	"github.com/balinwarren/FormulaAPI/internal/json"
	"go.mongodb.org/mongo-driver/bson"
)

func getAllDrivers(w http.ResponseWriter, r *http.Request) {
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
}

func getDriversByYear(w http.ResponseWriter, r *http.Request) {
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
}

func getDriverByFullName(w http.ResponseWriter, r *http.Request) {
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
}

func getDriversByLastName(w http.ResponseWriter, r *http.Request) {
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
}
