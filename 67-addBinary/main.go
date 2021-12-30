package main

import "fmt"

func addBinary(a string, b string) string {
	la := len(a)
	lb := len(b)
	if lb > la {
		return addBinary(b, a)
	}
	for i := 0; i < la-lb; i++ {
		b = "0" + b
	}
	fmt.Println(b)
	res := ""
	c := 0
	for i := la - 1; i >= 0; i-- {
		t := int(a[i]-'0') + int(b[i]-'0') + c
		res = string(t%2+'0') + res
		c = t / 2
	}
	if c == 1 {
		res = "1" + res
	}
	return res
}

func main() {

}
