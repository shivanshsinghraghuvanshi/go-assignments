//The following code merges two slices of integers but has redundancy and does not remove duplicate elements.
// Refactor the code to eliminate repetition and ensure that the resulting slice contains unique values.

package main

import "fmt"

func mergeSlices(a, b []int) []int {
	// Merge slices a and b
	result := append(a, b...)

	// TODO: Remove duplicates from the result slice

	return result
}

func main() {
	slice1 := []int{1, 2, 3}
	slice2 := []int{3, 4, 5}
	merged := mergeSlices(slice1, slice2)
	fmt.Println(merged)
}
