package main 

import "fmt"

type ListNode struct {
	Val int
	Next *ListNode
}

func main(){
	n1 := &ListNode{Val: 1}
	n2 := &ListNode{Val: 2}
	//n3 := &ListNode{Val: 3}
	//n4 := &ListNode{Val: 4}
	//n5 := &ListNode{Val: 5}
	n1.Next = n2
	//n2.Next = n3
	//n3.Next = n4
	//n4.Next = n5

	n := findNthFromEnd(n1, 3)
	fmt.Println("n=", n)
}

func findNthFromEnd(head *ListNode, n int) *ListNode {
	n1, n2 := head, head

	// 先让n2往后走n-1步
	for i := 0; i < n - 1; i++ {
		n2 = n2.Next
		if n2 == nil {
			fmt.Println(i)
			return head
		}
	}

	for n2.Next != nil {
		n1 = n1.Next
		n2 = n2.Next
	}
	return n1
}
