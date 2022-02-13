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
	var userTickets uint

	var bookings [50]string
	
	fmt.Println("Enter your First Name: ")
	fmt.Scan(&firstName)
	
	fmt.Println("Enter your Last Name: ")
	fmt.Scan(&lastName)
	
	fmt.Println("Enter your Email: ")
	fmt.Scan(&email)
	
	fmt.Println("How many tickets do you want to book: ")
	fmt.Scan(&userTickets)
	
	bookings[0] = firstName + " " + lastName
	fmt.Printf("The whole array %v\n", bookings)
	fmt.Printf("The first element %v\n", bookings[0])
	fmt.Printf("The type %T\n", bookings)
	fmt.Printf("The length %v\n", len(bookings))
	
	remainingTickets = remainingTickets - userTickets
	
	fmt.Printf("Thank you, %v %v. We received your order to book %v tickets. You will a confirmation email at %v.\n", firstName, lastName, email, userTickets)
	fmt.Printf("There are %v tickets left at %v conference\n", remainingTickets, conferenceName)
}
