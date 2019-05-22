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
	l8 := &ListNode{Val: 8}
	l1.Next = l3
	l3.Next = l5
	l5.Next = l7
	l2.Next = l4
	l4.Next = l6
	l6.Next = l8
	l := mergeTwoLists(l1, l2)
	var vals []int
	for n := l; n != nil; n = n.Next {
		vals = append(vals, n.Val)
	}
	fmt.Println("vals=", vals)
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	var h ListNode
	l := &h
	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			l.Next, l1 = l1, l1.Next
		} else {
			l.Next, l2 = l2, l2.Next
		}
		l = l.Next
	}
	if l1 != nil {
		l.Next = l1
	} else if l2 != nil {
		l.Next = l2
	}
	return h.Next
}
