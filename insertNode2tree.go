package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 最长路径 = 左子树 + 右子树 +1（根）
func main() {
	//var l ,r int
	t1 := &TreeNode{Val: 6}
	t2 := &TreeNode{Val: 2}
	t3 := &TreeNode{Val: 3}
	t4 := &TreeNode{Val: 4}
	t5 := &TreeNode{Val: 5}
	t6 := &TreeNode{Val: 7}

	t1.Left = t2
	t1.Right = t3

	t2.Left = t4
	t2.Right = t5

	t5.Right = t6
	InsertNode2BSortTree(t1, 8)

	valArray := preorderTravesal(t1)
	fmt.Println("val=", valArray)
}

func InsertNode2BSortTree(root *TreeNode, data int) *TreeNode {
	if root == nil {
		//建立一个新节点
		pTNode := &TreeNode{Val: data}
		root = pTNode
		fmt.Println("root=", root.Val)
		return root
	} else if root.Val > data {
		fmt.Println("1root=", root.Val)
		root.Left = InsertNode2BSortTree(root.Left, data)
	} else {
		fmt.Println("2root=", root.Val)
		root.Right = InsertNode2BSortTree(root.Right, data)
		fmt.Println("3root=", root.Right.Val)
	}
	return root
}

func preorderTravesal(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	ret := append([]int{}, root.Val)

	if left := preorderTravesal(root.Left); left != nil {
		ret = append(ret, left...)
	}

	if right := preorderTravesal(root.Right); right != nil {
		ret = append(ret, right...)
	}

	return ret
}
