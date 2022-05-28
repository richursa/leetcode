package main

import "fmt"

func main() {
	fmt.Println(maxProduct([]int{2, 3, -2, 4}))
}
func maxProduct(nums []int) int {
	startPos := 0
	maxArray := []int{}
	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			if startPos != i {
				maxArray = append(maxArray, findMaxProduct(nums[startPos:i]))
			}
			maxArray = append(maxArray, 0)
			startPos = i
		}
	}
	maxArray = append(maxArray, findMaxProduct(nums[startPos:]))
	return max(maxArray)
}
func findMaxProduct(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	product := 1
	negative := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] < 0 {
			negative++
		}
	}
	if negative%2 == 0 {
		for i := 0; i < len(nums); i++ {
			product *= nums[i]
		}
		return product
	}
	maxS := []int{}
	negLeft := negative
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] < 0 {
			if negLeft == negative {
				maxS = append(maxS, findMaxProduct(nums[:i]))
				maxS = append(maxS, findMaxProduct(nums[i:]))
			} else if negLeft == 1 {
				maxS = append(maxS, findMaxProduct(nums[:i]))
				maxS = append(maxS, findMaxProduct(nums[i+1:]))
			}
			negLeft--
		}
	}
	fmt.Println(maxS)
	return max(maxS)

}
func max(arr []int) int {
	curMax := -999999999
	for i := 0; i < len(arr); i++ {
		if arr[i] > curMax {
			curMax = arr[i]
		}
	}
	return curMax
}
