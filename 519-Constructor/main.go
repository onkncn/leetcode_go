package main

import (
	"fmt"
	"math/rand"
)

type Solution struct {
	m, n int
	len  int
	kv   map[int]int
}

func Constructor(m int, n int) Solution {
	len := n * m
	if len > 111111 {
		len = 111111
	}
	kv := make(map[int]int)
	return Solution{m, n, len, kv}
}

func (this *Solution) Flip() []int {
	random := rand.Intn(this.len)
	res, e := this.kv[random]
	if !e {
		res = random
	}
	last, e := this.kv[this.len-1]
	if !e {
		last = this.len - 1
	}
	this.kv[random] = last
	this.len--
	return []int{res / this.n, res % this.n}
}

func (this *Solution) Reset() {
	this.kv = make(map[int]int)
	this.len = this.m * this.n
}

/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(m, n);
 * param_1 := obj.Flip();
 * obj.Reset();
 */

func main() {
	s := Constructor(3, 1)
	fmt.Println(s.Flip())
	fmt.Println(s.Flip())
	fmt.Println(s.Flip())
	s.Reset()
	fmt.Println(s.Flip())
}
