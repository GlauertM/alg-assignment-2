package main

import (
	"fmt"
)

type Node struct {
	Value int
	Next  *Node
}

func printList(head *Node) {
	for head != nil {
		fmt.Printf("%d -> ", head.Value)
		head = head.Next
	}
	fmt.Println("nil")
}

func removeElements(head *Node, val int) *Node {
	var dummy *Node
	dummy = &Node{Value: 0}
	dummy.Next = head

	var prev *Node
	var current *Node
	prev = dummy
	current = head

	for current != nil {
		if current.Value == val {
			prev.Next = current.Next
		} else {
			prev = current
		}
		current = current.Next
	}

	return dummy.Next
}

func main() {
	//list
	node1 := &Node{Value: 1}
	node2 := &Node{Value: 2}
	node3 := &Node{Value: 3}
	node4 := &Node{Value: 8}
	node5 := &Node{Value: 5}

	node1.Next = node2
	node2.Next = node3
	node3.Next = node4
	node4.Next = node5
	node5.Next = nil
	// Подготовим для случая, где соостветсвия нет
	nodeFNone := node1

	// Проверка удаления первого элемента списка [head]
	resultIfHead := removeElements(node1, 1)
	printList(resultIfHead)
	// Если соостветсвия valREf=node[val] нет:
	resultNone := removeElements(nodeFNone, 1000)
	printList(resultNone)
	// На другом элементе, не первом
	result := removeElements(node1, 8)
	printList(result)

}
