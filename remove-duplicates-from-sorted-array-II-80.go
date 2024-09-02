//go:build ignore

package main

import "fmt"

func main() {
	var nums []int = []int{1, 1, 1, 2, 2, 3}
	var result []int
	count := 0
	allcount := 0
	for i, val := range nums {
		place := i - 1
		if place < 0 {
			place = 0
		}
		if val == nums[place] {
			count++
		} else {
			count = 1
		}
		if count < 3 {
			result = append(result, val)
			allcount++
		}
	}

	for i, _ := range result {
		nums[i] = result[i]
	}

	fmt.Println(nums)
	fmt.Println(allcount)
}
