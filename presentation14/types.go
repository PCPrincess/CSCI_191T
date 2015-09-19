package main

import (
	"fmt"
	"strconv"
	"math"
)

func main()  {
	output1 := "Schedule of Classes:\t Monday\t\t Tuesday\t Wednesday\n"
	output2 := "Class:\t\t\t CSCI 156\t CSCI 191T"
	fmt.Print(output1)
	fmt.Println(output2)

	miniAlpha := "ABCDEFGHIJKLM"
	miniMini := len(miniAlpha)
	fmt.Println(miniMini)
	miniMini += 3
	fmt.Println(miniMini)
	fmt.Println(miniAlpha[0:6])
	oneChar := miniAlpha[0:1]
	twoChar := miniAlpha[1:2]
	fmt.Println(oneChar)
	//fmt.Println(rune(oneChar)) //ERROR, can't convert string to rune
	fmt.Println(oneChar + twoChar)
	//Convert int to string w/ strconv.Itoa(int) or fmt.Sprint(int)
	var m int = 10;
	n, _ := strconv.Atoi("100") //MUST USE 2nd variable!
	fmt.Println(strconv.Itoa(m))
	fmt.Println(n)

	intage := 5
	floatage := 5.5
	fmt.Println(float64(intage) + floatage)
	fmt.Println(intage + int(floatage))

	byteToStr := []byte{'W', 'h', 'a', 't', 's'}
	newString := string(byteToStr) + "up?"
	fmt.Println(newString)

	strToByte := "Man, this presentation is long!"
	utfCheck := []byte(strToByte)
	//fmt.Println("Is my string: " + utfCheck + " in UTF format?")
	//Error! Can't mix string and byte-slices in an output
	fmt.Println(utfCheck)

	var x float64
	fmt.Print("Enter a decimal number no higher than 100: ")
	fmt.Scan(&x)
	fmt.Println(math.Ceil(x))


}
