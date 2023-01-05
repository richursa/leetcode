package main

import "math"

func main() {

}
func moveZeroes(nums []int) {
	integerIndex := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			integerIndex = int(math.Max(float64(i), float64(integerIndex)))
			for k := integerIndex; k < len(nums); k, integerIndex = k+1, integerIndex+1 {
				if nums[k] != 0 {
					nums[i], nums[k] = nums[k], nums[i]
					break
				}
			}
		}
	}
}
