//go:build ignore

package main

import "fmt"

func main() {
	var nums []int = []int{}
	var val int
	result := 0
	place := 0
	for _, num := range nums {
		if num != val {
			result++
			nums[place] = num
			place++
		}
	}

	fmt.Println(result, nums)
}
