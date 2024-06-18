package routes

import (
	"fmt"
	"net/http"
)

func Router() *http.ServeMux {
	mux := http.NewServeMux()

	//endpoint routes
	mux.HandleFunc("/", homePage)
	mux.HandleFunc("GET /drivers", getAllDrivers)
	mux.HandleFunc("GET /drivers/year/{year}", getDriversByYear)
	mux.HandleFunc("GET /drivers/name/{lastName}/{firstName}", getDriverByFullName)
	mux.HandleFunc("GET /drivers/name/{lastName}", getDriversByLastName)
	mux.HandleFunc("GET /drivers/wdcs", getAllWDCs)

	return mux
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to Formula API!")
	fmt.Println("Endpoint hit: homePage")
}
