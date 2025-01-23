package main

import (
	"fmt"
)

// Definition for singly-linked list.
type L1stNode struct {
	Val  int
	Next *L1stNode
}

// Function to create a new linked list
//func createLinkedList(values []int) *L1stNode {
//	if len(values) == 0 {
//		return nil
//	}
//
//	head := &L1stNode{Val: values[0]}
//	current := head
//	for _, value := range values[1:] {
//		current.Next = &L1stNode{Val: value}
//		current = current.Next
//	}
//	return head
//}

// Function to print the linked list
func printLinkedList(head *L1stNode) {
	current := head
	for current != nil {
		fmt.Print(current.Val, " ")
		current = current.Next
	}
	fmt.Println()
}

// Definition for reverseList
func reverseList(head *L1stNode) *L1stNode {
	var pre *L1stNode
	cur := head
	for cur != nil {
		next := cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}
	return pre
}

//func main() {
//	// Create a linked list 1 -> 2 -> 3 -> 4 -> 5
//	head := createLinkedList([]int{1, 2, 3, 4, 5})
//
//	// Print original linked list
//	fmt.Print("Original list: ")
//	printLinkedList(head)
//
//	// Reverse the linked list
//	reversedHead := reverseList(head)
//
//	// Print reversed linked list
//	fmt.Print("Reversed list: ")
//	printLinkedList(reversedHead)
//}
