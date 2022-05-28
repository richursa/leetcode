package main

import "fmt"

func main() {
	root := &TreeNode{5, &TreeNode{3, &TreeNode{2, &TreeNode{1, nil, nil}, nil}, &TreeNode{4, nil, nil}}, &TreeNode{6, nil, nil}}
	fmt.Println(kthSmallest(root))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var counter int

func kthSmallest(root *TreeNode, k int) int {
	counter = 0
	val, _ := inOrderRecursive(root, k)
	return val
}
func inOrderRecursive(root *TreeNode, k int) (int, bool) {
	if root == nil {
		return 0, false
	}
	if root.Left == nil {
		counter++
		if counter == k {
			return root.Val, true
		}
		return inOrderRecursive(root.Right, k)
	}
	val, ok := inOrderRecursive(root.Left, k)
	if ok {
		return val, ok
	}
	counter++
	if counter == k {
		return root.Val, true
	}
	return inOrderRecursive(root.Right, k)
}
