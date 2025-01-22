package main

import "fmt"

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

// 创建链表
func CreateList(values []int) *ListNode {
	head := &ListNode{} // 创建头结点
	current := head
	for _, val := range values {
		current.Next = &ListNode{Val: val} // 创建新节点并连接
		current = current.Next             // 移动到新节点
	}

	return head // 返回头结点
}

func RemoveElements(head *ListNode, val int) *ListNode {
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
func PrintList(head *ListNode) {
	current := head.Next // 跳过头结点
	for current != nil {
		fmt.Print(current.Val, " ")
		current = current.Next
	}
	fmt.Println()
}

func main() {
	// 创建链表并打印
	values := []int{1, 2, 6, 3, 4, 5, 6}
	head := CreateList(values) // 使用输入数组创建链表
	RemoveElements(head, 6)
	PrintList(head) // 打印链表
}
