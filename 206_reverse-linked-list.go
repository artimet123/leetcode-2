package main 

import "fmt"

type ListNode struct {
	Val int
	Next *ListNode
}

func main(){
	n1 := &ListNode{Val: 1}
	n2 := &ListNode{Val: 2}
	n3 := &ListNode{Val: 3}
	n4 := &ListNode{Val: 4}
	n1.Next = n2
	n2.Next = n3
	n3.Next = n4

	n := reverseList(n1)
	if n.Val != 4 {
		fmt.Println("n=", n.Val)
	}
	return
}

func reverseList(head *ListNode) *ListNode {
	
	if head == nil || head.Next == nil {
		// 没有或只有一个元素时，不需要反转
		return head
	}

	node, next := head, head.Next
	for next != nil {
		// 1: 先记录下来next.Next的地址
		tmp := next.Next
		// 2: 换next和head的地址
		next.Next = node
		node = next // 这步非常关键，会让node 始终指向开头的元素

		// 3: 移动next
		next = tmp
	}
	head.Next = nil // 这个步骤的意义？
	return node
}