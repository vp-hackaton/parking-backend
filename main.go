package main

import (
	"fmt"
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

	allUsers := make([]assignment, len(users))
	for i, user := range users {
		days, err := FreeDaysToSlice(user.Wfh, time.Now().Format("2006-01-02"))
		if err != nil {
			log.Fatal("Error: ", err)
			return
		}
		assign := assignment{Email: user.Email, Days: append(days, user.Freedays...)}
		allUsers[i] = assign
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

	pointedUsers, err := createUserAssignation(allUsers, configurations["car_slots"], 0)
	if err != nil {
		log.Fatal("Error: ", err)
		return
	}

	remainingUsers, err := createUserAssignation(allUsers, configurations["users_size"]-configurations["car_slots"], configurations["car_slots"])
	if err != nil {
		log.Fatal("Error: ", err)
		return
	}

	// Process the data
	assignedDays := InitAssignedDaysMap(configurations["current_month"], configurations["car_slots"])

	for key, value := range assignedDays {
		currentSlotsAssigned := 0
		for _, userInfo := range pointedUsers {
			if !ContainsDate(key, userInfo.Days) {
				value[currentSlotsAssigned] = userInfo.Email
				currentSlotsAssigned++
			}
		}
		if currentSlotsAssigned < configurations["car_slots"] {
			for _, userInfo := range remainingUsers {
				if !ContainsDate(key, userInfo.Days) {
					value[currentSlotsAssigned] = userInfo.Email
					currentSlotsAssigned++
					if currentSlotsAssigned == configurations["car_slots"] {
						break
					}
				}
			}
		}
	}

	fmt.Println(assignedDays)
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

func createUserAssignation(allUsers []assignment, numberOfSlots int, currentPos int) ([]assignment, error) {
	i := currentPos
	var assigmentUsers []assignment
	for i < numberOfSlots {
		actualValue := (allUsers)[i]
		assigmentUsers = append(assigmentUsers, actualValue)
		allUsers = append((allUsers)[:i], (allUsers)[i+1:]...)
		i++
	}
	return assigmentUsers, nil
}
