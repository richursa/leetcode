package main

func main() {

}
func hammingWeight(num uint32) int {
	startBit := uint32(1)
	count := 0
	for i := 0; i < 32; i++ {
		if startBit&num != 0 {
			count++
		}
		startBit = startBit << 1
	}
	return count
}
