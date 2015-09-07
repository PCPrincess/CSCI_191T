package myStuff

//Reversing a String

func reverseString(s string) {
	runeSlice := []rune(s)
	for i, j := 0, len(runeSlice) - 1; i < len(runeSlice) / 2; i, j = i + 1, j - 1 {
		runeSlice[i], runeSlice[j] = runeSlice[j], runeSlice[i]
	}
}
