package main

import (
	"fmt" // package for IO and other functions like print , etc.
	"time" // provides delay in the thread
	"sync" // provides basic synchronization
)
 


const conferenceTicket uint = 50

var conferenceName = "Go Conference" // syntatic sugar declaration := which means it dynamically declares and assigns the datatype and the value respectively
var remainingTickets uint = 50
var bookings = make([]userData,0) // 0 -> it's the size and as the map grows dynamically , we dont need to worry about initializing the size . Can set to 0 or any number

type userData struct {
	firstName string
	lastName string
	email string
	numberOfTickets uint
}

var wg= sync.WaitGroup{} // syntax for declaring WaitGroup
func main() {

	userGreet()

	
		firstName, lastName, email, userTickets := getUserInput()

		isValidName, isValidEmail, isValidTicketNumber := validateUserInput(firstName, lastName, email, userTickets, remainingTickets)

		if isValidName && isValidEmail && isValidTicketNumber {
			bookTicket(userTickets, firstName, lastName, email)

			wg.Add(1) // Sets the no of goroutines to wait for (increases the counter by the provided number)

			go sendTickets(userTickets, firstName, lastName, email)

			// call func print first name
			firstNames := getFirstNames()
			fmt.Printf("The first names of bookings are: %v\n", firstNames)

			if remainingTickets == 0 {
				//end the program
				fmt.Println("Our conference is booked up. Come back next year")
				
			}

		} else {
			if !isValidName {
				fmt.Println("First name or Last name you've entered is too short")
			}
			if !isValidEmail {
				fmt.Println("Email Address not contains @ sign")
			}
			if !isValidTicketNumber {
				fmt.Println("You have entered wrong ticket")
			}
		}
	wg.Wait() // blocks untill the WaitGroup counter is 0
}


func userGreet() {
	fmt.Printf("Welcome to our %v booking application\n", conferenceName)
	fmt.Printf("We have total of %v conferenceTicket tickets and %v remainingTickets are available\n", conferenceTicket, remainingTickets)
	fmt.Println("Get your tickets here to attend the conference")

}

func getUserInput() (string, string, string, uint) {
	var firstName string
	var lastName string
	var email string
	var userTickets uint

	fmt.Print("Enter your First name: ")
	fmt.Scan(&firstName)

	fmt.Print("Enter your Last name: ")
	fmt.Scan(&lastName)

	fmt.Print("Enter your Email address: ")
	fmt.Scan(&email)

	fmt.Print("How many Tickets do you want to book: ")
	fmt.Scan(&userTickets)

	return firstName, lastName, email, userTickets
}

func bookTicket(userTickets uint, firstName string, lastName string, email string) {
	remainingTickets -= userTickets

	// create a map for a user
	var userData = userData{
		firstName: firstName,
		lastName: lastName,
		email: email,
		numberOfTickets: userTickets,
	}
	
	// make(map[key -> string] value ->string) make is the in-built method to create a empty map and map is the in-built keyword to declare a map
	// userData["firstName"] = firstName      // we cannot mix data types as values in map 
	// userData["lastName"] = lastName
	// userData["email"] = email
	// userData["numberOfTickets"] = strconv.FormatUint(uint64(userTickets), 10)

	bookings = append(bookings, userData) // append() is an in-built which is used to add the elements to the slices

	fmt.Printf("List of bookings is %v\n\n", bookings)
	
	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets remaining for the %v\n", remainingTickets, conferenceName)
}

func getFirstNames() []string { //(bookings []string -> input parameters) []string -> output parameters
	firstNames := []string{}
	for _, booking := range bookings {
		
		//var names = strings.Fields(booking) // strings.Fields() is used to split a string into substrings based on whitespace (spaces, tabs, or newlines). It returns a slice of the substrings.
		firstNames = append(firstNames, booking.firstName)
	}
	return firstNames
}

func sendTickets(userTickets uint, firstName string , lastName string, email string){
	time.Sleep(5 * time.Second)
	ticket := fmt.Sprintf("%v tickets for %v %v\n" ,userTickets, firstName, lastName)
	fmt.Println("#########################")
	fmt.Printf("Sending ticket:\n  %v \nto email address%v\n", ticket , email  )
	fmt.Println("#########################")
	wg.Done()  // Decrements the WaitGroup counter by 1 .. So this is called by the goroutine to indicate that it's finished
}