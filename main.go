package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"
)

var endpointURL string

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
	userEndpointURL := "https://vpparking-de51c.firebaseio.com/users.json"

	var records []user

	endpointHandler(userEndpointURL, &records)

	for _, record := range records {
		fmt.Println(record)
		for _, freeday := range record.Freedays {
			fmt.Println(getFixedTime(freeday))
		}
	}

	// Process the data
	// Send data to the REST endpoint
}

func endpointHandler(endpointURL string, records interface{}) error {
	// Build the request
	req, err := http.NewRequest("GET", endpointURL, nil)
	if err != nil {
		log.Fatal("NewRequest: ", err)
		return err
	}

	// For control over HTTP client headers,
	// redirect policy, and other settings,
	// create a Client
	// A Client is an HTTP client
	client := &http.Client{}

	// Send the request via a client
	// Do sends an HTTP request and
	// returns an HTTP response
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal("Do: ", err)
		return err
	}

	// Callers should close resp.Body
	// when done reading from it
	// Defer the closing of the body
	defer resp.Body.Close()

	// Fill the record with the data from the JSON

	// Use json.Decode for reading streams of JSON data
	if err := json.NewDecoder(resp.Body).Decode(&records); err != nil {
		log.Println(err)
	}
	return err
}
