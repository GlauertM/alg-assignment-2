package main

import (
	"fmt"
)

type Node struct {
	Value int
	Next  *Node
}

func reverseLinkedList(head *Node) *Node {
	var prev *Node
	var current *Node

	prev = nil
	current = head

	for current != nil {
		next := current.Next
		current.Next = prev
		prev = current
		current = next
	}

	head = prev

	return head
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
	node7.Next = nil

	nodeNil := (*Node)(nil)
	resultNil := reverseLinkedList(nodeNil)
	fmt.Printf("%v\n", resultNil)

	result := reverseLinkedList(node1)
	for testN := result; testN != nil; testN = testN.Next {
		fmt.Printf("%v\n", testN)
	}
}
