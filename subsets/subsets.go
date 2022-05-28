package main

import "fmt"

func main() {
	fmt.Println(subsets([]int{1, 2, 3}))
}

var sets [][]int
var current []int

func subsets(nums []int) [][]int {
	sets = [][]int{}
	current = make([]int, len(nums)+1)
	recursiveSubsetCreator(nums, 0, 0)
	return sets
}

func recursiveSubsetCreator(nums []int, index int, currlength int) {
	if index >= len(nums) {
		temp := make([]int, len(current[:currlength]))
		copy(temp, current[:currlength])
		sets = append(sets, temp)
		return
	}
	recursiveSubsetCreator(nums, index+1, currlength)
	current[currlength] = nums[index]
	recursiveSubsetCreator(nums, index+1, currlength+1)
}
