package main

import "fmt"

func main() {

	type reminder struct {
		name string
		date string
		userID string
	}

	notes := make(map[reminder]string)

	rem1 := reminder{}
    rem1.name = "Arriana"
    notes[rem1] = "new note"

	var num2 *int = new(int)
	// new(int) IS a pointer in its declaration. Therefore it is not lawful to
	// use: var num2 int = new(int) BECAUSE the 'int' is a REGULAR int, not a pointer!

	fmt.Println(num2) // returns address
	fmt.Println(*num2) // returns 0
	fmt.Println(&num2) // returns address

    fmt.Println(notes)

}
