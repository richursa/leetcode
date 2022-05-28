package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sortedArrayToBST(nums []int) *TreeNode {
	root := &TreeNode{}
	recursiveBST(nums, 0, len(nums)-1, root)
	return root
}
func recursiveBST(nums []int, left, right int, root *TreeNode) {
	if left > right {
		return
	}
	if left == right {
		root.Val = nums[left]
		return
	}
	middle := (left + right) / 2
	root.Val = nums[middle]

	if left <= middle-1 {
		root.Left = &TreeNode{}
		recursiveBST(nums, left, middle-1, root.Left)
	}
	if right >= middle+1 {
		root.Right = &TreeNode{}
		recursiveBST(nums, middle+1, right, root.Right)
	}

}
func main() {

}
