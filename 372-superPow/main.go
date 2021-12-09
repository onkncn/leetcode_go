package main

func Pow10(a int) int {
	t := a
	for i := 0; i < 9; i++ {
		t = (a * t) % 1337
	}
	return t
}

func superPow(a int, b []int) int {
	if a == 1 {
		return 1
	}
	sum := 1
	a = a % 1337
	for i := len(b) - 1; i >= 0; i-- {
		for j := 0; j < b[i]; j++ {
			sum = (sum * a) % 1337
		}
		a = Pow10(a)
	}
	return sum
}

func main() {

}
