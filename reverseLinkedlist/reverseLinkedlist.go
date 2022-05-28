package main

func main() {}

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	return recursiveReverse(head, nil)
}

func recursiveReverse(current *ListNode, prev *ListNode) *ListNode {
	if current == nil {
		return nil
	}
	if current.Next == nil {
		current.Next = prev

		return current
	}
	if prev == nil {
		head := recursiveReverse(current.Next, current)
		current.Next = nil
		return head
	}
	head := recursiveReverse(current.Next, current)
	current.Next = prev
	return head
}
