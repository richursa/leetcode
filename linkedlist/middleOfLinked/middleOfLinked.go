// https://leetcode.com/problems/middle-of-the-linked-list/
/**
 * Definition for singly-linked list.
 */
package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func middleNode(head *ListNode) *ListNode {

	slow := head
	fast := head

	for fast != nil && fast.Next != nil {
		slow = slow.Next		
		fast = fast.Next.Next
	}

	return slow
}
