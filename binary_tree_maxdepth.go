package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	t1 := &TreeNode{Val: 1}
	t2 := &TreeNode{Val: 2}
	t3 := &TreeNode{Val: 3}
	t4 := &TreeNode{Val: 4}
	t5 := &TreeNode{Val: 5}
	t6 := &TreeNode{Val: 6}

	t1.Left = t2
	t1.Right = t3

	t2.Right = t4
	t3.Left = t5

	t5.Right = t6
	ret := maxDepth(t1)
	fmt.Println("ret=", ret)
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	var l, r int

	if root.Left != nil {
		l = maxDepth(root.Left)
	}

	if root.Right != nil {
		r = maxDepth(root.Right)
	}

	if l > r {
		return l + 1
	}

	return r + 1
}
