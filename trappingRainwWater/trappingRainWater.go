package main

import (
	"fmt"
)

func main() {
	fmt.Println(trap([]int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}))
}
func trap(height []int) int {
	if len(height) <= 1 {
		return 0
	}
	leftTall := height[0]
	rightTall := 0
	rightIndex := 0
	for i, h := range height[1:] {
		if h >= leftTall {
			rightTall = h
			rightIndex = i
			break
		}
		if h > rightTall {
			rightTall = h
			rightIndex = i
		}
	}
	normaLizedLength := min(leftTall, rightTall)
	area := normaLizedLength * (rightIndex + 2)
	areaOfBars := 0
	for i := 0; i <= rightIndex+1; i++ {
		areaOfBars = areaOfBars + min(height[i], normaLizedLength)
	}
	maxRainWater := area - areaOfBars
	return maxRainWater + trap(height[rightIndex+1:])
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
