// Prompt:
// Write a function that takes a non-empty list of integers
// and returns a list of the same length where each element of
// the output array is equal to the product of every other number
// in the input array

// Example:
//
// IN: [5, 1, 4, 2]
//
// OUT: [8, 40, 10, 20]
//

package main

import "fmt"

// O(n) time | O(n) space
func ArrayOfProducts(array []int) []int {
	totalProduct := 1
	numberOfZeros := 0

	// calculate the total product of all the values in the array
	// if a zero is encountered, increment the count of zeros and continue
	for i := 0; i < len(array); i++ {
		if array[i] == 0 {
			numberOfZeros += 1
			continue
		}
		totalProduct *= array[i]
	}

	// create array to hold the results
	resultArray := make([]int, len(array))

	// if more than 1 zero appears in input list, all output values
	// must be zero
	if numberOfZeros > 1 {
		return resultArray
	}

	// if one zero in input, all positions in output will be zero except
	// for the one that was already zero in the input array
	if numberOfZeros == 1 {
		for i := 0; i < len(array); i++ {
			if array[i] == 0 {
				resultArray[i] = totalProduct
			}
		}
		return resultArray
	}

	// otherwise compute product for each position by dividing the
	// total product by the value at that position
	for i := 0; i < len(array); i++ {
		resultArray[i] = totalProduct / array[i]
	}
	return resultArray
}

func main() {
	inputArray := []int{5, 1, 4, 2}
	outputArray := ArrayOfProducts(inputArray)

	fmt.Printf("Expected Output: [8 40 10 20], Actual Output: %v\n", outputArray)
}
