package main

import (
	"fmt"
	"github.com/PCPrincess/CSCI_191T/names"
)

func main() {
	var name string
	name = "Lori"
	message := name + " is one cool chick"

	fmt.Println(names.MyName)
	fmt.Println(names.yourName)//Can't Use!
	fmt.Println(names.secondTestFunc)//Can't Use!


//Below is Personal Use [Not Part Of Assignment]
	var message2 = "Go Can Infer This Is A String"
	var d, e, f = 2, false, 3
	message3 := "var Not Necessary Inside Functions"

	sum := 0
	for _, value := range array {
		sum += value
	}

	fmt.Println(message)
	fmt.Println(message2, d)
	fmt.Println(message3, e, f)

}

