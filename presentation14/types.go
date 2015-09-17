package main

import "fmt"

func main()  {
	output1 := "Schedule of Classes:\t Monday\t\t Tuesday\t Wednesday\n"
	output2 := "Class:\t\t\t CSCI 156\t CSCI 191T"
	fmt.Println(output1)
	fmt.Println(output2)

	miniAlpha := "ABCDEFGHIJKLM"
	miniMini := len(miniAlpha)
	fmt.Println(miniMini)
	miniMini += 3
	fmt.Println(miniMini)
}
