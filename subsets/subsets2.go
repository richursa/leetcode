package main

func main() {

}

var sets [][]int

func subsetsWithDup(nums []int) [][]int {
	if len(nums) == 1 {
		return
	}
}

func recursiveSubsetsWithDup(nums []int) {
	if len(nums) == 1 {
		sets = append(sets, nums)
		return
	}
	recursiveSubsetsWithDup(nums[1:])
}
