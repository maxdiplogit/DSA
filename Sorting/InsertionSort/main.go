package main

import (
	"fmt"
)

// In Insertion Sort, we treat the left part as sorted and the right part as unsorted
func main() {
	arr := [...]int{4, 0, 9, 8, 1, 2, 6, 3, 7, 5}

	fmt.Printf("Array: %#v\n", arr)

	for i := 1; i < len(arr); i++ {
		key := arr[i]

		j := i - 1
		for j >= 0 && arr[j] > key {
			arr[j+1] = arr[j]
			j--
		}

		arr[j+1] = key
	}

	fmt.Printf("Sorted Array: %#v\n", arr)
}
