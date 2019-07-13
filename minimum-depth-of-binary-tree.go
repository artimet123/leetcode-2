package main 

import "fmt"

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func main(){
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

        l := minDepth(t1)
        fmt.Println("l=", l)

}

func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	if root.Left == nil && root.Right == nil {
		return 1
	}
	if root.Left == nil {
		return minDepth(root.Right) + 1
	}
	if root.Right == nil {
		return minDepth(root.Left) + 1
	}

	l, r := minDepth(root.Left), minDepth(root.Right)
	if l < r {
		return l + 1
	}
	return r + 1
}