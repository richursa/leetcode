// https://leetcode.com/problems/rotate-list/description/
package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func rotateRight(head *ListNode, k int) *ListNode {
	length := 0
	headClone := head
	for headClone != nil {
		length++
		headClone = headClone.Next
	}
	if length == 0 {
		return head
	}
	k = k % length
	if k == 0 {
		return head
	}
	headClone = head
	for i := 1; i < length-k; i++ {
		headClone = headClone.Next
	}
	temp := headClone.Next
	headClone.Next = nil
	newHead := temp
	for temp.Next != nil {
		temp = temp.Next
	}
	temp.Next = head
	return newHead
}
