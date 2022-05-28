package main

import "fmt"

func main() {
	array := []int{1, 2, 3, 4, 5}
	fmt.Println(maxProfit(array))

}
func maxProfit(prices []int) int {

	boughtPrice := 10001
	profit := 0
	for i := 0; i < len(prices)-1; i++ {
		if boughtPrice > prices[i] {
			boughtPrice = prices[i]
		}
		if prices[i] > prices[i+1] {
			profit = profit + prices[i] - boughtPrice
			boughtPrice = prices[i+1]
		}
		if i == (len(prices) - 2) {
			profit = profit + prices[i+1] - boughtPrice
		}
	}
	return profit
}
