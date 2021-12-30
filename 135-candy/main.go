package main

func candy(ratings []int) int {
	l := len(ratings)
	ratings = append(ratings, ratings[0])
	minCandy := make([]int, l)
	minCandy[0] = 1
	for i := 1; i < l; i++ {
		if ratings[i] > ratings[i-1] {
			minCandy[i] = minCandy[i-1] + 1
		} else {
			minCandy[i] = 1
		}
	}
	for i := l - 2; i >= 0; i-- {
		if ratings[i] > ratings[i+1] && minCandy[i+1]+1 > minCandy[i] {
			minCandy[i] = minCandy[i+1] + 1
		}
	}
	sum := 0
	//fmt.Println(minCandy)
	for _, t := range minCandy {
		sum += t
	}
	return sum
}
func main() {

}
