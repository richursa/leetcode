package main

func main() {

}

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	number := 0
	temp := *l1
	for temp != nil {
		number = temp.Val
	}
}
