package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	zero := &ListNode{Next: head}
	for temp := zero.Next; temp.Next != nil; {
		if temp.Val == temp.Next.Val {
			temp.Next = temp.Next.Next
		} else {
			temp = temp.Next
		}
	}
	return zero.Next
}

func main() {
	zero := &ListNode{}
	zero.Next = &ListNode{Val: 1, Next: &ListNode{Val: 1, Next: &ListNode{Val: 2}}}
	zero.Next.Display()
	deleteDuplicates(zero.Next)
	zero.Next.Display()
	zero.Next = &ListNode{Val: 1, Next: &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 3}}}}}
	zero.Next.Display()
	deleteDuplicates(zero.Next)
	zero.Next.Display()
}

func (head *ListNode) Display() {
	fmt.Print(head.Val)
	if head.Next != nil {
		fmt.Print("->")
		head.Next.Display()
	} else {
		fmt.Println()
		return
	}
}
