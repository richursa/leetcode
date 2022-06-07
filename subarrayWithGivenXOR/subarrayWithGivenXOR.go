// https://www.interviewbit.com/problems/subarray-with-given-xor/
package main

import "fmt"

func main() {
	fmt.Println(solve([]int{1, 2, 2, 3}, 0))
}

func solve(A []int, B int) int {
	prefixXORMap := make(map[int]int)
	maxCount := 0
	xor := 0
	prefixXORMap[0] = 1
	for i := 0; i < len(A); i++ {
		xor = xor ^ A[i]
		maxCount += prefixXORMap[xor^B]
		prefixXORMap[xor]++
	}
	return maxCount
}
