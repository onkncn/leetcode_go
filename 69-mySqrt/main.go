package main

func mySqrt(x int) int {
	t := 1
	for t*t > x || (t+1)*(t+1) <= x {
		t = (x/t + t) / 2
	}
	return t
}

func main() {

}
