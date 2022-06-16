// https://leetcode.com/problems/reverse-nodes-in-k-group/
package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	count := 0
	headClone := head
	var prevHead, newHead, prev, next, newKhead *ListNode
	firstHead := false
	prevHead = head
	for headClone != nil {
		walker := headClone
		for i := 0; i < k; i++ {
			if walker == nil {
				return newHead
			}
			walker = walker.Next
			if i == 0 {
				newKhead = walker
			}
		}

	}
	return newHead
}
