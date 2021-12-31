package main

func isRectangleCover(rectangles [][]int) bool {
	mp := make(map[[2]int]int)
	s := 0
	min, max := [2]int{rectangles[0][0], rectangles[0][1]}, [2]int{rectangles[0][2], rectangles[0][3]}
	for i := range rectangles {
		if min[0] > rectangles[i][0] || min[1] > rectangles[i][1] {
			min = [2]int{rectangles[i][0], rectangles[i][1]}
		}
		if max[0] < rectangles[i][2] || max[1] < rectangles[i][3] {
			max = [2]int{rectangles[i][2], rectangles[i][3]}
		}
		mp[[2]int{rectangles[i][0], rectangles[i][1]}]++
		mp[[2]int{rectangles[i][2], rectangles[i][3]}]++
		mp[[2]int{rectangles[i][0], rectangles[i][3]}]++
		mp[[2]int{rectangles[i][2], rectangles[i][1]}]++
		s += (rectangles[i][0] - rectangles[i][2]) * (rectangles[i][1] - rectangles[i][3])
		//fmt.Println((rectangles[i][0]-rectangles[i][2])*(rectangles[i][1]-rectangles[i][3]))
	}
	times := 0
	for k, v := range mp {
		if v == 1 {
			if k != min && k != max && k != [2]int{min[0], max[1]} && k != [2]int{max[0], min[1]} {
				return false
			}
			times++
		}
		if v == 3 {
			return false
		}
		//fmt.Println(k,v)
	}
	//fmt.Println(s,min,max)
	return times == 4 && s == (min[0]-max[0])*(min[1]-max[1])
}

func main() {

}
