package main

import "fmt"

func main() {
	fmt.Println(firstUniqChar("leetcode"))
}
func firstUniqChar(s string) int {
	charMap := make(map[rune]int)
	for _, char := range s {
		charMap[char]++
	}
	for i, char := range s {
		if charMap[char] == 1 {
			return i
		}
	}
	return -1
}
