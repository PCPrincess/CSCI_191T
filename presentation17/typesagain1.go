package main

import (
	"fmt"
)

func main()  {
	type myType []int
	type myType2 bool
	var trueOrfalse myType2 = 3 < 4
	fmt.Println(trueOrfalse)

}

