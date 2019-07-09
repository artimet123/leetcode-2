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
	t1.Left = t2
	t2.Left = t3
	t3.Right = t4

	t21 := &TreeNode{Val: 1}
	t22 := &TreeNode{Val: 2}
	t23 := &TreeNode{Val: 3}
	t24 := &TreeNode{Val: 4}
	t21.Left = t22
	t22.Left = t23
	t23.Right = t24

	b := IsSameTree(t1, t21)
	fmt.Println("b=", b)

}

func IsSameTree(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}

	if p != nil && q != nil && p.Val == q.Val {
		return IsSameTree(p.Left, q.Left) && IsSameTree(p.Right, q.Right)
	}
	return false
}
