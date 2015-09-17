package main

// Note: move this file to a new repository for CSCI_156 - remove this package (myStuff)
import "fmt"

type Contact struct {
	name string
	address string
	phone string
}

func main() {
	var dataBase [5]Contact

	var a Contact
	//var b Contact
	//var c Contact

	var d = Contact{name: "Lori", address: "555 West Ave"}

	a.name = "Jessie"
	a.address = "1010 Park"
	a.phone = "555-5555"

	dataBase[0] = a
	dataBase[1] = d
	fmt.Println(dataBase[0].name)

}



