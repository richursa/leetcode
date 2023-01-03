// https://leetcode.com/problems/missing-number/
package main

func main() {

}
func missingNumber(nums []int) int {
	n := len(nums)
	requiredSum := n * (n + 1) / 2
	actualSum := 0
	for _, elem := range nums {
		actualSum += elem
	}
	return requiredSum - actualSum
}
