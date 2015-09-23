package main

import "fmt"

func main() {

	aMap := map[int]string{
		0: "Zero",
		1: "One",
		2: "Two",
		3: "Three", // Weird, MUST have comma after last entry too!
	}
	fmt.Println(aMap)

	if element, exists := aMap[3]; exists {
		delete(aMap, 3)
		element = "Six" //I can change the Key's VALUE (element)!
		fmt.Println(element)
		fmt.Println(aMap[3]) // This prints NOTHING! (empty string) It's DELETED
		fmt.Println(aMap[2])
	}

	for key, value := range aMap {
		fmt.Println(key, " goes with ", value)
	}
	fmt.Println("The length of aMap is: ", len(aMap))

	var usingNew *int = new(int)
	fmt.Println(usingNew)
	fmt.Println(*usingNew)
	fmt.Println("True, it returns a pointer!")

	var stringNew *string = new(string)
	fmt.Println(stringNew)
	fmt.Println(*stringNew)

	var boolNew *bool = new(bool)
	fmt.Println(boolNew)
	fmt.Println(*boolNew)

	var usingMake []int = make([]int, 10)
	fmt.Println(usingMake)

	var mapMake map[int]string = make(map[int]string)
	fmt.Println(mapMake)
}
