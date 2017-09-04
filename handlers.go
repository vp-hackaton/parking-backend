package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

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

func firebaseEndpointHandler(endpointName string, records interface{}) error {
	firebaseRootURL := "https://vpparking-de51c.firebaseio.com/%s.json"
	endpointURL := fmt.Sprintf(firebaseRootURL, endpointName)
	return endpointHandler(endpointURL, records)
}
