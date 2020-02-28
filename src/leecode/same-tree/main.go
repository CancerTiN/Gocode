package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSameTree(p *TreeNode, q *TreeNode) bool {
	if p == nil || q == nil {
		if p == q {
			return true
		} else {
			return false
		}
	}
	var b bool
	if p.Val == q.Val {
		b = isSameTree(p.Left, q.Left)
		if b {
			b = isSameTree(p.Right, q.Right)
		}
	}
	return b
}

func main() {
	var p *TreeNode
	var q *TreeNode
	p = &TreeNode{Val: 1, Left: &TreeNode{Val: 2}, Right: &TreeNode{Val: 3}}
	q = &TreeNode{Val: 1, Left: &TreeNode{Val: 2}, Right: &TreeNode{Val: 3}}
	fmt.Println(p, q, isSameTree(p, q))
	p = &TreeNode{Val: 1, Left: &TreeNode{Val: 2}}
	q = &TreeNode{Val: 1, Right: &TreeNode{Val: 3}}
	fmt.Println(p, q, isSameTree(p, q))
	p = &TreeNode{Val: 1, Left: &TreeNode{Val: 2}, Right: &TreeNode{Val: 1}}
	q = &TreeNode{Val: 1, Left: &TreeNode{Val: 1}, Right: &TreeNode{Val: 2}}
	fmt.Println(p, q, isSameTree(p, q))
}
