package main

func setZeroes(matrix [][]int) {
	n, m := len(matrix), len(matrix[0])
	flagn, flagm := make([]bool, m), make([]bool, n)
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if matrix[i][j] == 0 {
				flagn[j] = true
				flagm[i] = true
			}
		}
	}
	ch := make(chan int)
	go func() {
		for i := 0; i < m; i++ {
			if flagn[i] == true {
				for j := 0; j < n; j++ {
					matrix[j][i] = 0
				}
			}
		}
	}()
	for i := 0; i < n; i++ {
		if flagm[i] == true {
			for j := 0; j < m; j++ {
				matrix[i][j] = 0
			}
		}
	}
	<-ch
}
func main() {

}
