package main

import (
	"fmt"
	"strings"
)

func main() {

	const conferenceName = "Go Conference"

	const ticketCount = 50
	var remainingTickets = 50

	var bookings = []string{}
	fmt.Printf("ticket size is %v\n", ticketCount)

	fmt.Println("Welcome to to the ", conferenceName, "booking application")

	fmt.Printf("Type of the ticketCount is %T", ticketCount)

	for {
		var firstName string
		var lastName string
		var email string
		var tickets int

		fmt.Println("Enter your first Name")
		fmt.Scan(&firstName)

		fmt.Println("Enter your last Name")
		fmt.Scan(&lastName)

		bookings = append(bookings, firstName+" "+lastName)

		fmt.Println("Enter your email address")
		fmt.Scan(&email)

		fmt.Println("Enter the number of tickets you want")
		fmt.Scan(tickets)

		if remainingTickets > tickets {
			remainingTickets = ticketCount - tickets
		}
		bookings = append(bookings, firstName+" "+lastName)

		firstNames := []string{}
		for _, booking := range bookings {
			var names = strings.Fields(booking)
			var fname = names[0]
			firstNames = append(firstNames, fname)
			fmt.Printf("The first names of the booking are: %v\n", firstNames)
		}

	}

}
