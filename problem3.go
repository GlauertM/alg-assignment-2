package main

import (
	"fmt"
)

type Node struct {
	Value int
	Next  *Node
}

func middleNode(head *Node) *Node {
	var slow *Node
	var fast *Node

	slow = head
	fast = head

	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	return slow
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

	result := middleNode(node1)
	fmt.Printf("%v\n", result)
	//Проверка пустого списка:
	nodeNil := (*Node)(nil)
	resultNil := middleNode(nodeNil)
	fmt.Printf("%v\n", resultNil)
	//Проверка, если эллемент один:
	nodeSingle := &Node{Value: 100}
	resultSingle := middleNode(nodeSingle)
	fmt.Printf("%v\n", resultSingle)

}
