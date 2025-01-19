package main

type Node struct {
	Value int
	Next  *Node
}

func NewLinkedList() *Node {
	//创建头节点
	return &Node{}
}

func (head *Node) Add(val int) {
	newNode := &Node{Value: val}
	current := head
}
