package main

import "fmt"

const conferenceTickets int = 50
var conferenceName = "Crush Conference" // Syntactic sugar
var remainingTickets uint = 50
var bookings = make([]UserData, 0)

type UserData struct {
	firstName string
	lastNam string
	email string
	userTickets uint
}

func main() {

	greetUsers()

	for {
		
		firstName, lastName, email, userTickets := getUserInformation()

		if isValidateUserInfo(firstName,lastName,email, userTickets) {
			bookTickets(firstName, lastName, email, userTickets)
	
			firstNames := getFirstNames()
			fmt.Printf("The first name of all our bookings: %v\n", firstNames)
	
			if remainingTickets == 0 {
				fmt.Printf("Our conference is booked out. Please come back next year.\n")
				break
			}
		}
		
	}
	
}

func greetUsers()  {
	fmt.Printf("conferenceName is %T, conferenceTickets is %T and remainingTickets is %T\n", conferenceName, conferenceTickets, remainingTickets)
	fmt.Printf("Welcome to %v Booking Application\n", conferenceName)
	fmt.Printf("We have a total of %v conference tickets and %v are still available\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend😋")
}

func getUserInformation() (string, string, string, uint) {

	var firstName string
	var lastName string
	var email string
	var userTickets uint

	fmt.Println("Enter your First Name: ")
	fmt.Scan(&firstName)
		
	fmt.Println("Enter your Last Name: ")
	fmt.Scan(&lastName)
		
	fmt.Println("Enter your Email: ")
	fmt.Scan(&email)
		
	fmt.Println("How many tickets do you want to book: ")
	fmt.Scan(&userTickets)
	
	return firstName, lastName, email, userTickets
}

func getFirstNames() []string {
	firstNames := []string{}
	for _, booking := range bookings {
		firstNames = append(firstNames, booking.firstName)
	}
	return firstNames
}

func bookTickets(firstName string, lastName string, email string, userTickets uint)  {

	var userData = UserData{
		firstName: firstName,
		lastNam: lastName,
		email: email,
		userTickets: userTickets,
	}
 
	bookings = append(bookings, userData)
		
	remainingTickets = remainingTickets - userTickets
	
	fmt.Printf("Thank you, %v %v. We received your order to book %v tickets. You will a confirmation email at %v.\n", firstName, lastName,userTickets, email)
	fmt.Printf("There are %v tickets left at %v conference\n", remainingTickets, conferenceName)
}