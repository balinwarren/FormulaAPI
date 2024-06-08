package main

import (
	"context"
	"fmt"
	"strconv"
	"strings"

	"github.com/gofor-little/env"
	"github.com/xuri/excelize/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type driver struct {
	firstName   string
	lastName    string
	dateOfBirth string
	dateOfDeath string
	nationality string
	teamsRaced  []string
	wdc         int
	wins        int
	podiums     int
	points      float32
	uniqueGpWon []string
}

func main() {
	//load env
	if err := env.Load("../.env"); err != nil {
		panic(err)
	}

	dbUri := env.Get("DB_URI", "")

	//mongo connection
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(dbUri).SetServerAPIOptions(serverAPI)

	client, err := mongo.Connect(context.TODO(), opts)

	if err != nil {
		panic(err)
	}
	defer func() {
		if err = client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()

	var result bson.M
	if err := client.Database("admin").RunCommand(context.TODO(), bson.D{{"ping", 1}}).Decode(&result); err != nil {
		panic(err)
	}
	fmt.Println("Pinged your deployment. You successfully connected to MongoDB!")

	//read spreadsheet
	f, err := excelize.OpenFile("Formula 1 Grid.xlsx")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer func() {
		if err := f.Close(); err != nil {
			fmt.Println(err)
		}
	}()

	rows, err := f.GetRows("Drivers")
	if err != nil {
		fmt.Println(err)
		return
	}
	for i, row := range rows {
		var tempDriver = driver{}
		if i == 0 {
			continue
		} else {

			for col, colCell := range row {
				//fmt.Print(col, " ", colCell, "\t")
				switch col {
				case 0:
					split := strings.Split(colCell, " ")
					length := len(split)
					tempDriver.firstName = split[0]
					tempDriver.lastName = strings.Join(split[1:length], " ")
				case 1:
					tempDriver.dateOfBirth = colCell
				case 2:
					tempDriver.dateOfDeath = colCell
				case 3:
					tempDriver.nationality = colCell
				case 4:
					tempDriver.teamsRaced = strings.Split(colCell, ",")
				case 5:
					i, err := strconv.Atoi(colCell)
					if err != nil {
						panic(err)
					}
					tempDriver.wdc = i
				case 6:
					i, err := strconv.Atoi(colCell)
					if err != nil {
						panic(err)
					}
					tempDriver.wins = i
				case 7:
					i, err := strconv.Atoi(colCell)
					if err != nil {
						panic(err)
					}
					tempDriver.podiums = i
				case 8:
					i, err := strconv.ParseFloat(colCell, 32)
					if err != nil {
						panic(err)
					}
					f32 := float32(i)
					tempDriver.points = f32
				case 9:
					if colCell == "0" {
						continue
					} else {
						gps := strings.Split(colCell, ",")
						tempDriver.uniqueGpWon = gps
					}
				}
			}
		}
		fmt.Print(tempDriver)
		// TO DO //
		// ADD ALL DRIVER STRUCTS AS DOCUMENTS TO MONGO DB //
		fmt.Println()
	}

	// TO DO //
	// ADD FUNCTION FOR CIRCUIT INFORMATION //

	// TO DO //
	// ADD FUNCTION FOR CONSTRUCTOR INFORMATION //

	f.Close()
}
