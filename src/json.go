package main

import (
	"encoding/json"

	"go.mongodb.org/mongo-driver/bson"
)

func convertJSON(dataMap bson.M) string {
	data := reorderDriverMap(dataMap)
	jsonBytes, err := json.MarshalIndent(data, "", "   ")
	if err != nil {
		panic(err)
	}

	return string(jsonBytes)
}
