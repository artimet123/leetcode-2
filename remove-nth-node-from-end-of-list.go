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
	n5 := &ListNode{Val: 5}
	n1.Next = n2
	n2.Next = n3
	n3.Next = n4
	n4.Next = n5

	n := removeNthFromEnd(n1, 3)
	fmt.Println("n=", n)
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	n1, n2 := head, head
	for i := 0; i < n; i++ {
		n2 = n2.Next
	}

	if n2 == nil {
		return head.Next
	}
	for n2.Next != nil {
		n1 = n1.Next
		n2 = n2.Next
	}
	n1.Next = n1.Next.Next
	return head
}
