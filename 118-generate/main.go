package main

func generate(numRows int) [][]int {
	if numRows == 0 {
		return [][]int{}
	}
	if numRows == 1 {
		return [][]int{{1}}
	}
	res := [][]int{{1}}
	for i := 1; i < numRows; i++ {
		a := make([]int, i+1)
		a[0], a[i] = 1, 1
		for j := 1; j < i; j++ {
			a[j] = res[i-1][j-1] + res[i-1][j]
		}
		res = append(res, a)
	}
	return res
}

func main() {

}
