package main

import "fmt"

func boolFunc() {
	bool1 := (true && false) || (false && true)
	bool2 := !(false && false)
	fmt.Println(bool1 || bool2)
	return
}

func twoParam(name string, age int)  {
	name = name
	age = age
	fmt.Println(name, "is", age, "years old")
}

func twoReturns(name string, age int) (int, bool)  {
	doggie := age * 7
	old := (age > 25)
	if old {
		fmt.Println(name, "is", doggie, "in dog years and is old")
		return age, (age > 25)
	} else {
		fmt.Println(name, "is", doggie, "in dog years and is not old", old)
		return age, (age > 25)
	}

}

func main()  {
	boolFunc()
	twoParam("John", 27)
	twoReturns("Jane", 30)
}
