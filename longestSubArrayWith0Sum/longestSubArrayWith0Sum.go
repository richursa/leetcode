// https://takeuforward.org/data-structure/length-of-the-longest-subarray-with-zero-sum/
package main

import "fmt"

func main() {
	fmt.Println(longestSubArrayWith0Sum([]int{15, -2, 2, -8, 1, 7, 10, 23}))
}

func longestSubArrayWith0Sum(nums []int) int {
	prefixMap := make(map[int]int)
	maxLength := 0
	currentSum := 0
	for i := 0; i < len(nums); i++ {
		currentSum += nums[i]
		if index, ok := prefixMap[currentSum]; ok {
			if (i - index) > maxLength {
				maxLength = i - index
			}
		} else {
			prefixMap[currentSum] = i
		}
	}
	return maxLength
}
