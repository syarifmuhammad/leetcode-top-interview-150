//go:build ignore
// +build ignore

package main

import "fmt"

func main() {
	var nums []int = []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}

	count := make(map[int]int)

	for _, val := range nums {
		count[val]++
	}

	result := nums[0]

	for key, val := range count {
		if val > count[result] {
			result = key
		}
	}

	fmt.Println(result)

}
