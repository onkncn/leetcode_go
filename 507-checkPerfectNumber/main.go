package main

func checkPerfectNumber(num int) bool {
	sum := 0
	if num < 6 {
		return false
	}
	i := 0
	for i = 2; i*i < num; i++ {
		if num%i == 0 {
			sum += i + num/i
		}
	}
	sum += 1
	if i*i == num {
		sum += i
	}
	return sum == num
}

func main() {

}
