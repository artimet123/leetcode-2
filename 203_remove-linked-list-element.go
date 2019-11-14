package main 

import "fmt"

// 这是个简单的链表移动指定的数字的算法
type LinkNode struct{
	Val int
	Next *LinkNode
}
func main(){
	// 初始化
	head := &LinkNode{Val:0}
	node2 := &LinkNode{Val:2}
	node3 := &LinkNode{Val:3}
	node4 := &LinkNode{Val:4}
	head.Next = node2
	node2.Next = node3
	node3.Next = node4
	link := removeNode(head, 0)

	// 第一步是link是0	
	for l := link.Next; l != nil; l = l.Next{
		fmt.Println("l=", l.Val)
	}
}

func removeNode(head *LinkNode, val int)*LinkNode{
	ptr := &LinkNode{Next: head}

	cur, next := ptr, ptr.Next

	for next != nil {
		if next.Val == val {
			cur.Next = next.Next
		} else {
			cur = next
		}
		next = next.Next
	}
	return ptr
}