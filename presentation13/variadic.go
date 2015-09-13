package main

import (
	"fmt"
)

func main() {
	var name string
	fmt.Print("What is your name?")
	fmt.Scan(&name)
	tripleDot(name, "Middle ", "Last")

	intLine := []int{3, 6, 9, 12}
	sliceArg(intLine)
}

func tripleDot(s string, varies ...string )  {
	fmt.Println(s, varies)

}

func sliceArg(intSlice []int) {
	for i := 0; i <= len(intSlice) ; i++ {
		fmt.Println(i)
	}
}
