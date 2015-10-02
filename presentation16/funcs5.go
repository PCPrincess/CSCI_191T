package main

import "fmt"

func calBack(a string, callback func(string))  {
	callback(a)
}

func fibFunc(n int) int {
	if n == 0 {
		return 0
	} else if n == 1 {
		return 1
	} else {
		return fibFunc(n - 1) + fibFunc(n - 2)
	}
}

func endOfProg() {
	fmt.Println("That's All Folks!")
}

func main()  {
	defer endOfProg()
	calBack("Lori", func(name string){
		fmt.Println(name)
	} )

	fmt.Println(fibFunc(8))
}
