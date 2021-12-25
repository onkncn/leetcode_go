package main

func singleNumber(nums []int) int {
	res := 0
	for _, a := range nums {
		res ^= a
	}
	return res
}

func main() {

}
