package main

import "fmt"

func main() {
	const conferenceName = "Go Conference"
	const conferenceTickets = 50
	var remainingTickets = 50

	fmt.Println("Welcome to the", conferenceName, "booking application!")
	fmt.Println("Get your tickets here to attend!")
	fmt.Println("We have a total of", conferenceTickets, "and", remainingTickets, "are still available.")



}