package main 

import "fmt"

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func main(){
	var nArray =[]int{0,1,2,3,4,6,7,8,9,10}
	t1 := sortedArray2BST(nArray)

	fmt.Println("root.val=", t1.Val)
	fmt.Println("root.Left=", t1.Left.Val, " root.Right=", t1.Right.Val)

}

func sortedArray2BST(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}

	m := len(nums) / 2
	root := &TreeNode{Val: nums[m]}

	if 0 < m {
		root.Left = sortedArray2BST(nums[:m])
	}

	if m < len(nums) - 1 {
		root.Right = sortedArray2BST(nums[m+1:])
	}

	return root
}