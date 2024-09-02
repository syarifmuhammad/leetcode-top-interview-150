//go:build ignore
// +build ignore

package main

import "fmt"

func main() {
	var nums []int = []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	pivot := nums[0]
	place := 1
	result := 1
	for i := 1; i < len(nums); i++ {
		if nums[i] != pivot {
			result++
			pivot = nums[i]
			nums[place] = nums[i]
			place++
		}
	}

	fmt.Println(result, nums)

}
