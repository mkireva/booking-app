package helper

import "strings"

func ValidateUserInput(firstName string, lastName string, email string, userTickets uint, remainignTickets uint) (bool, bool, bool) {
	isValidName := len(firstName) >= 2 && len(lastName) >= 2
	isValidEmail := strings.Contains(email, "@")
	isValidTicketNumber := userTickets > 0 && userTickets <= remainignTickets
	return isValidName, isValidEmail, isValidTicketNumber
}