package main

import "fmt"

//https://leetcode.com/problems/majority-element-ii/
func main() {
	fmt.Println(majorityElement([]int{11, 33, 33, 11, 33, 11, 22}))
}
func majorityElement(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}
	var a, b, aCount, bCount int
	a = arr[0]
	b = arr[1]
	for i := 0; i < len(arr); i++ {
		if arr[i] == a {
			aCount++
		} else if arr[i] == b {
			bCount++
		} else if aCount == 0 {
			a = arr[i]
			aCount = 1
		} else if bCount == 0 {
			b = arr[i]
			bCount = 1
		} else {
			bCount--
			aCount--
		}
	}
	aCount = 0
	bCount = 0
	for i := 0; i < len(arr); i++ {
		if arr[i] == a {
			aCount++
		} else if arr[i] == b {
			bCount++
		}
	}
	result := []int{}
	if aCount > len(arr)/3 {
		result = append(result, a)
	}
	if bCount > len(arr)/3 {
		result = append(result, b)
	}
	return result
}
