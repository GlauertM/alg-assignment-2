package main

import (
	"fmt"
)

type Node struct {
	Value int
	Next  *Node
}

func hasCycle(head *Node) bool {
	if head == nil || head.Next == nil {
		return false
	}
	var slow *Node
	var fast *Node
	slow = head
	fast = head

	for slow != fast {
		if fast == nil || fast.Next == nil {
			return false
		}
		slow = slow.Next
		fast = fast.Next.Next
	}

	return true
}

func main() {

	node1 := &Node{Value: 1}
	node2 := &Node{Value: 3}
	node3 := &Node{Value: 5}
	node4 := &Node{Value: 9}
	node5 := &Node{Value: 17}
	node6 := &Node{Value: 13}
	node7 := &Node{Value: 45}

	node1.Next = node2
	node2.Next = node3
	node3.Next = node4
	node4.Next = node5
	node5.Next = node6
	node6.Next = node7
	node7.Next = node3

	result := hasCycle(node1)
	fmt.Printf("%v\n", result)
}
