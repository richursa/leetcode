package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(fizzBuzz(6))
}
func fizzBuzz(n int) []string {
	var answer []string
	for i := 1; i <= n; i++ {
		if i%3 == 0 && i%5 == 0 {
			answer = append(answer, "FizzBuzz")
		} else if i%5 == 0 {
			answer = append(answer, "Buzz")

		} else if i%3 == 0 {
			answer = append(answer, "Fizz")
		} else {
			answer = append(answer, strconv.FormatInt(int64(i), 10))
		}
	}
	return answer
}
