package main

func singleNumber(nums []int) int {
	a, b := 0, 0
	for _, num := range nums {
		a, b = b&^a&num|a&^b&^num, (b^num)&^a
	}
	return b
}

func main() {

}
