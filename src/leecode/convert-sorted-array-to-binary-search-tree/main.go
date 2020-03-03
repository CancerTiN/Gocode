package main

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func (t *TreeNode) LeftCollect(slice []int, tail bool) {
	if len(slice) > 0 {
		if tail {
			t.Val = slice[len(slice)-1]
			if len(slice[:len(slice)-1]) > 0 {
				t.Left = &TreeNode{}
				t.Left.LeftCollect(slice[:len(slice)-1], tail)
			}
		} else {
			t.Val = slice[0]
			if len(slice[1:]) > 0 {
				t.Left = &TreeNode{}
				t.Left.LeftCollect(slice[1:], tail)
			}
		}
	}
}

func leftAdd(nums []int) (root *TreeNode) {
	if len(nums) > 0 {
		root = &TreeNode{Val: nums[len(nums)-1]}
		root.Left = leftAdd(nums[:len(nums)-1])
	}
	return
}

func sortedArrayToBST(nums []int) *TreeNode {
	var root *TreeNode
	if len(nums) > 0 {
		i := len(nums) >> 1
		root = &TreeNode{Val: nums[i]}
		root.Left = sortedArrayToBST(nums[:i])
		root.Right = sortedArrayToBST(nums[i+1:])
	}
	return root
}

func main() {
	nums := []int{-10, -3, 0, 5, 9}
	root := sortedArrayToBST(nums)
	root.dlrDisplay()
}

func (t *TreeNode) dlrDisplay() {
	if t != nil {
		fmt.Println(t.Val)
		t.Left.dlrDisplay()
		t.Right.dlrDisplay()
	}
}
