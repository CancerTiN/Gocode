package main

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var missLeaves bool

func getLDRSlice(root *TreeNode) (slice []int) {
	if root.Left == nil || root.Right == nil {
		if !(root.Left == nil && root.Right == nil) {
			missLeaves = true
		}
	}
	if root.Left != nil {
		slice = append(slice, getLDRSlice(root.Left)...)
	}
	slice = append(slice, root.Val)
	if root.Right != nil {
		slice = append(slice, getLDRSlice(root.Right)...)
	}
	return
}

func isSymmetricFail(root *TreeNode) bool {
	missLeaves = false
	if root == nil {
		return true
	}
	slice := getLDRSlice(root)
	if missLeaves {
		return false
	}
	for i := 0; i < len(slice); i++ {
		if i >= len(slice)-1-i {
			break
		}
		if slice[i] != slice[len(slice)-1-i] {
			return false
		}
	}
	return true
}

func backDown(p, q *TreeNode) (same bool) {
	if p.Val == q.Val {
		if p.Left != nil && q.Right != nil {
			same = backDown(p.Left, q.Right)
		} else if p.Left == nil && q.Right == nil {
			same = true
		}
		if !same {
			return
		}
		if p.Right != nil && q.Left != nil {
			same = backDown(p.Right, q.Left)
		} else if p.Right == nil && q.Left == nil {
			same = true
		} else {
			same = false
		}
	}
	return
}

func isSymmetric(root *TreeNode) bool {
	if root == nil || (root.Left == nil && root.Right == nil) {
		return true
	} else if root.Left == nil || root.Right == nil {
		return false
	}
	return backDown(root.Left, root.Right)
}

func main() {
	var root *TreeNode
	root = &TreeNode{Val: 1,
		Left:  &TreeNode{Val: 2, Left: &TreeNode{Val: 3}, Right: &TreeNode{Val: 4}},
		Right: &TreeNode{Val: 2, Left: &TreeNode{Val: 4}, Right: &TreeNode{Val: 3}}}
	fmt.Println(root, isSymmetric(root))
	root = &TreeNode{Val: 1,
		Left:  &TreeNode{Val: 2, Right: &TreeNode{Val: 3}},
		Right: &TreeNode{Val: 2, Right: &TreeNode{Val: 3}}}
	fmt.Println(root, isSymmetric(root))
	root = &TreeNode{Val: 1}
	fmt.Println(root, isSymmetric(root))
	root = &TreeNode{Val: 2,
		Left:  &TreeNode{Val: 3, Left: &TreeNode{Val: 4}, Right: &TreeNode{Val: 5}},
		Right: &TreeNode{Val: 3, Right: &TreeNode{Val: 4}}}
	fmt.Println(root, isSymmetric(root))
	root = &TreeNode{Val: 2,
		Left:  &TreeNode{Val: 3, Left: &TreeNode{Val: 4}, Right: &TreeNode{Val: 5}},
		Right: &TreeNode{Val: 3, Left: &TreeNode{Val: 5}}}
	fmt.Println(root, isSymmetric(root))
}
