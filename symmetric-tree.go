package main

import "fmt"

// 计算二叉树是否对称
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	t1 := &TreeNode{Val: 1}

	t2 := &TreeNode{Val: 2}
	t22 := &TreeNode{Val: 2}


	t3 := &TreeNode{Val: 3}
	t4 := &TreeNode{Val: 4}

	t1.Left = t2
	t1.Right = t22

	t2.Left = t3
	t2.Right = t4

	t22.Right = t3
	t22.Left = t4

	b := isSymmetric(t1)
	fmt.Println("b=", b)

}

func isSymmetric(root *TreeNode)bool{
	if root == nil {
		return true
	}

	return same(root.Left, root.Right)
}

func same(left, right *TreeNode) bool {
	if left == nil && right == nil {
		return true
	}

	if left == nil || right == nil {
		return false
	}

	if left.Val != right.Val {
		return false
	}

	return same(left.Left, right.Right) && same(left.Right, right.Left)
}
