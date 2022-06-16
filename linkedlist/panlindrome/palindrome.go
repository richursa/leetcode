// https://leetcode.com/problems/palindrome-linked-list/
package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func isPalindrome(head *ListNode) bool {
	slow, fast := head, head

	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	var prev *ListNode
	for slow != nil {
		tmp := slow.Next
		slow.Next = prev
		prev = slow
		slow = tmp
	}

	for prev != nil && head != nil {
		if prev.Val != head.Val {
			return false
		}
		prev = prev.Next
		head = head.Next
	}

	return true
}
