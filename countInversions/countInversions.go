package main

import "fmt"

// https://www.codingninjas.com/codestudio/problems/615?topList=striver-sde-sheet-problems&utm_source=striver&utm_medium=website
func main() {
	fmt.Println(mergesort([]int{3, 2, 1}))
	fmt.Println(inversions)
}

var inversions int

func mergesort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}
	mid := len(arr) / 2
	left := mergesort(arr[0:mid])
	right := mergesort(arr[mid:])
	sortedArr := merge(left, right)
	return sortedArr
}
func merge(a, b []int) []int {
	arr := []int{}
	i := 0
	j := 0
	for i < len(a) && j < len(b) {
		if a[i] <= b[j] {
			arr = append(arr, a[i])
			i++
		} else {
			arr = append(arr, b[j])
			inversions += len(a[i:])
			j++
		}
	}
	if i != (len(a)) {
		for i < len(a) {
			arr = append(arr, a[i])
			i++
		}
	}
	if j != (len(b)) {
		for j < len(b) {
			arr = append(arr, b[j])
			j++
		}
	}
	return arr
}
