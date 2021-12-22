package main

func getRow(rowIndex int) []int {
	a := make([]int, rowIndex+1)
	a[0] = 1
	for i := 0; i < rowIndex; i++ {
		a[i+1] = 1
		for j := i; j >= 1; j-- {
			a[j] = a[j] + a[j-1]
		}
	}
	return a
}

func main() {
	getRow(5)
}
