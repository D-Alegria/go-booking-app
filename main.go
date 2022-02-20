package main

import (
	"fmt"
	"strings"
)

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

	var bookings []string

	for {
		fmt.Println("Enter your First Name: ")
		fmt.Scan(&firstName)
		
		fmt.Println("Enter your Last Name: ")
		fmt.Scan(&lastName)
		
		fmt.Println("Enter your Email: ")
		fmt.Scan(&email)
		
		fmt.Println("How many tickets do you want to book: ")
		fmt.Scan(&userTickets)

		isValidName := len(firstName) > 2 && len(lastName) > 2 
		isValidEmail := strings.Contains(email, "@")
		isValidNumberOfTickets := userTickets > 0 && userTickets <= remainingTickets

		if isValidName && isValidEmail && isValidNumberOfTickets {
			bookings = append(bookings, firstName + " " + lastName)
		
			remainingTickets = remainingTickets - userTickets
			
			fmt.Printf("Thank you, %v %v. We received your order to book %v tickets. You will a confirmation email at %v.\n", firstName, lastName,userTickets, email)
			fmt.Printf("There are %v tickets left at %v conference\n", remainingTickets, conferenceName)
	
			firstNames := []string{}
			for _, booking := range bookings {
				firstNames = append(firstNames, strings.Fields(booking)[0])
			}
	
			fmt.Printf("The first name of all our bookings: %v\n", firstNames)
	
			if remainingTickets == 0 {
				fmt.Printf("Our conference is booked out. Please come back next year.\n")
				break
			}
		} else {
			if !isValidName{
				fmt.Printf("Invalid firstname or lastname, please try again.\n")
			}
			
			if !isValidEmail{
				fmt.Printf("Invalid email, please try again.\n")
			}

			if !isValidNumberOfTickets{
				fmt.Printf("Invalid email, please try again.\n")
			}
		}

		
		if userTickets > remainingTickets{
			fmt.Printf("You can't book %v because there are only %v tickets left.", userTickets, remainingTickets)
			continue
		}
		
		
	}
	
}
