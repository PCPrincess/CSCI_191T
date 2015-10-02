package main

import "fmt"

func namedReturn(age int) (dogYears int) {
	return age * 7
}

func variad(n ...int) int {
	temp := 0
	for _, value := range n {
		if value > temp {
			temp = value
		}
	}
	return temp
}

func main()  {
	//namedReturn(41)
	fmt.Println(namedReturn(41))

	fmt.Println(variad(1, 3, 5))
}
