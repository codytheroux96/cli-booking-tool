package main

import "fmt"

func main() {
	conferenceName := "Go Conference"
	const conferenceTickets int = 50
	var remainingTickets uint = 50

	fmt.Printf("Welcome to the %v booking application!\n", conferenceName)
	fmt.Printf("Get your tickets here to attend!\n")
	fmt.Printf("We have a total of %v tickets and %v are still available.\n", conferenceTickets, remainingTickets)

	var bookings []string

	var firstName string
	var lastName string
	var email string
	var userTickets uint

	fmt.Println(("Enter your first name: "))
	fmt.Scan(&firstName)

	fmt.Println(("Enter your last name: "))
	fmt.Scan(&lastName)

	fmt.Println(("Enter your email: "))
	fmt.Scan(&email)

	fmt.Println(("Enter number of tickets: "))
	fmt.Scan(&userTickets)

	remainingTickets = remainingTickets - userTickets 
	bookings = append(bookings, firstName + " " + lastName )
	
	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)
	fmt.Printf("There are %v tickets left!\n", remainingTickets)
}
