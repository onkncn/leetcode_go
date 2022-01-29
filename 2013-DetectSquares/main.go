package main

import "fmt"

type DetectSquares struct {
	pointnums map[[2]int]int
	x2y       map[int][]int
}

func Constructor() DetectSquares {
	return DetectSquares{
		pointnums: map[[2]int]int{},
		x2y:       map[int][]int{},
	}
}

func (this *DetectSquares) Add(point []int) {
	this.pointnums[[2]int{point[0], point[1]}]++
	if this.pointnums[[2]int{point[0], point[1]}] == 1 {
		this.x2y[point[0]] = append(this.x2y[point[0]], point[1])
	}
}

func (this *DetectSquares) Count(point []int) int {
	res := 0
	sy := this.x2y[point[0]]
	for i := range sy {
		numsm := this.pointnums[[2]int{point[0], sy[i]}]
		fmt.Println(numsm)
		if point[1] == sy[i] {
			continue
			res += numsm * (numsm - 1) * (numsm - 2) / 6
			fmt.Println(res)
		}
		c := point[1] - sy[i]
		rx0, ry0 := point[0]+c, point[1]
		rx1, ry1 := point[0]+c, sy[i]
		numsr := this.pointnums[[2]int{rx0, ry0}] * this.pointnums[[2]int{rx1, ry1}]
		lx0, ly0 := point[0]-c, point[1]
		lx1, ly1 := point[0]-c, sy[i]
		numsl := this.pointnums[[2]int{lx0, ly0}] * this.pointnums[[2]int{lx1, ly1}]
		//fmt.Println(numsm,numsl,numsr)
		res += numsm*numsr + numsm*numsl
		//fmt.Println(res)
	}
	//fmt.Println(res)
	return res
}
func main() {
	d := Constructor()
	d.Add([]int{1, 1})
	d.Add([]int{1, 1})
	d.Add([]int{1, 1})
	d.Add([]int{1, 1})
	d.Add([]int{1, 1})
	d.Add([]int{1, 1})
	fmt.Println(d.Count([]int{1, 1}))
}
