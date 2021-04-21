/*
Prompt
You're given a non empty array of integers where each integer represents
the maximum number of steps you can take forward in the array. Write a
function that returns the minimum number of steps needed to reach the
final index
*/
package main

import (
	"fmt"
	"math"
)

func MinNumberOfJumps(array []int) int {
	return MinNumberOfJumpsHelper(array, 0)
}

func MinNumberOfJumpsHelper(array []int, pos int) int {
	if pos == len(array)-1 {
		return 0
	}

	if pos > len(array)-1 {
		return math.MaxInt32
	}

	minJumps := math.MaxInt32
	for i := 1; i <= array[pos]; i++ {
		result := MinNumberOfJumpsHelper(array, pos+i)
		minJumps = min(minJumps, result)
	}

	return 1 + minJumps
}

// helper function to get the min between two values
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	array := []int{3, 4, 2, 1, 2, 3, 7, 1, 1, 1, 3}
	result := MinNumberOfJumps(array)

	fmt.Printf("Expected: %v, Received: %v\n", 4, result)
}
