package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {

	http.HandleFunc("/", handleIndex)
	http.HandleFunc("/montly_assigment", handleMonthlyAssignment)
	err := http.ListenAndServe(":3001", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func handleIndex(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "VP-Parking its ALIVE!!! muajajajaja :}")
}

func handleMonthlyAssignment(w http.ResponseWriter, r *http.Request) {

	r.ParseForm() //Parse url parameters and body

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

	// update the current month and year
	if configurations["current_month"] == 12 {
		configurations["current_month"] = 1
	} else {
		configurations["current_month"] = configurations["current_month"] + 1
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

	// Init the assignedDays map by month and valid dates
	assignedDays := InitAssignedDaysMap(configurations["current_month"], configurations["car_slots"])

	// Fill the parking slots with the pointed users array and if necessary the remainingUsers
	assignedDays = assignMonthlyParking(assignedDays, pointedUsers, remainingUsers, configurations["car_slots"])

	// update the index for the next month users list
	configurations["monthly_pointer"] = configurations["monthly_pointer"] + configurations["car_slots"]

	// Send data to the REST endpoint
	jsonString, err := json.Marshal(assignedDays)
	if err != nil {
		log.Fatal("Error: Cannot parse the json response", err)
		return
	}
	fmt.Fprintln(w, string(jsonString))
}

func assignMonthlyParking(assignedDays map[string][]string, pointedUsers []assignment, remainingUsers []assignment, slotsSize int) map[string][]string {
	for key, value := range assignedDays {
		currentSlotsAssigned := 0
		for _, userInfo := range pointedUsers {
			if !ContainsDate(key, userInfo.Days) {
				value[currentSlotsAssigned] = userInfo.Email
				currentSlotsAssigned++
			}
		}
		if currentSlotsAssigned < slotsSize {
			for _, userInfo := range remainingUsers {
				if !ContainsDate(key, userInfo.Days) {
					value[currentSlotsAssigned] = userInfo.Email
					currentSlotsAssigned++
					if currentSlotsAssigned == slotsSize {
						break
					}
				}
			}
		}
	}
	return assignedDays
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
