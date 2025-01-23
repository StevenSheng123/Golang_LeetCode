package main

func removeNthFromEnd(head *ListNode, N int) *ListNode {
	dummy := &ListNode{Next: head}
	pre, mov := dummy, dummy

	for i := 0; i < N; i++ {
		mov = mov.Next
	}
	for mov.Next != nil {
		pre = pre.Next
		mov = mov.Next
	}
	pre.Next = pre.Next.Next
	return dummy.Next
}

//func main() {
//	L := CreateList([]int{1, 2, 3, 4, 5})
//	L = removeNthFromEnd(L, 2)
//	PrintList(L)
//
//}
