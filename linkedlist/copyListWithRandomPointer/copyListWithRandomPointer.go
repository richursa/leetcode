//https://leetcode.com/problems/copy-list-with-random-pointer/
package main

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

func copyRandomList(head *Node) *Node {
	if head == nil {
		return nil
	}
	pointerMap := make(map[*Node]*Node)
	headClone := head
	var newHead *Node
	if head.Next == nil {
		newHead = &Node{Val: head.Val, Next: nil}
	} else {
		newHead = &Node{Val: head.Val, Next: &Node{}}
	}
	newHeadClone := newHead
	pointerMap[head] = newHead
	pointerMap[head.Next] = newHead.Next
	headClone = headClone.Next
	newHeadClone = newHead.Next
	for headClone != nil {
		newHeadClone.Val = headClone.Val
		if headClone.Next != nil {
			newHeadClone.Next = &Node{}
		}
		pointerMap[headClone.Next] = newHeadClone.Next
		headClone = headClone.Next
		newHeadClone = newHeadClone.Next
	}
	headClone = head
	newHeadClone = newHead
	for headClone != nil {
		newHeadClone.Random = pointerMap[headClone.Random]
		newHeadClone = newHeadClone.Next
		headClone = headClone.Next
	}
	return newHead
}
