package main

func main() {

}
func firstMissingPositive(nums []int) int {
	for i := 0; i < len(nums); i++ {
		Bounce(nums, nums[i])
	}
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			return i
		}
	}
	return -1
}

func Bounce(nums []int, n int) {
	temp := nums[n]
	if temp == n {
		nums[n] = 0
		return
	}
	if temp >= 0 && temp < len(nums) {
		nums[n] = 0
		Bounce(nums, temp)
	}
}
