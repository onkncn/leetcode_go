package main

func twoSum(nums []int, target int) []int {
	mp := make(map[int]int)
	nums = append([]int{0}, nums...)
	for i := 1; i < len(nums); i++ {
		if mp[target-nums[i]] != 0 {
			return []int{mp[target-nums[i]] - 1, i - 1}
		} else {
			mp[nums[i]] = i
		}
	}
	return []int{}
}

func main() {

}
