package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	tree := &TreeNode{Val: 1, Left: nil, Right: &TreeNode{Val: 2, Left: &TreeNode{Val: 3, Left: nil, Right: nil}, Right: nil}}
	fmt.Println(inorderTraversal(tree))

}

var list []int

func inorderRecursive(root *TreeNode) {
	if root == nil {
		return
	}
	inorderRecursive(root.Left)
	list = append(list, root.Val)
	inorderRecursive(root.Right)
}

func inorderTraversal(root *TreeNode) []int {
	list = []int{}
	inorderRecursive(root)
	return list
}
