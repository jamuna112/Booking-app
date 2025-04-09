package main

import (
	"fmt"
)

// Writing a program to accept first name, last name and number of bookings.
func main() {

	const TOTALNUMBEROFTICKETS uint = 20

	var (
		firstName        string
		lastName         string
		email            string
		remainingTickets uint
		bookedTickets    uint
	)
	fmt.Println("===============================================================")
	fmt.Printf("WELCOME TO BOOKING APP\n")
	fmt.Println("===============================================================")

	fmt.Printf("Total tickets available %v\n", TOTALNUMBEROFTICKETS)
	fmt.Printf("Enter firstname: ")
	fmt.Scanf("%v", &firstName)

	fmt.Printf("Enter lastname: ")
	fmt.Scanf("%v\n", &lastName)

	fmt.Printf("Enter email: ")
	fmt.Scanf("%v\n", &email)

	fmt.Printf("Enter number of tickets: ")
	fmt.Scanf("%v\n", &bookedTickets)

	/*
		if booked tickets are greater than total number of tickets, ask user to re-enter within the range
		if all tickets are booked, after accepting email inform user no tickets available
	*/

	for bookedTickets > TOTALNUMBEROFTICKETS {
		fmt.Printf("Please re-enter the tickets within range 1 - %v\n", TOTALNUMBEROFTICKETS)
		fmt.Printf("Enter number of tickets: ")
		fmt.Scanf("%v\n", &bookedTickets)
	}
	remainingTickets = TOTALNUMBEROFTICKETS - bookedTickets

	/*
		user should enter valid number of tickets, non-zero or non-negative
		program should display there are no remaining tickets if user booked all available tickets
		how to run program for multiple users - plan
		Also, read about fmt.Scanf
	*/
	//  else if remainingTickets == 0 {
	// 	fmt.Printf("no tickets availabe\n")
	// }
	/*
		if user has successfully booked the tickets, send him confirmation mail
		if no tickets available send him update, we will inform you next booking slot in email
	*/

	fmt.Println("===============================================================")
	fmt.Printf("WELCOME TO BOOKING APP, %v\n", firstName)
	fmt.Println("===============================================================")
	fmt.Printf("Thanks for booking %v tickets.\nconfirmation will be send to your email at %v\n", bookedTickets, email)
	fmt.Printf("Remaining tickets: %v\n", remainingTickets)
	// fmt.Printf("First name: %v\n", firstName)
	// fmt.Printf("Last name: %v\n", lastName)
	// fmt.Printf("Number of bookings: %v\n", bookedTickets)
	// fmt.Printf("Remaining tickets: %v\n", remainingTickets)
	// fmt.Printf("We will send you confirmation on your email id = %v\n", email)

}
