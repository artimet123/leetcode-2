package main 

import (
	"fmt"
)

// 二叉树的中序遍历是:左->中->右
type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func main(){

	t1 := newNode(1)
	t2 := newNode(2)
	t3 := newNode(3)
	t4 := newNode(4)
	t5 := newNode(5)
	t6 := newNode(6)
	t7 := newNode(7)
	t8 := newNode(8)
	t9 := newNode(9)

	t1.setLeft(t2)
	t1.setRight(t3)
	t2.setLeft(t4)
	t2.setRight(t5)
	t3.setLeft(t6)
	t3.setRight(t7)
	t4.setLeft(t8)
	t4.setRight(t9)

	ret := inorderTraversal(t1)

	fmt.Println("ret=", ret)



}

func newNode(val int)*TreeNode{
	return &TreeNode{Val: val}
}

func (t *TreeNode)setLeft(left *TreeNode){
	t.Left = left
}

func (t *TreeNode)setRight(right * TreeNode){
	t.Right = right
}



func inorderTraversal(root *TreeNode) []int{
	if root == nil {
		return nil
	}

	ret := []int{}

	if left := inorderTraversal(root.Left); left != nil {
		ret = append(ret, left...)
	}

	ret = append(ret, root.Val)
	if right := inorderTraversal(root.Right); right != nil {
		ret = append(ret, right...)
	}

	return ret
}