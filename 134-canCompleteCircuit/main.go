package main

func canCompleteCircuit(gas []int, cost []int) int {
	l := len(gas)
	for i := 0; i < l; i++ {
		has := 0
		j := 0
		for ; j < l; j++ {
			has += gas[(i+j)%l] - cost[(i+j)%l]
			if has < 0 {
				i = i + j
				break
			}
		}
		if has >= 0 {
			return i
		}
	}
	return -1
}

func main() {

}
