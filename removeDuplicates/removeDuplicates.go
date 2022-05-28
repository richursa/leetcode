package main

import "fmt"

func main() {
	var array = []int{1, 2, 3, 4, 5, 5, 5, 6, 6, 7}
	fmt.Println(removeDuplicates(array))
	fmt.Println(array)
}
func removeDuplicates(nums []int) int {
	lastUniqueNumber := 0
	for i := 0; i < len(nums)-1; i++ {
		if nums[lastUniqueNumber] != nums[i+1] {
			lastUniqueNumber++
			nums[lastUniqueNumber] = nums[i+1]
		}
	}
	return lastUniqueNumber + 1
}
