package main

func maxProfit(prices []int) int {
	l := len(prices)
	if l < 2 {
		return 0
	}
	min := 11111
	max := 0
	res := -11111
	for i := 0; i < l; i++ {
		if prices[i] <= min {
			min = prices[i]
			max = 0
		}
		if prices[i] >= max {
			max = prices[i]
			if max-min > res {
				res = max - min
			}
		}
	}
	return res
}
func main() {

}
