// https://leetcode.com/problems/intersection-of-two-linked-lists/
package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	lengthA, lengthB := 0, 0
	headAClone := headA
	headBClone := headB
	for headAClone != nil {
		lengthA++
		headAClone = headAClone.Next
	}
	for headBClone != nil {
		lengthB++
		headBClone = headBClone.Next
	}
	if lengthA > lengthB {
		lengthA, lengthB = lengthB, lengthA
		headA, headB = headB, headA
	}
	diff := lengthB - lengthA
	headAClone = headA
	headBClone = headB
	for i := 0; i < diff; i++ {
		headBClone = headBClone.Next
	}
	for headAClone != nil && headBClone != nil {
		if headAClone == headBClone {
			return headAClone
		}
		headAClone = headAClone.Next
		headBClone = headBClone.Next
	}
	return nil
}
