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
	age     string
	phone   string
}

/*metodo per inserire un nuovo utente
  con i dati dati in input
  quindi nome, cognome, email, età e numero telefono
*/
func main() {

	//connect to db
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017/"))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("SEI CONNESSO AL MONGODB")
	fmt.Println()

	ctx, _ := context.WithTimeout(context.Background(), 200*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	//ricerca database esistente
	usersDatabase := client.Database("user_db")
	//new collection
	usersCollection := usersDatabase.Collection("user_details")

	fmt.Println("PER INSERIRE UN NUOVO UTENTE, PREMI 1\n" +
		"PER MODIFICARE UTENTE, PREMI 2\n" +
		"PER ELIMINARE UN UTENTE, PREMI 3")

	var input int
	fmt.Scanln(&input)

	switch input {
	case 1:
		var user_result User = insertNewUser()
		var name string = user_result.name
		var surname string = user_result.surname
		var email string = user_result.email
		var age string = user_result.age
		var phone string = user_result.phone

		user_final, err := usersCollection.InsertOne(ctx, bson.D{
			{Key: "name", Value: &name},
			{Key: "surname", Value: &surname},
			{Key: "email", Value: &email},
			{Key: "age", Value: &age},
			{Key: "phone", Value: &phone},
		})

		if err != nil {
			log.Fatal(err)
		}

		//generazione automatica id univoco
		fmt.Println(user_final.InsertedID)
		fmt.Println("Complimenti! Hai inserito un nuovo utente!")

	case 2:
		var mail string = updateUser()
		var new_name string
		var new_surname string
		var new_email string
		var new_age string
		var new_phone string

		var update_user User

		if mail != "" && usersCollection.FindOne(ctx, bson.D{{Key: "email", Value: &mail}}) != nil {

			fmt.Println("INSERISCI NOME: ")
			fmt.Scanln(&new_name)
			update_user.name = new_name

			fmt.Println("INSERISCI COGNOME: ")
			fmt.Scanln(&new_surname)
			update_user.surname = new_surname

			fmt.Println("INSERISCI EMAIL: ")
			fmt.Scanln(&new_email)
			update_user.email = new_email

			fmt.Println("INSERISCI ETA': ")
			fmt.Scanln(&new_age)
			update_user.age = new_age

			fmt.Println("INSERISCI TELEFONO: ")
			fmt.Scanln(&new_phone)
			update_user.phone = new_phone
			/*
				usersCollection.ReplaceOne(ctx, bson.D{
					{Key: "name", Value: new_name},
					{Key: "surname", Value: new_surname},
					{Key: "email", Value: new_email},
					{Key: "age", Value: new_age},
					{Key: "phone", Value: new_phone},
				}, bson.D{
					{Key: "age", Value: &new_age},
				},
				)*/

		}

	case 3:
		deleteUser()

	}

	defer client.Disconnect(ctx)

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
	var age string
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
/*func updateByAge() string {
	fmt.Println("QUALE UTENTE INTENDI MODIFICARE? DIGITA LA SUA MAIL")
	var mail string
	fmt.Scanln(&mail)

	//if la mail è contenuta nella lista di mail presa tramite query dal db, allora procedi con la modifica di età dell'utente

	return mail
}*/
/*
//modifica numero cellulare utente
func updateByPhone() string {
	fmt.Println("QUALE UTENTE INTENDI MODIFICARE? DIGITA LA SUA MAIL")
	var phone string
	fmt.Scanln(&phone)

	//if la mail è contenuta nella lista di mail presa tramite query dal db, allora procedi con la modifica di numero di telefono dell'utente
	return phone
}*/

//metodo per modificare utente
func updateUser() string {
	fmt.Println("QUALE UTENTE INTENDI MODIFICARE? DIGITA LA SUA MAIL")
	var mail string
	fmt.Scanln(&mail)

	return mail
}

//elimina utente
func deleteUser() string {
	fmt.Println("QUALE UTENTE INTENDI ELIMINARE? DIGITA LA SUA MAIL")
	var mail string
	fmt.Scanln(&mail)

	//if la mail è contenuta nella lista di mail presa tramite query dal db, allora procedi con l'eliminazione dell'utente
	return mail
}
