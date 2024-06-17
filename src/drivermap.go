package main

import (
	"go.mongodb.org/mongo-driver/bson"
)

type Driver struct {
	ID            interface{} `json:"_id"`
	First_Name    interface{} `json:"first_name"`
	Last_Name     interface{} `json:"last_name"`
	Date_Of_Birth interface{} `json:"date_of_birth"`
	Date_Of_Death interface{} `json:"date_of_death"`
	Active        interface{} `json:"active"`
	Car_Number    interface{} `json:"car_number"`
	Nationality   interface{} `json:"nationality"`
	Teams_Raced   interface{} `json:"teams_raced"`
	Years_Active  interface{} `json:"years_active"`
	Race_Starts   interface{} `json:"race_starts"`
	Wdcs          interface{} `json:"wdcs"`
	Wins          interface{} `json:"wins"`
	Podiums       interface{} `json:"podiums"`
	Points        interface{} `json:"points"`
	Unique_Gp_Won interface{} `json:"unique_gp_won"`
}

func reorderDriverMap(dataMap bson.M) Driver {
	m := Driver{
		ID:            dataMap["_id"],
		First_Name:    dataMap["firstName"],
		Last_Name:     dataMap["lastName"],
		Date_Of_Birth: dataMap["dateOfBirth"],
		Date_Of_Death: dataMap["dateOfDeath"],
		Active:        dataMap["active"],
		Car_Number:    dataMap["car_number"],
		Nationality:   dataMap["nationality"],
		Teams_Raced:   dataMap["teamsRaced"],
		Years_Active:  dataMap["yearsActive"],
		Race_Starts:   dataMap["raceStarts"],
		Wdcs:          dataMap["wdcs"],
		Wins:          dataMap["wins"],
		Podiums:       dataMap["podiums"],
		Points:        dataMap["points"],
		Unique_Gp_Won: dataMap["uniqueGpWon"],
	}

	return m
}
