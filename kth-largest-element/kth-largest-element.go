//https://leetcode.com/problems/kth-largest-element-in-an-array/
package main

func main() {

}
var kLargest []int
type node struct {
	left  *node
	right *node
	val   int
}

func (n *node) insert(val int) {
	if n == nil {
		n = &node{nil, nil, val}
	} else if n.left.val <0= val {
		n.left.insert(val)
	} else {
		n.right.insert(val)
	}
}
var count int
func (n *node) getKLargest(k int){
	if n==nil || count ==k{
		return
	} 
	if n.right == nil {
		kLargest = append(kLargest, n.val)
		count++
	}
	n.right.getKLargest(k)
	n.left.getKLargest(k)
}
func findKthLargest(nums []int, k int) int {

}
