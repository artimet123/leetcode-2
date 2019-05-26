package main 

import "fmt"

type ListNode struct {
	Val int
	Next *ListNode
}

func main() {
	l1 := &ListNode{Val: 1}
	l2 := &ListNode{Val: 2}
	l3 := &ListNode{Val: 3}
	l4 := &ListNode{Val: 4}
	l5 := &ListNode{Val: 5}
	l6 := &ListNode{Val: 6}
	l7 := &ListNode{Val: 7}

	l1.Next = l2
	l2.Next = l3
	l3.Next = l4
	l4.Next = l5
	l5.Next = l6
	l6.Next = l7
	l := swapPair(l1)
	var vals []int
	for n := l; n != nil; n = n.Next {
		vals = append(vals, n.Val)
	}
	fmt.Println("vals=", vals)
}

func swapPair(head *ListNode) *ListNode {
	
	if head == nil || head.Next == nil {
		return head
	}
	next := head.Next
	head.Next = swapPair(next.Next)
	next.Next = head
	return next

}