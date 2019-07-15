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

	ret := PathSum(t1, 7)

	fmt.Println("ret=", ret)
}

func PathSum(root *TreeNode, sum int) [][]int {
	if root == nil {
		return nil
	}

	if sum -= root.Val; root.Left == nil && root.Right == nil {
		if sum == 0 {
			return [][]int{{root.Val}}
		}
		return nil
	}

	var ret [][]int

	if root.Left != nil {
		ret = append(ret, PathSum(root.Left, sum)...)
	}

	if root.Right != nil {
		ret = append(ret, PathSum(root.Right, sum)...)
	}

	for i := 0; i < len(ret); i++ {
		ret[i] = append([]int{root.Val}, ret[i]...)
	}
	return ret
}
