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
	ret := levelOrder(t1)
	fmt.Println("ret=", ret)
}

func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}

	var ret [][]int

	for child := []*TreeNode{root}; child != nil; child = levelSubTreeNode(child) {
		vals := make([]int, len(child))
		for i, n := range child {
			vals[i] = n.Val
		}
		ret = append(ret, vals)
	}
	return ret
}

func levelSubTreeNode(nodes []*TreeNode) (child []*TreeNode) {
	for _, n := range nodes {
		if n.Left != nil {
			child = append(child, n.Left)
		}
		if n.Right != nil {
			child = append(child, n.Right)
		}
	}
	return
}
