package main

import "fmt"

func main() {
	myMap1 := map[int]string {
		1:	"One",
		2:	"Two",
		3:	"Three",
	}

	fmt.Println(myMap1[1])
	myMap1[4] = "Four"
	fmt.Println(myMap1[4])
	delete(myMap1, 2)
	fmt.Println(myMap1)
	myMap1[2] = "Two"

	if val, exists := myMap1[2]; exists {
		fmt.Println("The value", val, "exists. Delete?", exists)
		delete(myMap1, 2)
	}
	fmt.Println(myMap1)

	for key, val := range myMap1 {
		fmt.Println(key, "=", val)
	}

}
