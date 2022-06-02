package main

import (
	"fmt"
	"sort"
)

// https://leetcode.com/problems/4sum/

func main() {
	fmt.Println(fourSum([]int{1, -5, 1, -4, 2, 1, -3}, 1))
}
func fourSum(nums []int, target int) [][]int {
	currentI := -1000000001
	result := [][]int{}
	sort.Ints(nums)
	for i := 0; i < len(nums)-2; i++ {
		currentJ := -1000000001
		if currentI != nums[i] {
			currentI = nums[i]
			for j := i + 1; j < len(nums)-1; j++ {
				if currentJ != nums[j] {
					currentJ = nums[j]
					twoSums := twosum(nums[j+1:], target-nums[i]-nums[j])
					for k := 0; k < len(twoSums); k++ {
						result = append(result, append([]int{nums[i], nums[j]}, twoSums[k]...))
					}
				}
			}
		}
	}
	return result
}

func twosum(nums []int, target int) [][]int {
	currentI := -1000000001
	result := [][]int{}
	for i := 0; i < len(nums)-1; i++ {
		if currentI != nums[i] {
			currentI = nums[i]
			index := sort.SearchInts(nums[i+1:], target-nums[i])
			if len(nums[i+1:]) != index && nums[i+1+index] == (target-nums[i]) {
				result = append(result, []int{nums[i], target - nums[i]})
			}
		}
	}
	return result
}
