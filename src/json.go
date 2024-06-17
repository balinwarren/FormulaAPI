package main

import (
	"encoding/json"

	"go.mongodb.org/mongo-driver/bson"
)

func convertJSON(dataMap bson.M) string {
	jsonBytes, err := json.MarshalIndent(dataMap, "", "   ")
	if err != nil {
		panic(err)
	}

	return string(jsonBytes)
}
