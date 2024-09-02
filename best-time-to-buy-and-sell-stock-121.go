//go:build ignore
// +build ignore

package main

import "fmt"

func main() {
	var prices []int = []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}

	if len(prices) <= 1 {
		fmt.Println(0)
	}

	result := 0
	min := prices[0]

	for _, val := range prices {
		if val-min > result {
			result = val - min
		}
		if val < min {
			min = val
		}
	}

	fmt.Println(result)

}
