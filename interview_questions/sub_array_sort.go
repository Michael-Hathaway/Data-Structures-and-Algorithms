/*
PROMPT
Write a function that takes in an array of at least two integers
and returns an array of the starting and ending indices of the smallest
subarray in the input array that needs to be sorted in place in order
for the entire input array to be sorted in ascending order.

If the array is already sorted, return [-1, -1]
*/
package main

import (
	"fmt"
	"math"
)

func SubarraySort(array []int) []int {
	minNum := math.MaxInt32
	// find the smallest number that isn't in the right place
	for i := 0; i < len(array)-1; i++ {
		if array[i] > array[i+1] {
			minNum = min(minNum, array[i+1])
		}
	}

	// if the array is already sorted return [-1, -1]
	if minNum == math.MaxInt32 {
		return []int{-1, -1}
	}

	// Find the start index of the unsorted array, by finding the
	// first number less than or equal to the smallest number in the
	// array
	start := 0
	for i := range array {
		if array[i] >= minNum {
			start = i
			if array[i] == minNum {
				start += 1
			}
			break
		}
	}

	// Find the end index of the unsorted array
	var end int
	itr := start
	maxInRange := math.MinInt32
	for ; itr < len(array); itr++ {
		if array[itr] <= maxInRange {
			end = itr
		} else {
			maxInRange = array[itr]
		}

		if array[itr] <= minNum {
			end = itr
		}
	}

	return []int{start, end}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	array := []int{1, 2, 4, 7, 10, 11, 7, 12, 6, 7, 16, 18, 19}
	result := SubarraySort(array)
	expected := []int{3, 9}

	fmt.Printf("Expected: %v, Recieved: %v\n", expected, result)
}
