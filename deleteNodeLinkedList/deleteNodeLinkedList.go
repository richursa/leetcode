package main

func main() {

}

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteNode(node *ListNode) {
	if node.Next.Next == nil {
		node.Val = node.Next.Val
		node.Next = nil
		return
	}
	node.Val = node.Next.Val
	deleteNode(node.Next)
}
