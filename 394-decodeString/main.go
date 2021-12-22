package main

func decodeString(s string) string {
	stack := []int{}
	index := 0
	res := ""
	times := 0
	flag := false
	for i := 0; i < len(s); i++ {
		if s[i] >= '0' && s[i] <= '9' && !flag {
			times = times*10 + int(s[i]-'0')
		} else if s[i] >= 'a' && s[i] <= 'z' && !flag {
			res += string(s[i])
		} else if s[i] == '[' {
			flag = true
			stack = append(stack, i)
			index++
		} else if s[i] == ']' {
			index--
			if index == 0 {
				for j := 0; j < times; j++ {
					res += decodeString(s[stack[index]+1 : i])
				}
				times = 0
				flag = false
			}
			stack = stack[:index]
		}
	}
	return res
}

func main() {

}
