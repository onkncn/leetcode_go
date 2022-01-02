package main

func hammingWeight(num uint32) int {
	times := 0
	for num != 0 {
		if num%2 == 1 {
			times++
		}
		num /= 2
	}
	return times
}

func main() {

}
