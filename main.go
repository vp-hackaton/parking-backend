package main

import (
	"encoding/json"
	"log"
	"net/http"
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
	userEndpointURL := "https://vpparking-de51c.firebaseio.com/users.json"
	err := endpointHandler(userEndpointURL, &users)
	if err != nil {
		log.Fatal("Cannot get users info", err)
		return
	}

	var confs []configs
	configEndpointURL := "https://vpparking-de51c.firebaseio.com/configuration.json"
	err = endpointHandler(configEndpointURL, &confs)
	if err != nil {
		log.Fatal("Cannot get configuration info", err)
		return
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
