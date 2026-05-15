package main

import (
	"fmt"
)

// In Merge Sort, we keep on dividing until we reach a single element which is already sorted (subproblem), and to do this we keep on halving the array
// Halving means log(n) times we have to halve to get individual elements for array of size n
// Now when we start merging each and every pair for a particular merge level, their will always be n comparisons in total across all the levels
// So at each level we are doing n comparisons each, hence the time complexity of O(n.logn)
func main() {
	arr := [...]int{4, 0, 9, 8, 1, 2, 6, 3, 7, 5}
	arrSlice := arr[:]

	fmt.Printf("Array: %#v\n", arr)

	sorted_arrSlice := merge_sort(arrSlice)

	fmt.Printf("Sorted Array: %#v\n", sorted_arrSlice)
}

func merge_sort(arrSlice []int) []int {
	if len(arrSlice) == 1 {
		return arrSlice
	}

	half := len(arrSlice) / 2

	leftSortedSlice := merge_sort(arrSlice[:half])
	rigtSortedSlice := merge_sort(arrSlice[half:])

	return merge(leftSortedSlice, rigtSortedSlice)
}

func merge(leftHalf, rightHalf []int) []int {
	res := make([]int, len(leftHalf)+len(rightHalf))

	k := 0

	i := 0
	j := 0

	for i < len(leftHalf) && j < len(rightHalf) {
		if leftHalf[i] <= rightHalf[j] {
			res[k] = leftHalf[i]
			i++
		} else {
			res[k] = rightHalf[j]
			j++
		}
		k++
	}

	if i < len(leftHalf) {
		for i < len(leftHalf) {
			res[k] = leftHalf[i]
			i++
			k++
		}
	}

	if j < len(rightHalf) {
		for j < len(rightHalf) {
			res[k] = rightHalf[j]
			j++
			k++
		}
	}

	return res
}
