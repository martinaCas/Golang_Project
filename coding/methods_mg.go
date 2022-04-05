package methods

import (
	"fmt"
)

/*metodo per inserire un nuovo utente
  con i dati dati in input
  quindi nome, cognome, email, età e numero telefono
*/
func main() {
	/*
	   Connect to my cluster
	*/
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017/"))
	if err != nil {
		log.Fatal(err)
	}

func updateByAge(new_age int) {

}

	fmt.Println("ciao")

}

func insertNewUser(name, surname, email string, age int, phone string) {

}
