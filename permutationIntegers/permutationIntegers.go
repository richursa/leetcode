package main

import "fmt"

var bigArray [][]int

func main() {
	nums := []int{1, 2, 3}
	permuter(nums, 0, len(nums)-1)

	fmt.Println(bigArray)

}

func permute(nums []int) [][]int {
	permuter(nums, 0, len(nums)-1)
	return bigArray
}

func permuter(nums []int, left, right int) {
	if left >= right {
		tmp := make([]int, len(nums))
		copy(tmp, nums)
		bigArray = append(bigArray, tmp)
		return
	}
	for l := left; l <= right; l++ {
		swap(nums, left, l)
		permuter(nums, left+1, right)
		swap(nums, left, l)
	}
}
func swap(nums []int, l, r int) {
	temp := nums[l]
	nums[l] = nums[r]
	nums[r] = temp
}
