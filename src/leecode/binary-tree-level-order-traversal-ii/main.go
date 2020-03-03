package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelAppend(root *TreeNode, level int, iSlice [][]int) (oSlice [][]int) {
	if root != nil {
		for level > len(iSlice)-1 {
			iSlice = append(iSlice, make([]int, 0))
		}
		iSlice[level] = append(iSlice[level], root.Val)
		iSlice = levelAppend(root.Left, level+1, iSlice)
		iSlice = levelAppend(root.Right, level+1, iSlice)
	}
	oSlice = iSlice
	return
}

func levelOrderBottom(root *TreeNode) [][]int {
	mSlice := make([][]int, 0)
	if root != nil {
		mSlice = levelAppend(root, 0, mSlice)
	}
	rSlice := make([][]int, len(mSlice))
	for i, slice := range mSlice {
		rSlice[len(rSlice)-1-i] = slice
	}
	return rSlice
}

func main() {
	var root *TreeNode
	root = &TreeNode{Val: 3,
		Left:  &TreeNode{Val: 9},
		Right: &TreeNode{Val: 20, Left: &TreeNode{Val: 15}, Right: &TreeNode{Val: 7}}}
	fmt.Println(root, levelOrderBottom(root))
}
