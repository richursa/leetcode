// https://leetcode.com/problems/linked-list-cycle-ii/
package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func detectCycle(head *ListNode) *ListNode {
	slowPointer := head
	fastPointer := head
	for slowPointer != nil && fastPointer != nil && fastPointer.Next != nil {
		slowPointer = slowPointer.Next
		fastPointer = fastPointer.Next.Next
		if slowPointer == fastPointer {
			for slowPointer != head {
				slowPointer = slowPointer.Next
				head = head.Next
			}
			return slowPointer
		}
	}
	return nil
}
