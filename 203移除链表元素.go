package main

// Definition for singly-linked list.

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

//func main() {
//	// 创建链表并打印
//	values := []int{1, 2, 6, 3, 4, 5, 6}
//	head := CreateList(values) // 使用输入数组创建链表
//	RemoveElements(head, 6)
//	PrintList(head) // 打印链表
//}
