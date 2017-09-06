package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUserAssignation(t *testing.T) {
	user1 := assignment{Email: "fake1@gmail.com", Days: []string{}}
	user2 := assignment{Email: "fake2@gmail.com", Days: []string{}}
	user3 := assignment{Email: "fake3@gmail.com", Days: []string{}}

	var allUsers []assignment
	allUsers = append(allUsers, user1, user2, user3)

	userAssignation, index, err := createUserAssignation(allUsers, 3, 0)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, allUsers, userAssignation)
	assert.Equal(t, 3, index)
}

func TestUserAssignationStartingOnNonZeroIndexn(t *testing.T) {
	user1 := assignment{Email: "fake1@gmail.com", Days: []string{}}
	user2 := assignment{Email: "fake2@gmail.com", Days: []string{}}
	user3 := assignment{Email: "fake3@gmail.com", Days: []string{}}

	var allUsers []assignment
	allUsers = append(allUsers, user1, user2, user3)

	userAssignation, index, err := createUserAssignation(allUsers, 3, 2)
	if err != nil {
		t.Fatal(err)
	}
	var expected []assignment
	expected = append(expected, user3, user1, user2)
	assert.Equal(t, expected, userAssignation)
	assert.Equal(t, 2, index)
}

func TestUserAssignationZeroUsers(t *testing.T) {

	var allUsers []assignment

	userAssignation, index, err := createUserAssignation(allUsers, 3, 2)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, []assignment([]assignment(nil)), userAssignation)
	assert.Equal(t, 0, index)
}

func TestUserAssignationMoreSlotsThanUsers(t *testing.T) {

	user1 := assignment{Email: "fake1@gmail.com", Days: []string{}}
	user2 := assignment{Email: "fake2@gmail.com", Days: []string{}}
	user3 := assignment{Email: "fake3@gmail.com", Days: []string{}}

	var allUsers []assignment
	allUsers = append(allUsers, user1, user2, user3)

	userAssignation, index, err := createUserAssignation(allUsers, 10, 0)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, allUsers, userAssignation)
	assert.Equal(t, 0, index)
}
