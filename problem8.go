package main

import (
	"fmt"
)

func isPalindromeTwoPointers(s string) bool {
	runes := []rune(s) // Используем руны для кириллицы
	left := 0
	right := len(runes) - 1

	for left < right {
		if runes[left] != runes[right] {
			return false
		}
		left++
		right--
	}

	return true
}

func main() {
	fmt.Printf("isPalindromeTwoPointers('racecar'): %v\n", isPalindromeTwoPointers("racecar"))
	fmt.Printf("isPalindromeTwoPointers('hello'): %v\n", isPalindromeTwoPointers("hello"))
	fmt.Printf("isPalindromeTwoPointers('шалаш'): %v\n", isPalindromeTwoPointers("шалаш"))
}
