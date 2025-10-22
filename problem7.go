package main

import (
	"fmt"
)

type Element struct {
	elementValue rune
}

func String(element *Element) string {
	return string(element.elementValue)
}

type Stack struct {
	elements     []*Element
	elementCount int
}

func NewStack() *Stack {
	return &Stack{
		elements: make([]*Element, 0),
	}
}

func (stack *Stack) Push(element *Element) {
	stack.elements = append(stack.elements[:stack.elementCount], element)
	stack.elementCount = stack.elementCount + 1
}

func (stack *Stack) Pop() *Element {
	if stack.elementCount == 0 {
		return nil
	}

	var length int
	var element *Element

	length = len(stack.elements)
	element = stack.elements[length-1]

	if length > 1 {
		stack.elements = stack.elements[:length-1]
	} else {
		stack.elements = stack.elements[0:]
	}
	stack.elementCount = len(stack.elements)
	return element
}

func isAlphanumeric(r rune) bool {
	return (r >= 'a' && r <= 'z') || (r >= '0' && r <= '9') || (r >= 'а' && r <= 'я') // Include Cyrillic for 'казак'.
}

func isPalindromeStack(s string) bool {
	stack := NewStack()

	runes := []rune(s)

	for _, r := range runes {
		// Only push alphanumeric characters
		if isAlphanumeric(r) {
			stack.Push(&Element{elementValue: r})
		}
	}

	for _, r := range runes {
		if !isAlphanumeric(r) {
			continue
		}

		if len(stack.elements) == 0 {
			return false // Should not happen if logic is correct.
		}
		poppedElement := stack.Pop()
		if r != poppedElement.elementValue {
			return false // Mismatch found.
		}
	}

	return true
}

func main() {
	fmt.Printf("isPalindromeStack('racecar'): %v\n", isPalindromeStack("racecar"))
	fmt.Printf("isPalindromeStack('hello'): %v\n", isPalindromeStack("hello"))
	fmt.Printf("isPalindromeStack('казак'): %v\n", isPalindromeStack("казак"))
}
