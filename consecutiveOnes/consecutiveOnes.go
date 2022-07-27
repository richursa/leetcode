package main

func main() {

}

func findMaxConsecutiveOnes(nums []int) int {
	maxOnes := 0
	currentOnes := 0
	for _, num := range nums {
		if num == 1 {
			currentOnes++
			if currentOnes > maxOnes {
				maxOnes = currentOnes
			}
		} else {
			currentOnes = 0
		}
	}
	return maxOnes
}
