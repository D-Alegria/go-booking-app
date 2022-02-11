package main

import "fmt"

func main() {
	var conferenceName = "Crush Conference"
	const conferenceTickets = 50
	var remainingTickets = 50
	fmt.Printf("Welcome to %v Booking Application\n", conferenceName)
	fmt.Printf("We have a total of %v conference tickets and %v are still available\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend😋")
}
