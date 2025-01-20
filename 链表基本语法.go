package main

import "fmt"

// 设置节点的结构体
type Node struct {
	Value int
	Next  *Node
}

// 设置单链表的结构体，head指向头节点，Tail指向尾节点
type LinkList struct {
	Head *Node
	Tail *Node
}

func NewLinkList() *LinkList {
	Head := &Node{}
	Tail := Head
	return &LinkList{Head: Head, Tail: Tail}
}

func TailInsert(L *LinkList, val int) {
	NewNode := &Node{Value: val}
	L.Tail.Next = NewNode
	L.Tail = NewNode
}

func DeleteByVal(L *LinkList, val int) {
	p := L.Head
	q := L.Head.Next
	for q != nil {

		if q.Value != val {
			p = q
			q = q.Next
		} else {
			if q.Next == nil {
				L.Tail = p
				p.Next = nil
				q = nil
			} else {
				p.Next = q.Next
				q = q.Next
			}
		}
	}
}

func PrintList(L *LinkList) {
	p := L.Head.Next
	for p != nil {
		fmt.Println(p.Value)
		p = p.Next
	}
}

func main() {
	L := NewLinkList()
	list := []int{1, 2, 6, 3, 4, 5, 6}

	for _, val := range list {
		TailInsert(L, val)
	}
	PrintList(L)
	DeleteByVal(L, 6)
	PrintList(L)
}
