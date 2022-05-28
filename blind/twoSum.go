// https://leetcode.com/problems/two-sum/
// 1. Two Sum
// Easy

// Given an array of integers nums and an integer target, return indices of the two numbers such that they add up to target.

// You may assume that each input would have exactly one solution, and you may not use the same element twice.

// You can return the answer in any order.

// Example 1:

// Input: nums = [2,7,11,15], target = 9
// Output: [0,1]
// Output: Because nums[0] + nums[1] == 9, we return [0, 1].
// Example 2:

// Input: nums = [3,2,4], target = 6
// Output: [1,2]
// Example 3:

// Input: nums = [3,3], target = 6
// Output: [0,1]

// Constraints:

// 2 <= nums.length <= 104
// -109 <= nums[i] <= 109
// -109 <= target <= 109
// Only one valid answer exists.

// Follow-up: Can you come up with an algorithm that is less than O(n2) time complexity?
package main

func twoSum(nums []int, target int) []int {
	numMap := make(map[int][]int)
	for i := 0; i < len(nums); i++ {
		numMap[nums[i]] = append(numMap[nums[i]], i)
	}
	for i := 0; i < len(nums); i++ {
		if _, ok := numMap[target-nums[i]]; ok {
			if nums[i] == target/2 {
				if len(numMap[nums[i]]) == 2 {
					return numMap[nums[i]]
				}
			} else {
				return []int{i, numMap[target-nums[i]][0]}
			}
		}
	}
	return []int{}
}
