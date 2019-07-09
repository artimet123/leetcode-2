package main 

import "fmt"

type ListNode struct {
	Val int
	Next *ListNode
}

func main(){
	l1 := &ListNode{Val: 1}
	l2 := &ListNode{Val: 1}
	l3 := &ListNode{Val: 2}
	l4 := &ListNode{Val: 3}
	l5 := &ListNode{Val: 3}
	l1.Next = l2
	l2.Next = l3
	l3.Next = l4
	l4.Next = l5
	l := deleteDuplicates(l1)

	for n := l; n != nil; n = n.Next {
		fmt.Println("val=", n.Val)
	}
}

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	p1, p2 := head, head.Next

	for p2 != nil {
		if p1.Val == p2.Val {
			p1.Next, p2 = p2.Next, p2.Next
		} else {
			p1, p2 = p2, p2.Next
		}
	}
	return head
}

