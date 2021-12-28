package main

func countQuadruplets(nums []int) (times int) {
	l := len(nums)
	times = 0
	mp := make(map[int]int)
	for i := 2; i < l-1; i++ {
		for j := 0; j < i-1; j++ {
			mp[nums[i-1]+nums[j]]++
		}
		for j := i + 1; j < l; j++ {
			times += mp[nums[j]-nums[i]]
		}
	}
	return
}

func main() {

}
