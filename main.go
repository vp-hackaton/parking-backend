package main

import (
	"log"
	"strconv"
	"strings"
	"time"
)

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

	var confs []configs
	err = firebaseEndpointHandler("configuration", &confs)
	if err != nil {
		log.Fatal("Cannot get configuration info", err)
		return
	}

	configurations := make(map[string]int)
	for _, conf := range confs {
		key, value := parseConfig(conf)
		configurations[key] = value
	}

	// Process the data
	// Send data to the REST endpoint
}

func parseConfig(config configs) (string, int) {
	configArr := strings.Split(config, ":")
	i, err := strconv.Atoi(configArr[1])
	if err != nil {
		log.Fatal("Error: ", err)
	}
	return configArr[0], i
}
