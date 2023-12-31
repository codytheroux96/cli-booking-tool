package main

import (
	"fmt"
	//"sync"
	"time"
)

const conferenceTickets int = 50
var conferenceName = "Go Conference"
var remainingTickets uint = 50
var bookings = make([]UserData, 0)

type UserData struct {
	firstName	string
	lastName	string
	email	string
	numberOfTickets uint
}

//var wg = sync.WaitGroup{}

func main() {
	greetUsers()

	for {
		firstName, lastName, email, userTickets := getUserInput()
		
		isValidName, isValidEmail, isValidTicketNumber := validateUserInput(firstName, lastName, email, userTickets, remainingTickets)
	
		if isValidName && isValidEmail && isValidTicketNumber {
			bookTicket(userTickets, firstName, lastName, email)

			//wg.Add(1)
			go sendTicket(userTickets, firstName, lastName, email)

			firstNames := getFirstNames()
			fmt.Printf("The first names of people who have booked so far are: %v\n", firstNames)
	
			if remainingTickets == 0 {
				fmt.Println("Sorry but there are no more tickets left. We hope to see you next year!")
				break
			} 
		} else {
			if !isValidName {
				fmt.Println("The first name or last name you entered is too short!")
			}
			if !isValidEmail {
				fmt.Println("The email address you entered does not contain @ sign!")
			}
			if !isValidTicketNumber {
				fmt.Println("The number of tickets you entered is invalid!")
			}
		}
	}
	//wg.Wait()
}

func greetUsers() {
	fmt.Printf("Welcome to the %v booking application!\n", conferenceName)
	fmt.Printf("Get your tickets here to attend!\n")
	fmt.Printf("We have a total of %v tickets and %v are still available.\n", conferenceTickets, remainingTickets)
}

func getFirstNames() []string {
	firstNames := []string{}
	for _, booking := range bookings {
		firstNames = append(firstNames, booking.firstName)
	}
	return firstNames
}

func getUserInput() (string, string, string, uint) {
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

	return firstName, lastName, email, userTickets
}

func bookTicket(userTickets uint, firstName string, lastName string, email string) {
	remainingTickets = remainingTickets - userTickets 

	var userData = UserData {
		firstName: firstName,
		lastName: lastName,
		email: email,
		numberOfTickets: userTickets,
	}

	bookings = append(bookings, userData)
	
	fmt.Println("-----------------")
	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)
	fmt.Printf("There are %v tickets left!\n", remainingTickets)
}

func sendTicket(userTickets uint, firstName string, lastName string, email string) {
	time.Sleep(50 * time.Second)
	var ticket = fmt.Sprintf("Sent %v tickets for %v %v", userTickets, firstName, lastName)
	fmt.Println("-----------------")
	fmt.Printf("Sending ticket...\n %v to email address %v\n", ticket, email)
	fmt.Println("-----------------")
	//wg.Done()
}