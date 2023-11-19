package main

import "fmt"

func main() {
	const conferenceName = "Go Conference"
	const conferenceTickets = 50
	var remainingTickets = 50

	fmt.Printf("Welcome to the %v booking application!\n", conferenceName)
	fmt.Printf("Get your tickets here to attend!\n")
	fmt.Printf("We have a total of %v tickets and %v are still available.", conferenceTickets, remainingTickets)
}
