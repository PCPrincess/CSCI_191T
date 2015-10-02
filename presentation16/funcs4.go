package main

import "fmt"

func varTest() {
	a := 5
	fmt.Println(a)
}

func returnFunc(x int) func() int {
	tempVar := x
	return func() int {
		tempVar = tempVar * 2
		return  tempVar
	}
}

// CLOSURE
func main()  {

	varTest()
	a := 10
	fmt.Println(a)

	sumOfA := func() int {
		var sum = a + a
		return sum
	}
	timesTwo := returnFunc(8)
	fmt.Println(timesTwo())
	fmt.Println(sumOfA())
}

