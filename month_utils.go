package main

import (
	"log"
	"strconv"
	"strings"
	"time"
)

// FreeDaysToSlice returns a slice of WFH days on a given range of dates
func FreeDaysToSlice(wfh string, dayOfMonth string) ([]string, error) {
	// set the starting date (in any way you wish)
	start, err := time.Parse("2006-01-02", dayOfMonth)
	// handle error
	if err != nil {
		log.Fatal("Cannot calculate date range", err)
		return nil, err
	}

	var days []string

	// set d to starting date and keep adding 1 day to it as long as month doesn't change
	for d := start; d.Month() == start.Month(); d = d.AddDate(0, 0, 1) {
		weekDay := d.Weekday().String()
		if weekDay == wfh {
			days = append(days, d.Format("2006-01-02"))
		}
	}
	return days, nil
}

// ContainsDate check if the given date exists on the slice
func ContainsDate(dateToFind string, datesSlice []string) bool {
	for _, date := range datesSlice {
		if date == dateToFind {
			return true
		}
	}
	return false
}

// IsWorkDay validates if a given date is a work day
func IsWorkDay(monthBase int, dayNumber int) bool {
	loc, _ := time.LoadLocation("UTC")
	nameDay := time.Date(time.Now().Local().Year(), time.Month(monthBase), dayNumber, 0, 0, 0, 0, loc).Weekday().String()
	if nameDay == "Saturday" || nameDay == "Sunday" {
		return false
	}
	return true

}

// LastDayOfMonth returns the last day number of a given month
func LastDayOfMonth(monthBase int) int {
	loc, _ := time.LoadLocation("UTC")
	return time.Date(time.Now().Local().Year(), time.Month(monthBase), 1, 0, 0, 0, 0, loc).AddDate(0, 1, -1).Day()
}

// InitAssignedDaysMap set the initial dates map to be assigned for a given month, users arrays starts empty
func InitAssignedDaysMap(monthToLoad int, slotSize int) map[string][]string {
	lastDay := LastDayOfMonth(monthToLoad)
	assignedDays := make(map[string][]string)
	for i := 1; i <= lastDay; i++ {
		if IsWorkDay(monthToLoad, i) {
			assignedDays[DayFullString(monthToLoad, i)] = make([]string, slotSize)
		}
	}
	return assignedDays
}

// DayFullString return the full formated date for a given day, it adds the 0 digit before < 10 day numbers. Example: 2017-11-09
func DayFullString(monthToLoad int, dayNumber int) string {
	if dayNumber < 10 {
		return strings.Join([]string{YearMonthString(monthToLoad), "-0", strconv.Itoa(dayNumber)}, "")
	}
	return strings.Join([]string{YearMonthString(monthToLoad), "-", strconv.Itoa(dayNumber)}, "")
}

// YearMonthString return a year-month concat, it adds the 0 digit before < 10 months numbers. Example: 2017-01
func YearMonthString(monthToLoad int) string {
	if monthToLoad < 10 {
		return strings.Join([]string{strconv.Itoa(time.Now().Local().Year()), "-0", strconv.Itoa(monthToLoad)}, "")
	}
	return strings.Join([]string{strconv.Itoa(time.Now().Local().Year()), "-", strconv.Itoa(monthToLoad)}, "")
}
