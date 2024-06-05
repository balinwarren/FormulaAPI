package main

import (
	"fmt"
	"net/http"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to Formula API!")
	fmt.Println("Endpoint hit: homePage")
}

func handleRequests() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", homePage)

	if err := http.ListenAndServe(":10000", mux); err != nil {
		fmt.Println(err.Error())
	}
}

func main() {
	handleRequests()
}
