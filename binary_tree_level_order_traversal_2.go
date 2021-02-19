package main

import (
    "fmt"
)

type ListNode struct {
    val int
    left *ListNode
    right *ListNode
}

func main() {
    head := &ListNode{val:0}
    node1 := &ListNode{val:1}
    node2 := &ListNode{val:2}
    node3 := &ListNode{val:3}
	node4 := &ListNode{val:4}

	node5 := &ListNode{val:5}
	node6 := &ListNode{val:6}
	node7 := &ListNode{val:7}
	node8 := &ListNode{val:8}

    head.left = node1
	head.right = node2
	
    node1.left = node3
	node1.right = node4
	
	node2.left = node5
	node2.right = node6
	node6.right = node7

	node7.left = node8
    
    ans := levelTree(head)
    fmt.Println(ans)
}

func levelTree(head *ListNode) []int {
    ans := make([]int, 0)
    if head == nil {
         return ans   
    }
    stack := make([]*ListNode, 0)
    stack = append(stack, head)
    for len(stack) > 0 {
        l := len(stack)
        for i := 0; i < l; i++ {
            ans = append(ans, stack[i].val)
            if stack[i].left != nil {
                stack = append(stack, stack[i].left)
            }
            if stack[i].right != nil {
                stack = append(stack, stack[i].right)
            }
        }
        stack = stack[l:]
    }
    return ans
}