package main

import (
	"fmt"
	"log"

	"github.com/zabawaba99/firego"
)

var fbEndPointBase string

func init() {
	fbEndPointBase = "https://vpparking-de51c.firebaseio.com/%s"
}

func getUsersFromFB(v *[]user) error {
	fb := firego.New(fmt.Sprintf(fbEndPointBase, "users"), nil)

	if err := fb.Value(&v); err != nil {
		log.Println("There was an error trying to get the users from firebase")
		return err
	}
	return nil
}

func getConfigsFromFB(v *configs) error {
	fb := firego.New(fmt.Sprintf(fbEndPointBase, "configuration"), nil)
	if err := fb.Value(&v); err != nil {
		log.Println("There was an error trying to get the configuration from firebase")
		return err
	}
	return nil
}
