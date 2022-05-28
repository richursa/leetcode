package main

import "fmt"

func main() {
	nums := []int{2, 3, 4, 5, 2, 3, 5}
	fmt.Println(singleNumber(nums))
}
func singleNumber(nums []int) int {
	xor := 0
	for _, val := range nums {
		xor ^= val
	}
	return xor
}
