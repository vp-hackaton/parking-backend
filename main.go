package main

import (
	"encoding/json"
	"fmt"
	"log"
	"time"
)

func main() {

	// Read data from the Firebase REST endpoints

	var users []user

	if err := getUsersFromFB(&users); err != nil {
		log.Fatal(err)
		return
	}

	var allUsers []assignment
	for _, user := range users {
		days, err := FreeDaysToSlice(user.Wfh, time.Now().Format("2006-01-02"))
		if err != nil {
			log.Fatal("Error: ", err)
			return
		}
		assign := assignment{Email: user.Email, Days: append(days, user.Freedays...)}
		allUsers = append(allUsers, assign)
	}

	var configurations configs
	if err := getConfigsFromFB(&configurations); err != nil {
		log.Fatal(err)
		return
	}

	pointedUsers, index, err := createUserAssignation(allUsers, configurations["car_slots"], 0)
	if err != nil {
		log.Fatal("Error: ", err)
		return
	}

	remainingUsers, index, err := createUserAssignation(allUsers, configurations["users_size"]-configurations["car_slots"], index)
	if err != nil {
		log.Fatal("Error: ", err)
		return
	}

	// Update index in the configurations to know the current position for the next assignation
	fmt.Println("Last index assigned, ", index)

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
	// Send data to the REST endpoint
	jsonString, err := json.Marshal(assignedDays)
	if err != nil {
		log.Fatal("Error: Cannot parse the json response", err)
		return
	}
	fmt.Println(string(jsonString))
}

func createUserAssignation(users []assignment, slots int, startAt int) ([]assignment, int, error) {
	if len(users) == 0 {
		return users, 0, nil
	}
	if slots > len(users) {
		return users, 0, nil
	}
	i := startAt
	count := 0
	var result []assignment
	for count < slots {
		if i == slots {
			i = 0
		}
		actualValue := (users)[i]
		result = append(result, actualValue)
		i++
		count++
	}
	return result, i, nil
}
