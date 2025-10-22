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

func mergeTwoLists(list1 *Node, list2 *Node) *Node {

	dummy := &Node{}
	current := dummy

	for list1 != nil && list2 != nil {
		if list1.Value < list2.Value {
			current.Next = list1
			list1 = list1.Next
		} else {
			current.Next = list2
			list2 = list2.Next
		}
		current = current.Next
	}

	if list1 != nil {
		current.Next = list1
	} else if list2 != nil {
		current.Next = list2
	}

	return dummy.Next
}

func main() {
	listA := &Node{Value: 3, Next: &Node{Value: 6, Next: &Node{Value: 8}}}
	listB := &Node{Value: 4, Next: &Node{Value: 7, Next: &Node{Value: 9, Next: &Node{Value: 11}}}}

	mergedList := mergeTwoLists(listA, listB)
	fmt.Print("Список после слияния: ")
	printList(mergedList)
}
