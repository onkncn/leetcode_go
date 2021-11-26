package main

import (
	"fmt"
	"math"
)

func poorPigs(buckets, minutesToDie, minutesToTest int) int {
	return int(math.Ceil(math.Log(float64(buckets)) / math.Log(float64(minutesToTest/minutesToDie+1))))
}
func main() {
	fmt.Println(poorPigs(4, 15, 15))
}
