package main
import "strings"


func validateUserInput(firstName string, lastName string, email string, userTickets uint , remainingTickets uint) (bool, bool, bool) {
	isValidName := len(firstName) >= 2 && len(lastName) >= 2
	isValidEmail := strings.Contains(email, "@") // strings.Contains is an in-built method which checks the specific value is present in the passed argument or not
	isValidTicketNumber := userTickets > 0 && userTickets <= remainingTickets

	return isValidName, isValidEmail, isValidTicketNumber
}

