package main

import (
	"fmt"
	"strconv"
)

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
	t7 := &TreeNode{Val: 7}

	t8 := &TreeNode{Val: 8}
	t9 := &TreeNode{Val: 9}
	t0 := &TreeNode{Val: 0}

	t1.Left = t2
	t1.Right = t3

	t2.Left = t4
	t2.Right = t5

	t3.Left = t6
	t3.Right = t7

	t4.Left = t8
	t4.Right = t9

	t5.Right = t0
	count := sumNumbers(t1)
	fmt.Println("count=", count)

}

func sumNumbers(root *TreeNode) int {
	if root == nil {
		return 0
	}
	sum := 0
	for _, n := range pathNumber(root) {
		v, _ := strconv.Atoi(n)
		sum += v
	}
	return sum
}

// 中序遍历，字符串的累加，根到叶子才会形成一个数字
func pathNumber(root *TreeNode) (nums []string) {
	var number = map[int]string{
		0: "0", 1: "1", 2: "2", 3: "3", 4: "4", 5: "5", 6: "6", 7: "7", 8: "8", 9: "9",
	}

	fmt.Println("root=", root.Val)
	val := number[root.Val]
	if root.Left == nil && root.Right == nil {
		return []string{val}
	}
	if root.Left != nil {
		for _, n := range pathNumber(root.Left) {
			fmt.Println("left-n=", n, " val=", val)
			nums = append(nums, val+n)
			fmt.Println("left-nums=", nums)
		}
	}
	if root.Right != nil {
		for _, n := range pathNumber(root.Right) {
			fmt.Println("right-n=", n)
			nums = append(nums, val+n)
			fmt.Println("right-nums=", nums)
		}
	}
	return
}
