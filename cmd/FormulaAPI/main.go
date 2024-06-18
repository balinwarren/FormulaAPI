package main

import (
	"fmt"
	"net/http"

	"github.com/balinwarren/FormulaAPI/internal/routes"
)

func main() {
	//start api
	mux := routes.Router()

	if err := http.ListenAndServe(":10000", mux); err != nil {
		fmt.Println(err.Error())
	}
}
