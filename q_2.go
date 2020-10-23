package main

import "fmt"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	result := &ListNode{}
	head := result
	for {
		if l1 == nil && l2 == nil {
			break
		}
		if l1 != nil {
			result.Val += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			result.Val += l2.Val
			l2 = l2.Next
		}
		var val int
		if result.Val >= 10 {
			val = result.Val / 10
			result.Val = result.Val % 10
		}
		if l1 != nil || l2 != nil || val != 0 {
			result.Next = &ListNode{Val: val}
			result = result.Next
		}
	}
	return head
}
func main() {
	l1 := &ListNode{Val: 9,
		Next: &ListNode{Val: 9, Next: &ListNode{Val: 9, Next: &ListNode{Val: 9, Next: &ListNode{Val: 9, Next: &ListNode{Val: 9, Next: &ListNode{Val: 9}}}}}}}
	l2 := &ListNode{Val: 9, Next: &ListNode{Val: 9, Next: &ListNode{Val: 9, Next: &ListNode{Val: 9}}}}
	r := addTwoNumbers(l1, l2)
	for {
		if r == nil {
			return
		}
		fmt.Println(r.Val)
		r = r.Next
	}
	/*
		[2,7,11,15]
		9
	*/
}
