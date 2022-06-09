// https://leetcode.com/problems/remove-nth-node-from-end-of-list/
package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {

}
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	clone := head
	length := 0
	for head != nil {
		length++
		head = head.Next
	}
	if length == n {
		return clone.Next
	}
	head = clone
	for i := length; i >= 0 && head != nil; i-- {
		if i == n+1 {
			head.Next = head.Next.Next
		}
		head = head.Next
	}
	return clone
}
