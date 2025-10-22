package main

import (
	"fmt"
)

type Queue struct {
	Value int
	Next  *Node
}

func isSubsequenceQueue(a, b string) bool {
	queue := []rune(a)

	for _, charB := range b {
		// q.peek() == el
		if len(queue) > 0 && queue[0] == charB {
			// q.pop()
			queue = queue[1:]
		}
	}

	// return q.getSize() == 0
	return len(queue) == 0
}

func main() {
	fmt.Printf("%v\n", isSubsequenceQueue("abc", "abcrfftm"))
	fmt.Printf("%v\n", isSubsequenceQueue("abc", "unttdt"))
}
