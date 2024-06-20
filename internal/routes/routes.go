package routes

import (
	"fmt"
	"net/http"
)

func Router() *http.ServeMux {
	mux := http.NewServeMux()

	//endpoint routes
	mux.HandleFunc("/api", homePage)
	mux.HandleFunc("GET /api/drivers", getAllDrivers)
	mux.HandleFunc("GET /api/drivers/year/{year}", getDriversByYear)
	mux.HandleFunc("GET /api/drivers/name/{lastName}/{firstName}", getDriverByFullName)
	mux.HandleFunc("GET /api/drivers/name/{lastName}", getDriversByLastName)
	mux.HandleFunc("GET /api/drivers/wdcs", getAllWDCs)
	mux.HandleFunc("GET /api/drivers/winners", getAllGpWinners)

	return mux
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to Formula API!")
	fmt.Println("Endpoint hit: homePage")
}
