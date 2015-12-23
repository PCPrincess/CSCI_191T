package main

import (
	"fmt"
	"encoding/json"
)

// ***********************************************
//   CAPITALIZATION MATTERS
// -- capitalized: available (public)
// -- lower case: not available (private)
// ***********************************************

type reminder struct {
	name string
	date string
	userID string
}

type Data struct {
	Email string
	LoggedIn string
	Note []reminder

}

func main() {

	rem1 := reminder{"Jane", "10202015", "jjooze"}
	fmt.Println(rem1.userID)
	rem2 := reminder{"John", "10212105", "patsfan"}
	fmt.Println(rem2.userID)

	var n Data

	str :=
	`{"Email":"jj@hotmail.com", "LoggedIn":"true", "Note":["rem1.userID", "rem2"]} `
	bytStr := []byte(str) // won't work for the structs, only bytes!
	json.Unmarshal(bytStr, &n)

	fmt.Println(n)
	fmt.Println(n.Email)
	fmt.Println(n.LoggedIn)
	fmt.Println(n.Note) // Remember, the structs won't print!
}

