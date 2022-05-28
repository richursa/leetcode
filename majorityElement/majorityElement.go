package main

//https://leetcode.com/problems/majority-element/submissions/
//Boyer Moore Voting Algorithm
func majorityElement(nums []int) int {
	m := nums[0]
	count := 1
	for i := 1; i < len(nums); i++ {
		if count == 0 {
			m = nums[i]
			count++
			continue
		}
		if nums[i] == m {
			count++
			continue
		}
		count--
	}
	return m
}
