// PROMPT
// Write a funcion that takes as arguments an array of integers
// and a target sum. The function should find all unique triplets
// in the array that sum to the target value. The function should
// return an array of arrays where each inner array is a triplet
// of numbers that sum to the target value

package main

import (
	"fmt"
	"sort"
)

func ThreeNumberSum(array []int, target int) [][]int {
	// sort array before starting
	sort.Ints(array)

	results := [][]int{}
	for i := 0; i < len(array)-2; i++ {
		// look at numbers greater than and less than the current
		// number and try to construct the target sum
		left := i + 1
		right := len(array) - 1

		for left < right {
			currentSum := array[i] + array[left] + array[right]
			if currentSum == target {
				// found target
				newTriplet := []int{array[i], array[left], array[right]}
				results = append(results, newTriplet)
				// adjust left and right to look for new target
				left, right = left+1, right-1
			} else if currentSum < target {
				// we need to increase our sum, so move left forward
				left += 1
			} else {
				// we need to decrease our sum, so move right back
				right -= 1
			}
		}
	}

	return results
}

func main() {
	array := []int{12, 3, 1, 2, -6, 5, -8, 6}
	target := 0
	threeNumberSums := ThreeNumberSum(array, target)

	fmt.Println("Three Number Sums: ", threeNumberSums)
}
