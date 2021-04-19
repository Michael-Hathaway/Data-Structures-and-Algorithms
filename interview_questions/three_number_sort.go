/*
Prompt:
You are given an array of integers and another array of three distinct
integers, The first array is guaranteed to only contain values that
are in the second array. The second array represents a desired order
for all the elements in the first array. For example if the second
array is [x, y, z] then the desired order for the second array is
[x, x, ..., x, y, y, ..., y, z, z, ..., z]

Write a function that sorts the first array based off the order given
in the second array. The sort should occur in O(1) space.
*/

package main

import "fmt"

func ThreeNumberSort(array []int, order []int) []int {
	// get counts of each number in the array
	count := map[int]int{}
	for _, value := range array {
		if _, ok := count[value]; ok {
			count[value] += 1
		} else {
			count[value] = 1
		}
	}

	// use the accounts to populate the values in array
	posInArray := 0
	for _, value := range order {
		countOfValue, _ := count[value]
		stopPos := posInArray + countOfValue

		for posInArray < stopPos {
			array[posInArray] = value
			posInArray += 1
		}
	}

	return array
}

func main() {
	array := []int{1, 0, 0, -1, -1, 0, 1, 1}
	order := []int{0, 1, -1}

	array = ThreeNumberSort(array, order)
	fmt.Println(array)
}
