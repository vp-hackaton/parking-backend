package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"
)

func getFixedTime(dateToFormat string) time.Time {
	splited := strings.Split(dateToFormat, "-")
	loc, _ := time.LoadLocation("UTC")
	year, err := strconv.Atoi(splited[2])
	if err != nil {
		log.Fatal("Cannot convert year")
	}
	month, err := strconv.Atoi(splited[1])
	if err != nil {
		log.Fatal("Cannot convert month")
	}
	day, err := strconv.Atoi(splited[0])
	if err != nil {
		log.Fatal("Cannot convert day")
	}
	return time.Date(year, time.Month(month), day, 0, 0, 0, 0, loc)
}

func main() {

	// Read data from the Firebase REST endpoints

	var users []user
	err := firebaseEndpointHandler("users", &users)
	if err != nil {
		log.Fatal("Cannot get users info", err)
		return
	}

	days, err := WFHToDays("Wednesday", "2017-09-01")
	if err != nil {
		log.Fatal("Error: ", err)
		return
	}

	fmt.Println("Day: ", days)

	var confs []configs
	err = firebaseEndpointHandler("configuration", &confs)
	if err != nil {
		log.Fatal("Cannot get configuration info", err)
		return
	}

	// Process the data
	// Send data to the REST endpoint
}
