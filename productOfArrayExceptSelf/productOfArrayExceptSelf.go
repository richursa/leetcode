package main

func main() {

}
func productExceptSelf(nums []int) []int {
	var product, numZero int
	product = 1
	var array []int
	for _, number := range nums {
		if number == 0 {
			numZero++
			continue
		}
		product = product * number
	}
	if numZero == 0 {
		for _, number := range nums {
			array = append(array, product/number)
		}
	} else if numZero == 1 {
		for _, number := range nums {
			if number == 0 {
				array = append(array, product)
				continue
			}
			array = append(array, 0)
		}
	} else {
		for _, _ = range nums {
			array = append(array, 0)
		}
	}
	return array
}
