// Prompt
// Write a function that takes in a non-empty array of intervals,
// merges any intervals that overlap, and returns the array of merged intervals
//
// Example:
//
// IN: [[1 2] [4 7] [3 5] [6 8] [9 10]]
//
// OUT: [[1 2] [3 8] [9 10]]
// * [3 5], [4 7], and [5 8] have been merged to form [3 8]

package main

import (
	"fmt"
	"sort"
)

type Intervals [][]int

func (is Intervals) Len() int {
	return len(is)
}

func (is Intervals) Less(i, j int) bool {
	return is[i][0] < is[j][0]
}

func (is Intervals) Swap(i, j int) {
	is[i], is[j] = is[j], is[i]
}

func SortIntervals(intervals Intervals) {
	sort.Sort(intervals)
}

// Function to merge a list of overlapping intervals
// O(n * log(n)) time | O(n) space
// sorting is O(n * log(n)), merging is O(n)
func MergeOverlappingIntervals(intervals Intervals) Intervals {
	// Sort intervals by the start value of each interval
	SortIntervals(intervals)
	newIntervals := [][]int{}

	// get the first interval and the index of the next
	currentInterval := intervals[0]
	nextIntervalIndex := 1

	for nextIntervalIndex <= len(intervals) {
		// next interval is invalid, so current interval must be the end
		if nextIntervalIndex >= len(intervals) {
			newIntervals = append(newIntervals, currentInterval)
			return newIntervals
		}

		currentIntervalEnd := currentInterval[1]
		nextIntervalStart := intervals[nextIntervalIndex][0]

		// if current and next interval overlap, we will merge
		if nextIntervalStart <= currentIntervalEnd {
			var newEnd int
			// determine end of the new interval by comparing the current and
			// next interval ends
			if nextIntervalEnd := intervals[nextIntervalIndex][1]; nextIntervalEnd > currentIntervalEnd {
				newEnd = nextIntervalEnd
			} else {
				newEnd = currentIntervalEnd
			}

			// merge intervals
			currentIntervalStart := currentInterval[0]
			currentInterval = []int{currentIntervalStart, newEnd}
			nextIntervalIndex += 1
			// dont add yet, we need to check if the next interval needs to
			// be merged into this one too
		} else {
			// no merge needed, just add current interval
			newIntervals = append(newIntervals, currentInterval)
			currentInterval = intervals[nextIntervalIndex]
			nextIntervalIndex += 1
		}
	}

	return newIntervals
}

func main() {
	intervals := Intervals{
		[]int{1, 2},
		[]int{4, 7},
		[]int{3, 5},
		[]int{6, 8},
		[]int{9, 10},
	}

	intervals = MergeOverlappingIntervals(intervals)
	fmt.Println(intervals) // [[1 2] [3 8] [9 10]]
}
