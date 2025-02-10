package main

import (
	"fmt"
	"strings"
)

func main() {
	var conferenceName string = "Go Conference"
	// conferenceName := "Go Conference" -- doesn't work with constants or if you want to explicitly define a type of your variable

	var bookings []string
	//Arrays and Slices -- slices are dinamic arrays

	const conferenceTickets int = 50
	var remainingTickets uint = 50

	fmt.Printf("conferenceTickets is %T, remainingTickets is %T, conferenceName is %T\n", conferenceTickets, remainingTickets, conferenceName)

	fmt.Printf("Welcome to %v booking application.\n", conferenceName)
	fmt.Printf("We have total of %v tickets, and %v are still available.\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")

	for {
		var firstName string
		var lastName string
		var email string
		var userTickets uint

		// ask user for their name -- user input using pointer &
		fmt.Println("Enter your first name: ")
		fmt.Scan(&firstName)

		fmt.Println("Enter your last name: ")
		fmt.Scan(&lastName)

		fmt.Println("Enter your email: ")
		fmt.Scan(&email)

		fmt.Println("Enter number of tickets: ")
		fmt.Scan(&userTickets)


		// validate user input 
		isValidName := len(firstName) >=2 && len(lastName) >= 2
		isValidEmail := strings.Contains(email, "@")
		isValidTicketNumber := userTickets > 0 && userTickets <= remainingTickets

		// OR operator ||
		// isValidCity := city = "Singapore" || city == "London"
		// validate city which is either Singapore or London, same as -- !isValidCity


		if isValidName && isValidEmail && isValidTicketNumber { 
			userTickets <= remainingTickets { // check if we have enough tickets. If yes - 
			remainingTickets = remainingTickets - userTickets
			bookings = append(bookings, firstName+" "+lastName)
	
			fmt.Printf("The whole slice: %v\n", bookings)
			fmt.Printf("The first value: %v\n", bookings[0])
			fmt.Printf("Slice type: %T\n", bookings)
			fmt.Printf("Slice length: %v\n", len(bookings))
	
			fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v.\n", firstName, lastName, userTickets, email)
			fmt.Printf("%v tickets remaining for %v.\n", remainingTickets, conferenceName)
	
			// To show only the first names
			// range expression allows to itterate through elements of diff data types, providing index and value for each element
			// for loop inside another for loop
			firstNames := []string{}
			for _, booking := range bookings { // To ignore index variable we are using blank identifier _
				var names = strings.Fields(booking) // Splits all the users' name into 2 strings - firs name and last name
				// var firstName = names[0]
				firstNames = append(firstNames, names[0])
			}
	
			fmt.Printf("The first names of our bookings are: %v\n", firstNames)
	
			// noTicketsRemaining := remainingTickets == 0
			if remainingTickets == 0 {
				// end program
				fmt.Println("Our conference is booked up. Come back next year")
				break
	
			}
		} else { // If we don't have enough tickets -- the first user input is bigger then the anout of tickets -- we can break the loop and end the program using break expression 
			// or we can reset the condition using continue expression, so we go to the next iterration of the loop
			// in this case we don't need continue, it will automatically go to the next iterration 
			// fmt.Printf("We only have %v tickets remaining, you cannot book %v tickets,\n", remainingTickets, userTickets)

			if !isValidName {
				fmt.Println("First name or last name you entered is too short")
			}
			if !isValidEmail {
				fmt.Println("Email address you entered does not contain @ sign")
			}
			if !isValidTicketNumber {
				fmt.Println("Number of tickets you entered is invalid")
			}
			// fmt.Println("Your input data is invalid, try again")
		}

	}
}
