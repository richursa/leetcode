// https://leetcode.com/problems/reverse-nodes-in-k-group/
package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	currentLength := 0
	headClone := head
	newHead := head
	lastTail := head
	for headClone != nil {
		if currentLength%k == 0 {
			currentHead, tail, next := reverseKNodes(headClone, k)
			if currentLength == 0 {
				newHead = currentHead
				lastTail = tail
			} else {
				lastTail.Next = currentHead
				lastTail = tail
			}
			currentLength = currentLength + k
			headClone = next
		}
	}
	return newHead
}

func reverseKNodes(head *ListNode, k int) (*ListNode, *ListNode, *ListNode) {
	headClone := head
	tail := headClone
	length := 0
	for headClone != nil && length < k {
		length++
		tail = headClone
		headClone = headClone.Next
	}
	if length < k {
		return head, tail, nil
	}
	headClone = head
	var prev, temp *ListNode
	currentLength := 0
	for headClone != nil && currentLength < k {
		temp = headClone.Next
		headClone.Next = prev
		prev = headClone
		headClone = temp
		currentLength++
	}
	return prev, head, temp
}
