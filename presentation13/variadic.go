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

	str := fmt.Sprint("Hey ", name, ",", " your Sprint test works!" )
	fmt.Println(str)
	//Note: Was unable to use a 'verb' in Sprint (above) - why?
	//Note: Fprintf for 'formatted' printing, including verbs (?)

	var textTest string
	fmt.Print("In order to test Sscan, I'll need you to enter a line of text and hit the enter key.")
	fmt.Sscan("%s", &textTest)
	fmt.Sprint(textTest)
	// This is going to take a bit more research *Scratches Head*
	// Note: Try using fmt.Printf (try soon)


}

func tripleDot(s string, varies ...string )  {
	fmt.Println(s, varies)

}

func sliceArg(intSlice []int) {
	for i := 0; i <= len(intSlice) ; i++ {
		fmt.Println(i)
	}
}
