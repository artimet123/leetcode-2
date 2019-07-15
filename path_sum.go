package main

import "fmt"

/*
	这个是计算路径的，走过的路径的val值是否+的合 == sum
*/
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
	t5 := &TreeNode{Val: 0}
	t6 := &TreeNode{Val: 3}

	t1.Left = t2
	t1.Right = t3

	t2.Right = t4
	t3.Left = t5

	t5.Right = t6

	ok := hasPathSum(t1, 7)
	fmt.Println("ok=", ok)

}

func hasPathSum(root *TreeNode, sum int) bool {
	if root == nil {
		return false
	}
	fmt.Println("sum=", sum, " val=", root.Val)
	if sum -= root.Val; root.Left == nil && root.Right == nil {
		return sum == 0
	}

	fmt.Println("sum=", sum, " val=", root.Val)
	if root.Left != nil && hasPathSum(root.Left, sum) {
		return true
	}

	if root.Right != nil && hasPathSum(root.Right, sum) {
		return true
	}
	return false
}
