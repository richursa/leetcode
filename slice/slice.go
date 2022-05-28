package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 4}

	appender(nums)
	fmt.Println(nums)
}

func appender(a []int) {
	a = append(a, 5)
}
