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
		"PER MODIFICARE NUMERO TELEFONO UTENTE, PREMI 2\n" +
		"PER MODIFICARE ETA' UTENTE, PREMI 3\n" +
		"PER ELIMINARE UN UTENTE, PREMI 4")

	var input int
	fmt.Scanln(&input)

	switch input {
	case 1:
		//caso di inserimento nuovo utente
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
		//caso in cui si modifica il numero di cellulare di un utente, data la sua email
		var e_mail = updateByPhone()

		var new_phone string
		fmt.Println("INSERISCI NUOVO NUMERO CELLULARE DELL'UTENTE CON LA SEGUENTE EMAIL: ", e_mail)
		fmt.Scanln(&new_phone)
		var nphone string = new_phone

		fmt.Println(nphone)

		if e_mail != "" {
			fmt.Println(usersCollection.FindOne(ctx, bson.D{
				{Key: "email", Value: e_mail},
			}).DecodeBytes())

			result := usersCollection.FindOneAndUpdate(ctx,
				bson.M{"email": e_mail},
				bson.D{
					{"$set", bson.D{{"phone", new_phone}}},
				},
			)

			fmt.Println("HAI MODIFICATO IL NUMERO DI CELLULARE COME RICHIESTO.")
			fmt.Println(result.DecodeBytes())
		}
		/*
			fmt.Println(usersCollection.FindOne(ctx, bson.D{
				{Key: "email", Value: e_mail},
			}).DecodeBytes())*/

	case 3:
		//caso in cui si modifica l'età di un utente, data la sua email
		var e_mail = updateByAge()

		var new_age string
		fmt.Println("INSERISCI NUOVA ETA' DELL'UTENTE CON LA SEGUENTE EMAIL: ", e_mail)
		fmt.Scanln(&new_age)
		var nage string = new_age

		fmt.Println(nage)

		if e_mail != "" {
			fmt.Println(usersCollection.FindOne(ctx, bson.D{
				{Key: "email", Value: e_mail},
			}).DecodeBytes())

			result := usersCollection.FindOneAndUpdate(ctx,
				bson.M{"email": e_mail},
				bson.D{
					{"$set", bson.D{{"age", new_age}}},
				},
			)

			fmt.Println("HAI MODIFICATO L'ETA' COME RICHIESTO.")
			fmt.Println(result.DecodeBytes())
		}

		/*fmt.Println(usersCollection.FindOne(ctx, bson.D{
			{Key: "email", Value: e_mail},
		}).DecodeBytes())*/

	case 4:
		//caso in cui si elimina un utente, data la sua email
		var e_mail = deleteUser()

		if e_mail != "" {
			fmt.Println(usersCollection.FindOne(ctx, bson.D{
				{Key: "email", Value: e_mail},
			}).DecodeBytes())
			fmt.Println("UTENTE TROVATO: " + e_mail)

			result := usersCollection.FindOneAndDelete(ctx, bson.M{"email": e_mail})

			fmt.Println("HAI ELIMINATO L'UTENTE: " + e_mail)
			fmt.Println(result.DecodeBytes())
		}

		/*	fmt.Println(usersCollection.FindOne(ctx, bson.D{
			{Key: "email", Value: e_mail},
		}).DecodeBytes())*/
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

//modifica numero cellulare utente
func updateByPhone() string {
	fmt.Println("QUALE UTENTE INTENDI MODIFICARE? DIGITA LA SUA MAIL")
	var mail string
	fmt.Scanln(&mail)

	return mail
}

//modifica età utente
func updateByAge() string {
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

	return mail
}
