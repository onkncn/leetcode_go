package main

import "fmt"

func originalDigits(s string) string {
	nums := [10]int{}
	al := [26]int{}
	res := ""
	for _, n := range s {
		al[n-'a']++
	}
	for al['z'-'a'] > 0 {
		for _, c := range "zero" {
			al[c-'a']--
		}
		nums[0]++
	}
	for al['x'-'a'] > 0 {
		for _, c := range "six" {
			al[c-'a']--
		}
		nums[6]++
	}
	for al['s'-'a'] > 0 {
		for _, c := range "seven" {
			al[c-'a']--
		}
		nums[7]++
	}
	for al['g'-'a'] > 0 {
		for _, c := range "eight" {
			al[c-'a']--
		}
		nums[8]++
	}
	for al['h'-'a'] > 0 {
		for _, c := range "three" {
			al[c-'a']--
		}
		nums[3]++
	}
	for al['r'-'a'] > 0 {
		for _, c := range "four" {
			al[c-'a']--
		}
		nums[4]++
	}

	for al['f'-'a'] > 0 {
		for _, c := range "five" {
			al[c-'a']--
		}
		nums[5]++
	}
	for al['i'-'a'] > 0 {
		for _, c := range "nine" {
			al[c-'a']--
		}
		nums[9]++
	}
	for al['n'-'a'] > 0 {
		for _, c := range "one" {
			al[c-'a']--
		}
		nums[1]++
	}
	for al['w'-'a'] > 0 {
		for _, c := range "two" {
			al[c-'a']--
		}
		nums[2]++
	}
	for i, n := range nums {
		for j := 0; j < n; j++ {
			res += (string)('0' + i)
		}
	}
	return res
}

func main() {
	fmt.Println(originalDigits("ereht"))
}

/*
-zero--
one----
two----
three--
four---
five---
six----
seven--
eight--
nine---


zero  z
six   x
seven s
eight g
three h
four  r
five  f
nine  i
one   n
two   w
*/
