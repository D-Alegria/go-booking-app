package main

import (
	"fmt"
	"strings"
)

func isValidateUserInfo(firstName string, lastName string, email string, userTickets uint) bool {
	isValidName := len(firstName) > 2 && len(lastName) > 2
	isValidEmail := strings.Contains(email, "@")
	isValidNumberOfTickets := userTickets > 0 && userTickets <= remainingTickets

	if !isValidName {
		fmt.Printf("Invalid firstname or lastname, please try again.\n")
	}

	if !isValidEmail {
		fmt.Printf("Invalid email, please try again.\n")
	}

	if !isValidNumberOfTickets {
		fmt.Printf("Invalid number of tickets, please try again.\n")
	}

	return isValidName && isValidEmail && isValidNumberOfTickets
}
