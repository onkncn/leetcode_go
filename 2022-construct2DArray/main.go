package main

func construct2DArray(original []int, m int, n int) [][]int {
	l := len(original)
	if l != m*n {
		return [][]int{}
	}
	nums := make([][]int, m)
	for i := range nums {
		nums[i] = make([]int, n)
		for j := range nums[i] {
			nums[i][j] = original[j+i*n]
		}
	}
	return nums
}

func main() {

}
