package main

func longestConsecutive(nums []int) int {
	mp := make(map[int]int)
	l := len(nums)
	max := 0
	for i := 0; i < l; i++ {
		t := 0
		if mp[nums[i]] == 0 {
			mp[nums[i]]++
			l, r := mp[nums[i]-1], mp[nums[i]+1]
			t = l + r + 1
			mp[nums[i]-l] = t
			mp[nums[i]+r] = t
		}
		if max < t {
			max = t
		}
	}
	return max
}

func main() {

}
