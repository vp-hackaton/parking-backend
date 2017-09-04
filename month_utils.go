package main

import (
	"log"
	"time"
)

// WFHToDays returns a slice of WFH days on a given range of dates
func WFHToDays(wfh string, dayOfMonth string) ([]time.Time, error) {
	// set the starting date (in any way you wish)
	start, err := time.Parse("2006-01-02", dayOfMonth)
	// handle error
	if err != nil {
		log.Fatal("Cannot calculate date range", err)
		return nil, err
	}

	var days []time.Time

	// set d to starting date and keep adding 1 day to it as long as month doesn't change
	for d := start; d.Month() == start.Month(); d = d.AddDate(0, 0, 1) {
		weekDay := d.Weekday().String()
		if weekDay == wfh {
			days = append(days, d)
		}
	}
	return days, nil
}
