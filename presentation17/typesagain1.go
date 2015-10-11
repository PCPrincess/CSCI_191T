package main

import (
	"fmt"
)

type beverage struct {
	coffee string
	sugar bool
	cream bool
	//More to learn about declarations (sigh)
}

func drink (b beverage) {

	fmt.Println("Today, I'm drinking", b.coffee)
}

func main()  {
	type myType []int
	type myType2 bool
	var trueOrfalse myType2 = 3 < 4
	fmt.Println(trueOrfalse)

	var myDrink beverage
	myDrink.coffee = "French Roast"
	myDrink.cream = true
	myDrink.sugar = true
	drink(myDrink)

}

