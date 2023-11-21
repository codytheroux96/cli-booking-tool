package main

import (
	"fmt"
	"strings"
)

func main() {
	conferenceName := "Go Conference"
	const conferenceTickets int = 50
	var remainingTickets uint = 50
	bookings := []string{}

	fmt.Printf("Welcome to the %v booking application!\n", conferenceName)
	fmt.Printf("Get your tickets here to attend!\n")
	fmt.Printf("We have a total of %v tickets and %v are still available.\n", conferenceTickets, remainingTickets)

	for {
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
	
		if userTickets > remainingTickets {
			fmt.Printf("We only have %v tickets remaining, you cannot buy %v tickets.\n", remainingTickets, userTickets)
			continue
		}

		remainingTickets = remainingTickets - userTickets 
		bookings = append(bookings, firstName + " " + lastName )
		
		fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)
		fmt.Printf("There are %v tickets left!\n", remainingTickets)
	
		firstNames := []string{}
		for _, booking := range bookings {
			var names = strings.Fields(booking)
			firstNames = append(firstNames, names[0])
		}
		fmt.Printf("The first names of our bookings are: %v\n", firstNames)

		if remainingTickets == 0 {
			fmt.Println("Sorry but there are no more tickets left. We hope to see you next year!")
			break
		} 
	}
}