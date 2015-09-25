package main

import "fmt"

func main()  {
	type customer struct {
		name string
		ident int
	}

	cust1 := customer{"David", 001}
	cust2 := customer{"Sarah", 002}

	fmt.Println(cust1.name)
	fmt.Println(cust1.ident)

	cust2.name = "Sally"
	fmt.Println(cust2.name)


	cust3 := new(customer)
	cust3.name = "Billy"
	fmt.Println(cust3)

	//Can't use 'make' on structs!

}
