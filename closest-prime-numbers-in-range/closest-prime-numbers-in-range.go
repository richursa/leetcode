// https://leetcode.com/problems/closest-prime-numbers-in-range/
package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(closestPrimes(1, 1000000))
}

func closestPrimes(left int, right int) []int {
	if left == 1 {
		left = int(math.Min(float64(left+1), float64(right)))
	}
	leftPrime := nextPrime(left, right)
	if leftPrime == right || leftPrime == -1 {
		return []int{-1, -1}
	}
	rightPrime := nextPrime(leftPrime+1, right)
	if rightPrime == -1 {
		return []int{-1, -1}
	}
	diff := rightPrime - leftPrime
	if diff == 2 {
		return []int{leftPrime, rightPrime}
	}
	next := closestPrimes(rightPrime, right)
	if next[0] == -1 {
		return []int{leftPrime, rightPrime}
	} else {
		if (next[1] - next[0]) < diff {
			return next
		}
	}
	return []int{leftPrime, rightPrime}
}
func nextPrime(current int, right int) int {
	if isPrime(current) {
		return current
	} else {
		if current%2 == 0 {
			current++
		}
	}
	for i := current; i <= right; i = i + 2 {
		if isPrime(i) {
			return i
		}
	}
	return -1
}
func isPrime(n int) bool {
	if n == 2 {
		return true
	} else {
		if n%2 == 0 {
			return false
		}
	}
	for i := 3; i <= int(math.Sqrt(float64(n))); i = i + 2 {
		if n%i == 0 {
			return false
		}
	}
	return true
}
