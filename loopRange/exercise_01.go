package main

import "fmt"

func main() {

	var input1 int
	fmt.Scan(&input1)

	var input2 int
	fmt.Scan(&input2)

	output := input1 % input2

	fmt.Println("The remainder of dividing ", input1, " by ", input2, " is ", output)
}
