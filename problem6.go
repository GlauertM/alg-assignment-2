package main

import (
	"fmt"
)

func isSubsequenceTwoPointers(a, b string) bool {
	runeA := []rune(a)
	runeB := []rune(b)

	i := 0
	j := 0

	for i < len(runeA) && j < len(runeB) {
		if runeA[i] == runeB[j] {
			i++
		}
		j++
	}

	return i == len(runeA)
}

func main() {
	fmt.Printf("%v\n", isSubsequenceTwoPointers("abc", "abcrfftm"))
	fmt.Printf("%v\n", isSubsequenceTwoPointers("abc", "unttdt"))

}
