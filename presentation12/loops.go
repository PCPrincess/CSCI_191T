package main

import "fmt"

func main(){

	for i := 0; i <= 1000 ; i += 2 {
		fmt.Println(i)
	}

	for i := 0; i <= 100 ; i++ {
	//First Version
		var iCheck3 int
		iCheck3 = i % 3
		var iCheck5 int
		iCheck5 = i % 5

		switch true {
		case iCheck3 == 0:
			fmt.Println("Fizz")
		    fallthrough
		case iCheck5 == 0:
			fmt.Println("Buzz")
			fallthrough
		case iCheck3 == 0 && iCheck5 ==0:
			fmt.Println("FizzBuzz")

		}

	}
	var sumOfIChecks int
	for i := 0; i <= 1000 ; i++ {
	//Second Version with Sum of Multiples
	//May Still Have Errors (Needs More Debugging)
		var iCheck3 int
		iCheck3 = i % 3
		var iCheck5 int
		iCheck5 = i % 5

		switch true {
		case iCheck3 == 0:
			fmt.Println("Fizz")
			sumOfIChecks += i
			fallthrough
		case iCheck5 == 0:
			fmt.Println("FizzBuzz")

		case iCheck5 == 0:
			fmt.Println("Buzz")
			sumOfIChecks += i
		}

	}
	fmt.Println(sumOfIChecks)

	var wordUp string
	wordUp = "Word Up!"
	var count int
	for len(wordUp) > count {
		count++
		fmt.Println(wordUp)
	}

	for i := 0; i < 5; i++ {
		fmt.Println("i is ", i)
	}

	for i, j := 1, 2; i <= 4; i, j = i + 1 , j + 2{
		var sum int = i + j
		fmt.Println(sum)
	}
}
