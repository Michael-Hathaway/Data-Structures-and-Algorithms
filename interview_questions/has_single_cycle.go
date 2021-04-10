// PROMPT
// You're given an array of intergers where each integer represents a jump
// to a new position in the array. For example if the value at index 0 was 2
// that would represent a jump forward two places to index 2.
// Given this array write a function the returns a boolean denoting whether
// or not the jumps in the array form a single cycle. A single cycle occurs
// following the jumps leads from the start pos to all other positions
// exactly once and then back to the start position

package main

import "fmt"

func AreAllMapValuesTrue(visited map[int]bool) bool {
	for _, hasVisited := range visited {
		if !hasVisited {
			return false
		}
	}
	return true
}

func HasSingleCycle(array []int) bool {
	visited := map[int]bool{}
	for index := range array {
		visited[index] = false
	}

	currentIndex := 0
	for {
		// if at the start index and have already visited start
		// check that all positions have been visited and return result
		haveVisitedStart, ok := visited[0]
		if currentIndex == 0 && ok && haveVisitedStart {
			complete := AreAllMapValuesTrue(visited)
			return complete
		}

		// if we've already visited a pos but it isn't start, return false
		alreadyVisited, ok := visited[currentIndex]
		if ok && alreadyVisited {
			return false
		}

		// update visited for current position
		visited[currentIndex] = true

		// get next index to jump to
		nextIndex := (currentIndex + array[currentIndex]) % len(array)
		if nextIndex < 0 {
			nextIndex = len(array) + (nextIndex % len(array))
		}

		currentIndex = nextIndex
	}
}

func main() {
	array := []int{2, 3, 1, -4, -4, 2}
	hasSingleCycle := HasSingleCycle(array)

	fmt.Println("Has a single cycle?: ", hasSingleCycle)
}
