package main

func merge(nums1 []int, m int, nums2 []int, n int) {
	p1, p2, i := 0, 0, 0
	t := make([]int, m+n)
	for p1 < m && p2 < n {
		if nums1[p1] <= nums2[p2] {
			t[i] = nums1[p1]
			p1++
		} else {
			t[i] = nums2[p2]
			p2++
		}
		i++
	}
	for p1 < m {
		t[i] = nums1[p1]
		i++
		p1++
	}
	for p2 < n {
		t[i] = nums2[p2]
		i++
		p2++
	}
	for i := 0; i < m+n; i++ {
		nums1[i] = t[i]
	}
}

func main() {

}
