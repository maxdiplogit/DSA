package main

import (
	"fmt"
)

// Counting Sort is a non-comparison based algorithm, which means that we dont compare elements with each other in the given input array
// In Counting Sort, we need to have the input array of given size n
// Also we need to know the range in which the numbers can be for the array, and the max of that range should be k
// If the max range is not given, we can calculate the max
// Also if the range is going into negative numbers, we also calculate the min as well
func main() {
	arr := []int{-1, 5, -2, 6, 1, 3, 0, 1, 2, -1, 3, 4}

	min, max := arr[0], arr[0]
	for i := 1; i < len(arr); i++ {
		if arr[i] > max {
			max = arr[i]
		}

		if arr[i] < min {
			min = arr[i]
		}
	}

	count_arr := make([]int, max-min+1)
	for _, v := range arr {
		count_arr[v-min]++
	}

	for i, v := range count_arr {
		for j := 0; j < v; j++ {
			fmt.Printf("%d ", i+min)
		}
	}

	fmt.Println()
}
