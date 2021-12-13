package main

func toLowerCase(s string) string {
	res := make([]byte, len(s))
	for i, c := range s {
		if c <= 'Z' && c >= 'A' {
			res[i] = byte(c - 'A' + 'a')
		} else {
			res[i] = byte(c)
		}
	}
	return string(res)
}

func main() {

}
