//go:build ignore
// +build ignore

package main

import "fmt"

func main() {
	var nums []int = []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	k := 3

	if k == len(nums) || len(nums) == 1 || k == 0 {
		return
	}
	var result []int

	count := 0
	i := len(nums) - k%len(nums)
	for count < len(nums) {
		result = append(result, nums[i])
		count++
		if i > len(nums)-2 {
			i = 0
		} else {
			i++
		}
	}

	for idx, val := range result {
		nums[idx] = val
	}

	fmt.Println(nums)

}
