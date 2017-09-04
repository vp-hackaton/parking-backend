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
		log.Fatal("Cannot convert year", err)
	}
	month, err := strconv.Atoi(splited[1])
	if err != nil {
		log.Fatal("Cannot convert month", err)
	}
	day, err := strconv.Atoi(splited[0])
	if err != nil {
		log.Fatal("Cannot convert day", err)
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

	allUsers := make(map[string][]string)
	for _, user := range users {
		days, err := FreeDaysToSlice(user.Wfh, time.Now().Format("2006-01-02"))
		if err != nil {
			log.Fatal("Error: ", err)
			return
		}
		allUsers[user.Email] = days
		allUsers[user.Email] = append(days, user.Freedays...)
	}

	for k, v := range allUsers {
		fmt.Println("Usero: ", k, v)
	}

	var confs []configs
	err = firebaseEndpointHandler("configuration", &confs)
	if err != nil {
		log.Fatal("Cannot get configuration info", err)
		return
	}

	// Process the data
	// Send data to the REST endpoint
}
