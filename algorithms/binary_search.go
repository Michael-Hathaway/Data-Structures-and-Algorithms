// Binary Search algorithm implemented in golang

package main

import (
	"fmt"
	"sort"
)

func BinarySearch(array []int, target int) bool {
	start, end := 0, len(array)-1

	for end >= start {
		mid := start + (end-1)/2
		if array[mid] == target {
			return true
		}

		if target > array[mid] {
			start = mid + 1
		} else {
			end = mid - 1
		}
	}

	return false
}

func main() {
	array := []int{9, 2, 4, 6, 5, 1, 3, 10, 7, 8}
	sort.Ints(array)

	fmt.Println(BinarySearch(array, 3))
	fmt.Println(BinarySearch(array, 12))
	fmt.Println(BinarySearch(array, 10))
	fmt.Println(BinarySearch(array, -1))
}
