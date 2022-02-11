package main

import "fmt"

func main() {
	conferenceName := "Crush Conference" // Syntactic sugar
	const conferenceTickets int = 50
	var remainingTickets uint= 50

	fmt.Printf("conferenceName is %T, conferenceTickets is %T and remainingTickets is %T\n", conferenceName, conferenceTickets, remainingTickets)

	fmt.Printf("Welcome to %v Booking Application\n", conferenceName)
	fmt.Printf("We have a total of %v conference tickets and %v are still available\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend😋")

	var firstName string
	var lastName string
	var email string
	var userTickets int

	userTickets = 4

	fmt.Println("Enter your First Name: ")
	fmt.Scan(&firstName)
	
	fmt.Println("Enter your Last Name: ")
	fmt.Scan(&lastName)
	
	fmt.Println("Enter your Email: ")
	fmt.Scan(&email)
	
	fmt.Println("How many tickets do you want to book: ")
	fmt.Scan(&userTickets)

	fmt.Printf("Thank you, %v %v. We received your order to book %v tickets. You will a confirmation email at %v.\n", firstName, lastName, email, userTickets)
}
