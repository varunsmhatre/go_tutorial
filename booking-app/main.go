package main

import (
	"booking-app/helper"
	"fmt"
	"sync"
	"time"
)

const conferenceTickets int = 50

var conferenceName string = "Go Conference"

var remainingTickets uint = 50

type UserData struct {
	firstName       string
	lastName        string
	email           string
	numberOfTickets uint
}

// var bookings []string
var bookings = make([]UserData, 0)

// bookings := []string{}

var wg = sync.WaitGroup{}

func main() {

	greetUsers()

	for {

		firstName, lastName, email, userTickets := getUserInput()

		isValidName, isValidEmail, isValidTicketNumber := helper.ValidateUserInput(firstName, lastName, email, userTickets, remainingTickets)

		if isValidName && isValidEmail && isValidTicketNumber {
			bookTickets(userTickets, firstName, lastName, email)
			wg.Add(1)
			go sendTicket(userTickets, firstName, lastName, email)

			fmt.Printf("Booking Names: %v\n", getFirstNames())

			fmt.Println("------------------------------------------------")

			if remainingTickets <= 0 {
				fmt.Printf("All Tickets sold for %v\n", conferenceName)
				wg.Wait()
				break
			}

		} else {
			if !isValidName {
				fmt.Println("First name or last name is too short")

			}
			if !isValidEmail {
				fmt.Println("Email is invalid")
			}
			if !isValidTicketNumber {
				fmt.Println("Number of tickets is invalid")
			}
			fmt.Println("------------------------------------------------")
		}
	}

}

func greetUsers() {
	fmt.Printf("Welcome to %v Booking Application\n", conferenceName)
	fmt.Printf("Total tickets are %v while remaining tickets are %v\n", conferenceTickets, remainingTickets)
	fmt.Println("Get you Tickets here")
}

func getFirstNames() []string {
	firstNames := []string{}
	for _, booking := range bookings {
		firstNames = append(firstNames, booking.firstName)
	}
	return firstNames
}

func getUserInput() (string, string, string, uint) {
	var firstName, lastName, email string
	var userTickets uint

	// Get User Details
	fmt.Print("Enter Your First Name: ")
	fmt.Scan(&firstName)

	fmt.Print("Enter Your Last Name: ")
	fmt.Scan(&lastName)

	fmt.Print("Enter Your Email: ")
	fmt.Scan(&email)

	fmt.Print("Enter the number of tickets you want: ")
	fmt.Scan(&userTickets)
	return firstName, lastName, email, userTickets

}

func bookTickets(userTickets uint, firstName string, lastName string, email string) {
	remainingTickets = remainingTickets - userTickets

	// create a struct
	var userData = UserData{
		firstName:       firstName,
		lastName:        lastName,
		email:           email,
		numberOfTickets: userTickets,
	}
	bookings = append(bookings, userData)

	fmt.Printf("Thank you %v %v for booking %v tickets, you will receive confimation email at %v\n", firstName, lastName, userTickets, email)
	fmt.Printf("%v Tickets are remaining for %v\n", remainingTickets, conferenceName)

}

func sendTicket(userTickets uint, firstName string, lastName string, email string) {
	time.Sleep(10 * time.Second)
	var ticket = fmt.Sprintf("%v tickets for %v %v\n", userTickets, firstName, lastName)
	fmt.Println("##################")
	fmt.Printf("Sending ticket:\n %v to email address %v\n", ticket, email)
	fmt.Println("##################")
	wg.Done()
}
