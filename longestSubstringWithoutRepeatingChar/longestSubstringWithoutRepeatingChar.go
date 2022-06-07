// https://leetcode.com/problems/longest-substring-without-repeating-characters/
package main

import "fmt"

func main() {
	fmt.Println(lengthOfLongestSubstring("abcabcbb"))
}

func lengthOfLongestSubstring(s string) int {
	prefixMap := make(map[byte]int)
	maxLength := 0
	prevLength := 0
	for i := 0; i < len(s); i++ {
		if index, ok := prefixMap[s[i]]; !ok {
			prevLength = prevLength + 1
		} else {
			if i-index <= prevLength {
				prevLength = i - index
			} else {
				prevLength = prevLength + 1
			}
		}
		prefixMap[s[i]] = i
		if prevLength > maxLength {
			maxLength = prevLength
		}
	}
	return maxLength
}
