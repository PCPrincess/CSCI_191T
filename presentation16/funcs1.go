package main

import "fmt"

func half(x int) (y int) {
	remain := x % 2
	y = x / 2
	if remain == 0 {
		fmt.Println(y, "True")
	} else  {
		fmt.Println(y, "False")
	}

	return
}

func toddsHalf(n int) (int, bool)  {
	return n / 2, n % 2 == 0
	// Bool used as return should be an EXPRESSION
	// This was the cause of my hair-pull
}

// First attempt @ exercise without looking at solution
// Attempting to create a slice to hold inputs
// Can fine-tune this to work.....  (then I cheated and peeked)
// Todd's solution below mine
/* func greatNum(a ...int) int  {
	numList := make([]int)
	temp := 0
	for len(numList) <= len(a) {
		append(numList, a)
	}

	for value := range numList {
		if value > temp {
			temp = value
		}
	}

	return fmt.Println(temp)
} */

func maxNum(numbers ...int) int {
	var largest int
	for _, v := range numbers {
		if v > largest {
			largest = v
		}
	}
	return  largest
}

func main()  {
	half(10)
	//toddsHalf(10) can't be called like this

	halfCheck, even := toddsHalf(10)
	fmt.Println(halfCheck, even)
	// using toddsHalf must be done like this (above)

	greatest := maxNum(2, 4, 6, 8)
	fmt.Println(greatest)
}
