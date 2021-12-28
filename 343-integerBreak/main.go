package main

func integerBreak(n int) int {
	if n < 3 {
		return 1
	}
	if n == 3 {
		return 2
	}
	times := n / 3
	sum := 1
	for i := 0; i < times; i++ {
		sum *= 3
	}
	if n%3 == 2 {
		sum *= 2
	}
	if n%3 == 1 {
		sum = sum / 3 * 4
	}
	return sum
}

func main() {

}
