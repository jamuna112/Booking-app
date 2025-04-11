package main

import (
	"fmt"
	"strconv"
)

// Writing a program to accept first name, last name and number of bookings.
func main() {

	const TOTALNUMBEROFTICKETS int = 20

	var (
		firstName          string
		lastName           string
		email              string
		response           string
		remainingTickets   int
		bookedTickets      int
		totalbookedTickets int
		cont               bool
	)

	//slice of map to store multiple users
	users := []map[string]string{}

	fmt.Println("===============================================================")
	fmt.Printf("WELCOME TO BOOKING APP\n")
	fmt.Println("===============================================================")
	fmt.Printf("Total tickets available %v\n", TOTALNUMBEROFTICKETS)
	remainingTickets = TOTALNUMBEROFTICKETS
	for {

		fmt.Printf("Enter firstname: ")
		fmt.Scanf("%v", &firstName)

		fmt.Printf("Enter lastname: ")
		fmt.Scanf("%v\n", &lastName)

		fmt.Printf("Enter email: ")
		fmt.Scanf("%v\n", &email)

		fmt.Printf("Enter number of tickets: ")
		fmt.Scanf("%v\n", &bookedTickets)

		for bookedTickets > remainingTickets || bookedTickets <= 0 {
			fmt.Printf("Please re-enter the tickets within range 1 - %v\n", remainingTickets)
			fmt.Printf("Enter number of tickets: ")
			fmt.Scanf("%v\n", &bookedTickets)
		}
		//map to add key value pair for each user
		userMap := map[string]string{
			"firstname":     firstName,
			"lastname":      lastName,
			"email":         email,
			"bookedTickets": strconv.Itoa(bookedTickets),
		}
		//add users to the slice of map
		users = append(users, userMap)
		remainingTickets = remainingTickets - bookedTickets
		totalbookedTickets += bookedTickets

		fmt.Printf("Thank you %s, for booking %d number of tickets. Email confirmation will be send to you email at %s\n", firstName, bookedTickets, email)
		fmt.Printf("Remaining tickets are %d\n\n", remainingTickets)

		if remainingTickets == 0 {
			fmt.Println("No tickets available at this time.")
			break
		}

		cont = true
		fmt.Printf("Do you want to continue?")
		fmt.Scanf("%s", &response)

		switch response {
		case "yes":
			cont = true
		case "no":
			cont = false
		default:
			for response != "yes" || response != "no" {
				fmt.Printf("Please enter yes or no\n")
				fmt.Printf("Do you want to continue?")
				fmt.Scanf("%s", &response)
				if response == "no" || response == "yes" {
					if response == "no" {
						cont = false

					}
					break
				}
			}
		}
		if !cont {
			break
		}

	}

	fmt.Println("================================")
	fmt.Printf("ALL BOOKINGS \n")
	fmt.Println("================================")
	fmt.Printf("|S.no| Name   | Bookings|\n")
	array := []int{}
	j := 0
	for i, val := range users {
		j, _ = strconv.Atoi(val["bookedTickets"])
		array = append(array, j)
		fmt.Printf("|%2v. | %6v | %6v| %6v\n", (i + 1), val["firstname"], val["bookedTickets"], j)

	}

	n := len(array)
	for i := 0; i < n-1; i++ {

		for j := 0; j < n-i-1; j++ {

			if array[j] < array[j+1] {

				array[j], array[j+1] = array[j+1], array[j]
			}
		}
	}
	fmt.Printf("Max booked ticket is %v\n", array[0])
	fmt.Println("================================")
	fmt.Printf("Total booked tickets: %v\n", totalbookedTickets)
	fmt.Println("================================")
	fmt.Printf("Thanks for using booking app!\n")

}
