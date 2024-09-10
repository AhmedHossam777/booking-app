package main

import "fmt"

func main() {
	var conferenceName string = "Go conference"
	const conferenceTickets int = 50
	var remainingTickets uint = 50

	fmt.Printf("Welcome to our %v booking application \n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still available \n",
		conferenceTickets, remainingTickets,
	)
	fmt.Println("Get your ticket here to attend")

	var firstname string
	var lastname string
	var email string
	var userTickets uint

	//ask users for their name
	fmt.Println("Enter your firstname")
	fmt.Scan(&firstname)

	fmt.Println("Enter your lastname")
	fmt.Scan(&lastname)

	fmt.Println("Enter you email")
	fmt.Scan(&email)

	fmt.Println("Enter the number of tickets you want to book")
	fmt.Scan(&userTickets)

	remainingTickets = remainingTickets - userTickets

	fmt.Printf("User %v %v with Email: %v booked %v ticket.\n", firstname,
		lastname, email, userTickets,
	)

	fmt.Printf("the remaining tickets %v for %v \n", remainingTickets,
		conferenceName,
	)
}