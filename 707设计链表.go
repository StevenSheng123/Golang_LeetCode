package main

import "fmt"

type myNode struct {
	val  int
	Next *myNode
}

type MyLinkedList struct {
	head *myNode
}

func Constructor() MyLinkedList {
	head := &myNode{}
	return MyLinkedList{head: head}
}

func (this *MyLinkedList) Get(index int) int {
	var a int
	cur := this.head
	if index < 0 {
		return -1
	}
	for i := 0; i <= index; i++ {
		if cur.Next != nil {
			cur = cur.Next
		} else {
			return -1
		}
	}
	a = cur.val
	return a
}

func (this *MyLinkedList) AddAtHead(val int) {
	newNode := &myNode{val: val}
	newNode.Next = this.head.Next
	this.head.Next = newNode
}

func (this *MyLinkedList) AddAtTail(val int) {
	cur := this.head
	for cur.Next != nil {
		cur = cur.Next
	}
	newNode := &myNode{val: val, Next: nil}
	cur.Next = newNode
}

func (this *MyLinkedList) AddAtIndex(index int, val int) {
	cur := this.head
	newNode := &myNode{val: val}
	//index在链表中，则cur.Next != nil
	for i := 0; i < index; i++ {

		//遍历index次，cur指向待插入节点的前一个节点
		if cur.Next == nil {
			//此时已经指到了最后一个节点
			//若i==index-2则插入最后
			if i == index {
				cur.Next = newNode
			} else {
				//i!=index,则index大于现有链表长度，不插入
				fmt.Println("插入失败")
				return
			}
		}
		cur = cur.Next //cur指向下一个节点
	}
	newNode.Next = cur.Next
	cur.Next = newNode
}

func (this *MyLinkedList) DeleteAtIndex(index int) {
	cur := this.head //cur始终指向待删除节点的前一个节点
	for i := 0; i < index; i++ {
		if cur.Next == nil {
			//此时cur已经来到了链表的末尾,删除失败
			return
		} else {
			cur = cur.Next
		}
	}
	if cur.Next != nil {
		cur.Next = cur.Next.Next
	}
}
func (this *MyLinkedList) Print() {
	cur := this.head
	for cur.Next != nil {
		cur = cur.Next
		fmt.Println(cur.val)
	}
}

/**
 * Your MyLinkedList object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Get(index);
 * obj.AddAtHead(val);
 * obj.AddAtTail(val);
 * obj.AddAtIndex(index,val);
 * obj.DeleteAtIndex(index);
 */
//func main() {
//	// Initialize a new linked list
//	obj := Constructor()
//
//	// Add nodes at head and tail
//
//	obj.AddAtIndex(2, 1) // 插入失败
//	obj.AddAtIndex(3, 4) // 插入失败
//	obj.AddAtTail(1)
//	obj.Get(0)
//	obj.DeleteAtIndex(0)
//	obj.Get(0)
//	//obj.Print()
//}
