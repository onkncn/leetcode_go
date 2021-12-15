package main

import (
	"fmt"
	"math"
	"sort"
)

func visiblePoints(points [][]int, angle int, location []int) int {
	res := 0
	lp := len(points)
	if lp < 2 {
		return lp
	}
	angle_point := []float64{}
	for i := 0; i < lp; i++ {
		if points[i][0] == location[0] && points[i][1] == location[1] {
			res++
		} else {
			ans := math.Atan2(float64(points[i][0]-location[0]), float64(points[i][1]-location[1])) / math.Pi * 180
			if ans < 0 {
				ans += 360
			}
			angle_point = append(angle_point, ans)

		}
	}
	lp = lp - res
	sort.Float64s(angle_point)
	fa := float64(angle)
	p1, p2, max := 0, 1, 0
	for p1 < lp {
		t := angle_point[p2%lp] + float64(p2/lp*360)
		for angle_point[p1]+fa >= t && p1 != p2%lp {
			p2++
			t = angle_point[p2%lp] + float64(p2/lp*360)
		}
		if p2-p1 > max {
			max = p2 - p1
			if max == lp {
				return max + res
			}
		}
		p1++
		if p1 == p2 {
			p2++
		}
	}
	return max + res
}

func main() {
	fmt.Println(visiblePoints([][]int{{1, 1}, {2, 2}, {3, 3}, {4, 4}, {1, 2}, {2, 1}}, 0, []int{0, 0}))
}
