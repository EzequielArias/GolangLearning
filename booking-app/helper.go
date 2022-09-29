package main

import "strings"

func validateUserInput(firstName string, lastName string, email string, userTickets uint) (bool, bool, bool) {
	isValidName := len(firstName) >= 2 && len(lastName) >= 2
	isValidEmail := strings.Contains(email, "@")
	isValidTicketNumber := userTickets > 0 && userTickets <= remainingTickets
	return isValidName, isValidEmail, isValidTicketNumber
}

//Para exportar la funcion solo necesitamos hacer la primera letra mayusculad de la funcion

//Para ejecutar todos los archivos .go solo hay que ejecutar en la consola desde la carpeta correcta |"go run."|
