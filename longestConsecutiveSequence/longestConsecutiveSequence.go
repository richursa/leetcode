// https://leetcode.com/problems/longest-consecutive-sequence/
package main

func main() {

}
func longestConsecutive(nums []int) int {
	maxConsecutive, currentConsecutive := 0, 0
	numMap := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		numMap[nums[i]] = 0
	}
	for val, _ := range numMap {
		currentConsecutive = calculateConsecutive(numMap, val)
		if currentConsecutive > maxConsecutive {
			maxConsecutive = currentConsecutive
		}
	}
	return maxConsecutive
}

func calculateConsecutive(numMap map[int]int, current int) int {
	if numMap[current] != 0 {
		return numMap[current]
	}
	i := 1
	if consecutive, ok := numMap[current+i]; ok {
		if consecutive == 0 {
			numMap[current+i] = calculateConsecutive(numMap, current+i)
		}
		numMap[current] = numMap[current+i] + 1
	} else {
		numMap[current] = 1
	}
	return numMap[current]
}
