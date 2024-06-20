package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/balinwarren/FormulaAPI/internal/routes"
)

func main() {
	//start api
	port := os.Getenv("PORT")
	mux := routes.Router()

	if err := http.ListenAndServe("0.0.0.0"+port, mux); err != nil {
		fmt.Println(err.Error())
	}
}
