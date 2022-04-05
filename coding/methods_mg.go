package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type User struct {
	name    string
	surname string
	email   string
	age     int
	phone   string
}

/*metodo per inserire un nuovo utente
  con i dati dati in input
  quindi nome, cognome, email, età e numero telefono
*/
func main() {

	var collection mongo.Collection = connect_toMongoDB()
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	fmt.Println("PER INSERIRE UN NUOVO UTENTE, PREMI 1\n" +
		"PER MODIFICARE ETA' UTENTE, PREMI 2\n" +
		"PER MODIFICARE NUMERO TELEFONO UTENTE, PRIMI 3\n" +
		"PER ELIMINARE UN UTENTE, PREMI 4")

	var input int
	fmt.Scanln(&input)

	switch input {
	case 1:
		var user_result User = insertNewUser()
		fmt.Println("sono nel case 1")

		user_final, err := collection.InsertOne(ctx, bson.D{
			{"name", user_result.name},
			{"surname", user_result.surname},
			{"email", user_result.email},
			{"age", user_result.age},
			{"phone", user_result.phone},
		})

		if err != nil {
			log.Fatal(err)
		}
		//generazione automatica id univoco
		fmt.Println(user_final.InsertedID)
		fmt.Println("Complimenti! Hai inserito un nuovo utente!")

	case 2:
		updateByAge()

	case 3:
		updateByPhone()

	case 4:
		deleteUser()
	}

}

//metodo per inserire un nuovo utente
func insertNewUser() User {
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

	var newUser User
	newUser.name = name
	newUser.surname = surname
	newUser.email = email
	newUser.age = age
	newUser.phone = phone

	fmt.Println(newUser)

	return newUser
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

//connessione al mongodb
func connect_toMongoDB() mongo.Collection {
	/*
	   Connect to my cluster
	*/
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017/"))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("SEI CONNESSO AL MONGODB")
	fmt.Println()

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(ctx)

	//new database
	usersDatabase := client.Database("user_db")
	//new collection
	usersCollection := usersDatabase.Collection("user_details")

	return *usersCollection
}
