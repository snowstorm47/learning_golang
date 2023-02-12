// since this is the main folder we use the main name for our pacakge name
package main

import (
	"fmt"
	"strings"
)

func main() {
	//var conferenceName = "Go Conference" this is the longer way of decraring a variable we can also do the below
	conferenceName := "Go Conference" // this can not be done on constant variables and it doesnot work when u want to explicitly define a type for the variable
	const confrerenceTickets = 50     //total number of tickets
	var remainingTickets uint = 50

	// var bookings = [50]string{} for a shorter declaration we can use
	//var bookings [50]string to change this array in to a slice we just leave out the size of the array its besically the same for slices too
	// for an alternaitve for slice we can also use bookings := []string{}
	var bookings []string //this is a slice

	for {

		var firstName string
		var lastName string
		var email string
		var userTickets uint

		fmt.Printf("Welcome to our %v Booking Application\n", conferenceName)
		fmt.Printf("we have a total of %v tickets and %v are still remaining.\n", confrerenceTickets, remainingTickets) //always do it in the right order when referencing the variables
		fmt.Println("Book your ticket here")

		fmt.Println("Enter your first name: ")
		fmt.Scan(&firstName) // the pointer to the memeory location and the variable to store the data

		fmt.Println("Enter your last name: ")
		fmt.Scan(&lastName)

		fmt.Println("Enter your email: ")
		fmt.Scan(&email)

		fmt.Println("Enter number of tickets: ")
		fmt.Scan(&userTickets)
		//before we set the remaining tickets we need to check if the usertickets are more then the remaining tickets to stop form errors happening

		if userTickets > remainingTickets {

			fmt.Printf("We only have %v Tickets remaining so you can not book %v tickets\n", remainingTickets, userTickets)
			continue
		}

		//adding logic to reduce the number of tickets in remaining tickets so as to show the tickets are being booked
		remainingTickets = remainingTickets - userTickets

		//here is the arrya assignment
		//bookings[0] = firstName + " " + lastName// to change this array assignmet to slice we use the append() function
		bookings = append(bookings, firstName+" "+lastName)

		//show the details of the slice
		// fmt.Printf("the whole slice displayed: %v\n", bookings)
		// fmt.Printf("the first instance of the slice: %v\n", bookings[0])
		// fmt.Printf("the length of the slice: %v\n", len(bookings))

		fmt.Printf("Thank your %v %v for booking %v tickets. you will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)
		fmt.Printf("%v tickets are remaining for the %v\n", remainingTickets, conferenceName)
		//to display the bookings with only first names
		firstNames := []string{}
		for _, booking := range bookings {
			var names = strings.Fields(booking)
			firstNames = append(firstNames, names[0])
		}
		fmt.Printf("All the bookings for the %v are: %v\n", conferenceName, firstNames)
		// with in our for loop we have to stop the program when all the tickets have been booked to do that we use if else statements

		if remainingTickets == 0 {
			// here end the program
			fmt.Println("The conference is full thank you so much hope to see you in the next one ")
			break
		}
	}

}
