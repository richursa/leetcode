package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(twoSum([]int{3, 3}, 6))
}
func twoSum(nums []int, target int) []int {
	mapper := make(map[int]int)
	mapper2 := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		if _, ok := mapper[nums[i]]; !ok {
			mapper[nums[i]] = i
		} else {
			mapper2[nums[i]] = i
		}
	}
	required := 0
	var indexes []int
	for k := range mapper {
		required = target - k
		if required == k {
			if _, ok := mapper2[required]; ok {
				indexes = []int{mapper[k], mapper2[required]}
				break
			}
		} else if _, ok := mapper[required]; ok {
			indexes = []int{mapper[k], mapper[required]}
			break
		}
	}
	sort.Ints(indexes)
	return indexes
}
