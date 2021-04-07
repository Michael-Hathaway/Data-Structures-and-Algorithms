// Implementation of the insertion sort algorithm in Go
//
//

package main

import "fmt"

func InsertionSort(array []int) {
	for i := 1; i < len(array); i++ {
		j := i - 1
		currentValue := array[i]

		for j >= 0 && array[j] > currentValue {
			array[j+1] = array[j]
			j -= 1
		}

		array[j+1] = currentValue
	}
}

func main() {
	array := []int{5, 1, 3, 75, 32, 43, 23}
	fmt.Printf("Array before sorting: %v\n", array)

	InsertionSort(array)

	fmt.Printf("Array after sorting: %v\n", array)
}
