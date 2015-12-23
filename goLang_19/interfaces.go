package main

import "fmt"

type Component struct {
	purpose string
}

type Monitor struct {
	Component
	wideScreen bool
}

type Memory struct {
	Component
	size int
}

// Everything implements an empty interface!
func specs(i interface{}) {
	fmt.Println(i)
}

func main() {

	dell := Monitor{Component{"viewing"}, true}
	ram := Memory{Component{"stack"}, 8}
	myPC := []interface{}{dell, ram}
	fmt.Println(myPC)
	specs(myPC)
}
