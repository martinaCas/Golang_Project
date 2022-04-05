package main

import (
	"fmt"
)

/*metodo per inserire un nuovo utente
  con i dati dati in input
  quindi nome, cognome, email, età e numero telefono
*/
func main() {
	fmt.Println("PER INSERIRE UN NUOVO UTENTE, PREMI 1\n" +
		"PER MODIFICARE ETA' UTENTE, PREMI 2\n" +
		"PER MODIFICARE NUMERO TELEFONO UTENTE, PRIMI 3\n" +
		"PER ELIMINARE UN UTENTE, PREMI 4")

	var input int
	fmt.Scanln(&input)

	switch input {
	case 1:
		insertNewUser()

	case 2:
		updateByAge()

	case 3:
		updateByPhone()

	case 4:
		deleteUser()
	}

}

//metodo per inserire un nuovo utente
func insertNewUser() {
	fmt.Println("INSERIMENTO NUOVO UTENTE, INSERISCI I DATI RICHIESTI.")
	fmt.Println("NOME: ")
	var name string
	fmt.Scanln(&name)

	fmt.Println("COGNOME: ")
	var surname string
	fmt.Scanln(&surname)

	fmt.Println("EMAIL: ")
	var email string
	fmt.Scanln(&email)

	fmt.Println("ETA': ")
	var age int
	fmt.Scanln(&age)

	fmt.Println("NUMERO DI TELEFONO: ")
	var phone string
	fmt.Scanln(&phone)

	fmt.Println("Complimenti! Hai inserito un nuovo utente!")

}

//modifica età utente
func updateByAge() {
	fmt.Println("QUALE UTENTE INTENDI MODIFICARE? DIGITA LA SUA MAIL")
	var mail string
	fmt.Scanln(&mail)

	//if la mail è contenuta nella lista di mail presa tramite query dal db, allora procedi con la modifica di età dell'utente

}

//modifica numero cellulare utente
func updateByPhone() {
	fmt.Println("QUALE UTENTE INTENDI MODIFICARE? DIGITA LA SUA MAIL")
	var mail string
	fmt.Scanln(&mail)

	//if la mail è contenuta nella lista di mail presa tramite query dal db, allora procedi con la modifica di numero di telefono dell'utente
}

//elimina utente
func deleteUser() {
	fmt.Println("QUALE UTENTE INTENDI ELIMINARE? DIGITA LA SUA MAIL")
	var mail string
	fmt.Scanln(&mail)

	//if la mail è contenuta nella lista di mail presa tramite query dal db, allora procedi con l'eliminazione dell'utente
}
