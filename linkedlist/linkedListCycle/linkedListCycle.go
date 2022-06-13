// https://leetcode.com/problems/linked-list-cycle/
package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func hasCycle(head *ListNode) bool {
	cycleDetected := false
	slowPointer := head
	fastPointer := head
	for slowPointer != nil && fastPointer != nil && fastPointer.Next != nil {
		slowPointer = slowPointer.Next
		fastPointer = fastPointer.Next.Next
		if slowPointer == fastPointer {
			cycleDetected = true
			break
		}
	}
	return cycleDetected
}
