package main

import "fmt"

func maxSumOfThreeSubarrays(nums []int, k int) []int {
	ln := len(nums)
	sum := make([]int, ln)
	for i := 0; i < k; i++ {
		sum[0] += nums[i]
	}
	for i := 1; i < ln-k+1; i++ {
		sum[i] = sum[i-1] - nums[i-1] + nums[i+k-1]
	}
	fmt.Println(sum)
	maxsum1, maxsum2, maxsum3 := sum[0], sum[0]+sum[k], sum[0]+sum[k]+sum[2*k]
	maxindex1, maxindex2, maxindex3 := 0, []int{0, k}, []int{0, k, 2 * k}
	i := 1
	for i+2*k < ln {
		fmt.Println(maxindex1, maxindex2, maxindex3, maxsum1, maxsum2, maxsum3)
		if maxsum1 < sum[i] {
			maxsum1 = sum[i]
			maxindex1 = i
		}
		if maxsum2 < maxsum1+sum[k+i] {
			maxsum2 = maxsum1 + sum[k+i]
			maxindex2[0] = maxindex1
			maxindex2[1] = k + i
		}
		if maxsum3 < maxsum2+sum[2*k+i] {
			maxsum3 = maxsum2 + sum[2*k+i]
			maxindex3[0] = maxindex2[0]
			maxindex3[1] = maxindex2[1]
			maxindex3[2] = 2*k + i
		}
		i++
	}
	return maxindex3
}

// 3 3 3 8 13 12 6 7 13 15 17 120 114 7 9 11 13 0 0

func main() {
	fmt.Println(maxSumOfThreeSubarrays([]int{1, 2, 1, 2, 6, 7, 5, 1, 6, 7}, 2))
}
