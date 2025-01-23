package main

import "fmt"

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

	return head.Next // 返回头结点
}

// 打印链表
func PrintList(head *ListNode) {
	current := head
	for current != nil {
		fmt.Print(current.Val, " ")
		current = current.Next
	}
	fmt.Println()
}

//func main() {
//	l := &ListNode{}
//	fmt.Printf("%t", l)
//}
