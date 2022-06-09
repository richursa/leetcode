// https://leetcode.com/problems/add-two-numbers/
package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	carry := 0
	newListNode := &ListNode{}
	clone := newListNode
	for l1 != nil && l2 != nil {
		newListNode.Val = (l1.Val + l2.Val + carry) % 10
		carry = (l1.Val + l2.Val + carry) / 10
		l1 = l1.Next
		l2 = l2.Next
		fmt.Println("loop1", newListNode.Val, carry)
		if l1 != nil || l2 != nil {
			newListNode.Next = &ListNode{}
			newListNode = newListNode.Next
		}
	}
	for l1 != nil {
		newListNode.Val = ((l1.Val + carry) % 10)
		carry = (l1.Val + carry) / 10
		fmt.Println("loop2", newListNode.Val, carry)
		l1 = l1.Next
		if l1 != nil {
			newListNode.Next = &ListNode{}
			newListNode = newListNode.Next
		}
	}
	for l2 != nil {
		newListNode.Val = ((l2.Val + carry) % 10)
		carry = (l2.Val + carry) / 10
		l2 = l2.Next
		fmt.Println("loop3", newListNode.Val, carry)
		if l2 != nil {
			newListNode.Next = &ListNode{}
			newListNode = newListNode.Next
		}
	}
	if carry != 0 {

		newListNode.Next = &ListNode{}
		newListNode.Next.Val = carry
	}
	return clone
}
