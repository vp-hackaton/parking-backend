package main

import (
	"log"
	"strconv"
	"strings"
)

type vehicle struct {
	Plate  string `json:"plate"`
	Model  int    `json:"model"`
	Color  string `json:"color"`
	Type   string `json:"type"`
	IsMain bool   `json:"is_main"`
}

type configs string

type user struct {
	Email    string    `json:"email"`
	Name     string    `json:"name"`
	Vehicle  []vehicle `json:"vehicles"`
	Wfh      string    `json:"wfh"`
	IsActive bool      `json:"is_active"`
	Password string    `json:"password"`
	Freedays []string  `json:"free_days"`
}

type assignment struct {
	Email string
	Days  []string
}

func (c configs) parse() (string, int) {
	confSlice := strings.Split(string(c), ":")
	i, err := strconv.Atoi(confSlice[1])
	if err != nil {
		log.Fatalf("Error: Cannot parse the config param %s %e", c, err)
	}
	return confSlice[0], i
}
