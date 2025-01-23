package main

func getIntersectionNode(head1 *ListNode, head2 *ListNode) *ListNode {
	dummy1 := &ListNode{Next: head1}
	dummy2 := &ListNode{Next: head2}
	Len1 := 0
	Len2 := 0
	cur1 := dummy1
	cur2 := dummy2
	for LNode := dummy1; LNode.Next != nil; LNode = LNode.Next {
		Len1++
	}
	for LNode := dummy2; LNode.Next != nil; LNode = LNode.Next {
		Len2++
	}
	if Len1 < Len2 {
		for i := 0; i < Len2-Len1; i++ {
			cur2 = cur2.Next
		}
	} else {
		for i := 0; i < Len1-Len2; i++ {
			cur1 = cur1.Next
		}
	}
	for cur1 != nil {
		if cur2 == cur1 {
			return cur2
		} else {
			cur2 = cur2.Next
			cur1 = cur1.Next
		}
	}
	return nil
}
