package main

func plusOne(digits []int) []int {
	ld := len(digits)

	digits[ld-1] += 1
	for i := ld - 1; i > 0; i-- {
		if digits[i]/10 > 0 {
			digits[i] -= 10
			digits[i-1]++
		} else {
			break
		}
	}
	if digits[0]/10 > 0 {
		digits[0] -= 10
		digits = append([]int{1}, digits...)
	}
	return digits
}

func main() {

}
