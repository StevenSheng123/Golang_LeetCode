package main

import "fmt"

// Definition for singly-linked list.
type L2stNode struct {
	Val  int
	Next *L2stNode
}

func Cr1ateList(values []int) *L2stNode {
	head := &L2stNode{} // 创建头结点
	current := head
	for _, val := range values {
		current.Next = &L2stNode{Val: val} // 创建新节点并连接
		current = current.Next             // 移动到新节点
	}

	return head.Next // 返回头结点
}

func SwapPairs(head *L2stNode) *L2stNode {
	dummy := &L2stNode{Next: head}
	cur := head
	pre := dummy
	for cur != nil && cur.Next != nil {
		pre.Next = cur.Next
		cur.Next = cur.Next.Next
		pre.Next.Next = cur
		cur = cur.Next
		pre = pre.Next.Next
	}
	return dummy.Next
}
func Pr1ntList(head *L2stNode) {
	current := head
	for current != nil {
		fmt.Print(current.Val, " ")
		current = current.Next
	}
	fmt.Println()
}
func main() {
	L := Cr1ateList([]int{1, 2, 3, 4})
	L = SwapPairs(L)
	Pr1ntList(L)
}
