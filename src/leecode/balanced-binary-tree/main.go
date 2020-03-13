package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func deep(root *TreeNode) (d int) {
	leftDeep, rightDeep := 0, 0
	if root.Left != nil {
		leftDeep = deep(root.Left) + 1
	}
	if root.Right != nil {
		rightDeep = deep(root.Right) + 1
	}
	if leftDeep > rightDeep {
		d = leftDeep
	} else {
		d = rightDeep
	}
	return
}

func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}
	leftDeep, rightDeep, deltaDeep := 0, 0, 0
	if root.Left != nil {
		leftDeep = deep(root.Left) + 1
	}
	if root.Right != nil {
		rightDeep = deep(root.Right) + 1
	}
	deltaDeep = leftDeep - rightDeep
	if -1 <= deltaDeep && deltaDeep <= 1 {
		leftIsBalanced, rightIsBalanced := true, true
		leftIsBalanced = isBalanced(root.Left)
		rightIsBalanced = isBalanced(root.Right)
		if leftIsBalanced && rightIsBalanced {
			return true
		} else {
			return false
		}
	} else {
		return false
	}
}

func main() {
	var root *TreeNode
	root = &TreeNode{Val: 3,
		Left:  &TreeNode{Val: 9},
		Right: &TreeNode{Val: 20, Left: &TreeNode{Val: 15}, Right: &TreeNode{Val: 7}},
	}
	fmt.Println(root, isBalanced(root))
	root = &TreeNode{Val: 1,
		Left: &TreeNode{Val: 2,
			Left:  &TreeNode{Val: 3, Left: &TreeNode{Val: 4}, Right: &TreeNode{Val: 4}},
			Right: &TreeNode{Val: 3}},
		Right: &TreeNode{Val: 2},
	}
	fmt.Println(root, isBalanced(root))
}
