package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/balinwarren/FormulaAPI/internal/data"
	"github.com/balinwarren/FormulaAPI/internal/routes"
)

func main() {
	//start api
	port := os.Getenv("PORT")
	mux := routes.Router()

	//connect DB
	data.Client, data.ClientErr = data.GetClient()

	if err := http.ListenAndServe("0.0.0.0:"+port, mux); err != nil {
		fmt.Println(err.Error())
	}

	//disconnect DB
	data.CloseConnection(data.Client, data.ClientErr)
}
