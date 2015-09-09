package main

import "fmt"

const myAge int  = 48

const (
	Broke = iota
	Quarter = iota * 25
	Fifty = iota * 25
	SevFive = iota * 25
	Dollar = iota * 25
)

func timesTwo(x *int){
	*x *= 2
}

/*
func timesTwoTest(x *int){
	x *= 2//Does This Work The Same As func timesTwo()?
}
*///The Above func timesTwoTest Throws an Error!
func main() {

	fmt.Println("My age is ", myAge)//Using myAge

	bookMark := 48
	fmt.Println("I saved a bookmark at ", &bookMark)//Printing Address

	var input int
	fmt.Print("Enter A Number From 1 - 10")
	fmt.Scan(&input)//Using Scan

	fmt.Println("Entered Input Using Scan: ", input)

	x := 5
	//y := 5
	timesTwo(&x)
	//timesTwoTest(&y)
	fmt.Println("Results Of func timesTwo: ", x)
	//fmt.Println("Results Of func timesTwoTest: ", y)

}
