package main

import "fmt"

// Create Slice:  []Type   (Bracket on Left) -or-
//                make([]Type, length, capacity) -or-
//                new([capacity]Type)[0:length]
// Index Slice:   Slice[:]  (Bracket on Right)
// Indexing Works on Slice[:] and "Strings"[2]

func main() {
	aSlice := make([]int, 3, 20)
	bSlice := []int{1, 3, 5}
	for index, value := range aSlice  {
		fmt.Println(index, " - ", value)
	}

	aSlice[0] = 2
	aSlice[1] = 4
	aSlice[2] = 6

	fmt.Println(aSlice[0], aSlice[1], aSlice[2])
// To add items to a slice past the length, up to and beyond capacity: use APPEND
// Go automatically adds double the length of the underlying array when the length exceeds capacity
	aSlice = append(aSlice, 8)
	fmt.Println(aSlice[0], aSlice[1], aSlice[2], aSlice[3])
	//fmt.Println(aSlice[4]) <--  Index out of range!!
// Appending a Slice to a Slice! (...)
	aSlice = append(aSlice, bSlice...) //More Possible Args Dots
	fmt.Println(aSlice)

	cSlice := []string{"Have", "you", "tried", "a", "Tootsie", "Roll", "Lately?"}
	fmt.Println(cSlice)

// Deleting by 'Overwriting'
	dSlice := append(cSlice[:0], cSlice[4:6]...) //Don't Forget Dots!
	fmt.Println(dSlice)
}
