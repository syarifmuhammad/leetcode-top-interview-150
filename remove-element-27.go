package main

import "fmt"

func main() {
	var nums []int = []int{}
	var val int
	result := 0
	place := 0
	exist := false
	for i, num := range nums {
		if num != val {
			result++
			nums[place] = num
			place++
		} else {
			exist = true
			if !exist {
				place = i
			}
		}
	}

	fmt.Println(result, nums)
}
