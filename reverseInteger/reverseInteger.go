package main

import (
	"fmt"
)

func main() {
	fmt.Println(reverse(1534236469))
}

func reverse(xx int) int {
	x := int32(xx)
	straight := x
	reverse := int32(0)
	firstDigit := int32(0)
	for x != 0 {
		reverse = reverse*10 + x%10
		firstDigit = x % 10
		x = x / 10
	}
	if reverse == 0 {
		return 0
	}
	if firstDigit != reverse%10 {
		return 0
	}
	if straight/reverse < 0 {
		return 0
	}
	return int(reverse)
}
