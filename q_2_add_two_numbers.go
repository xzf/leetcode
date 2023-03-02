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

//https://leetcode-cn.com/problems/add-two-numbers/
func addTwoNumbers(l1 *ListNode, l2 *ListNode) (head *ListNode) {
	var p *ListNode
	top := 0
	for {
		if l1 == nil && l2 == nil {
			break
		}
		val := top
		if l1 != nil {
			val += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			val += l2.Val
			l2 = l2.Next
		}
		top, val = val/10, val%10
		if head == nil {
			head = &ListNode{Val: val}
			p = head
		} else {
			p.Next = &ListNode{Val: val}
			p = p.Next
		}
	}
	if top != 0 {
		p.Next = &ListNode{Val: top}
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
}
