package main

import "fmt"

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

// 创建一个带头结点的链表
func createList() *ListNode {
	head := &ListNode{} // 创建头结点
	current := head
	values := []int{1, 2, 6, 3, 4, 5, 6}

	for _, val := range values {
		current.Next = &ListNode{Val: val} // 创建新节点并连接
		current = current.Next             // 移动到新节点
	}

	return head // 返回头结点
}

func removeElements(head *ListNode, val int) *ListNode {
	dummyHead := &ListNode{}
	dummyHead.Next = head
	cur := dummyHead
	for cur != nil && cur.Next != nil {
		if cur.Next.Val == val {
			cur.Next = cur.Next.Next
		} else {
			cur = cur.Next
		}
	}
	return dummyHead
}

// 打印链表
func printList(head *ListNode) {
	current := head.Next // 跳过头结点
	for current != nil {
		fmt.Print(current.Val, " ")
		current = current.Next
	}
	fmt.Println()
}

func main() {
	head := createList() // 创建链表
	removeElements(head, 6)
	printList(head) // 打印链表
}
