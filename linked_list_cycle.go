package main 

import (
	"fmt"
)

type ListNode struct  {
    Val int;
    next *ListNode;
}


func main(){
	n1 := &ListNode{Val: 1}
	n2 := &ListNode{Val: 2}
	n3 := &ListNode{Val: 3}
	n1.next = n2
	n2.next = n3
	n3.next = n1

	b := hashCycle(n1)

	fmt.Println("b=", b)

}

func hashCycle(head *ListNode) bool {
	if head == nil {
		return false
	}

	p1, p2 := head,head

	for p2.next != nil && p2.next.next != nil {
		p1 = p1.next
		p2 = p2.next.next

		if p1 == p2 {
			return true
		}
	}
	return false
}