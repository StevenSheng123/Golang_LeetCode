package main

import "fmt"

func detectCycle(head *ListNode) *ListNode {
	//一边遍历一边创建一个map,key为节点的地址,value为索引
	m := make(map[*ListNode]int)
	dummy := &ListNode{Next: head}
	cur := dummy
	i := -1
	for cur != nil {
		_, exists := m[cur]
		if exists {
			return cur
		}
		m[cur] = i
		i++
	}
	return nil
}
func CreateCircularList(values []int) *ListNode {
	if len(values) == 0 {
		return nil
	}

	head := &ListNode{Val: values[0]}
	current := head
	for _, value := range values[1:] {
		current.Next = &ListNode{Val: value}
		current = current.Next
	}
	// 形成环形链表，最后一个节点指向头节点
	current.Next = head.Next
	return head
}
func main() {
	l := CreateCircularList([]int{3, 2, 0, -4})
	//PrintList(l)
	Node := detectCycle(l)
	fmt.Println(Node == l.Next)
}
