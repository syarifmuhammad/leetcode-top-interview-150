package main

import "fmt"

func main() {
	nums1 := []int{-1, 0, 0, 3, 3, 3, 0, 0, 0}
	m := 6
	nums2 := []int{1, 2, 3}
	n := 3
	tmp := make([]int, m+n)
	left := 0
	right := 0
	for i := 0; i < m+n; i++ {
		if right >= n {
			tmp[i] = nums1[left]
			left++
		} else {
			if left < m && nums1[left] < nums2[right] {
				tmp[i] = nums1[left]
				left++
			} else {
				tmp[i] = nums2[right]
				right++
			}
		}
	}
	for i, val := range tmp {
		nums1[i] = val
	}

	fmt.Println(nums1)
}
