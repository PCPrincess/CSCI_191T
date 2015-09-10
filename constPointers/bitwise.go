package main

import "fmt"

func main()  {
	var ten int = 10
	var shiftTen int
	shiftTen = ten << 2

	var three int = 3
	var shiftThree int
	shiftThree = three << 2

	fmt.Println("The binary representation of ten is: ")
	fmt.Printf("%b", ten)
	fmt.Println("\nThe binary representation of shiftTen is: ")
	fmt.Printf("%b", shiftTen)
	fmt.Println("\nThe decimal representation of shiftTen is: ")
	fmt.Printf("%d", shiftTen)

	fmt.Println("\nThe binary representation of three is: ")
	fmt.Printf("%b\t", three)
	fmt.Println("\nThe binary representation of shiftThree is: ")
	fmt.Printf("%b\t", shiftThree)
	fmt.Println("\nThe decimal representation of shiftThree is: ")
	fmt.Printf("%d", shiftThree)


}
