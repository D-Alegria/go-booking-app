package main

import "fmt"

func main() {
	conferenceName := "Crush Conference" // Syntactic sugar
	const conferenceTickets = 50
	var remainingTickets = 50

	fmt.Printf("conferenceName is %T, conferenceTickets is %T and remainingTickets is %T\n", conferenceName, conferenceTickets, remainingTickets)

	fmt.Printf("Welcome to %v Booking Application\n", conferenceName)
	fmt.Printf("We have a total of %v conference tickets and %v are still available\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend😋")

	var userName string
	var userTickets int

	userName = "Demi"
	userTickets = 4

	fmt.Printf("User %v has %v tickets\n", userName, userTickets)
}
