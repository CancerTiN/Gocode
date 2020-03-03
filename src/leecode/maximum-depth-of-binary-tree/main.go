package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func recursiveDrillDown(root *TreeNode, iDeep int) (oDeep int) {
	if root == nil {
		oDeep = iDeep - 1
	} else {
		lDeep := recursiveDrillDown(root.Left, iDeep+1)
		rDeep := recursiveDrillDown(root.Right, iDeep+1)
		if lDeep > rDeep {
			oDeep = lDeep
		} else {
			oDeep = rDeep
		}
	}
	return
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return recursiveDrillDown(root, 1)
}

func main() {
	var root *TreeNode
	root = &TreeNode{Val: 3,
		Left:  &TreeNode{Val: 9},
		Right: &TreeNode{Val: 20, Left: &TreeNode{Val: 15}, Right: &TreeNode{Val: 7}}}
	fmt.Println(root, maxDepth(root))
}
