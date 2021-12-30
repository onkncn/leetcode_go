package main

func maxProfit(prices []int) int {
	b1, s1 := -prices[0], 0
	b2, s2 := -prices[0], 0
	for i := 1; i < len(prices); i++ {
		b1 = mymax(-prices[i], b1)
		s1 = mymax(prices[i]+b1, s1)
		b2 = mymax(-prices[i]+s1, b2)
		s2 = mymax(prices[i]+b2, s2)
	}
	return s2
}
func mymax(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {

}
