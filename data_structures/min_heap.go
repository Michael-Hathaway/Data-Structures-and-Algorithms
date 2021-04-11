package main

import (
	"fmt"
	"math"
)

// Comparable interface allows different types
// to be compared and used in the heap data structure
type Comparable interface {
	GetComparable() int
}

type MinHeap struct {
	heap []Comparable
}

// helper function to create a new MinHeap
func CreateMinHeap(capacity int) *MinHeap {
	newHeap := new(MinHeap)
	newHeap.heap = make([]Comparable, 0, capacity)

	return newHeap
}

// helper function to get the indices of the left and right
// children of a node in the heap
func getChildIndices(pos int) (int, int) {
	return (2 * pos) + 1, (2 * pos) + 2
}

// Helper function returns the index of the child node
// with the largest value
func getIndexOfMaxChild(list []Comparable, pos int) (int, bool) {
	left, right := getChildIndices(pos)

	// if indices are both out of range return 0 with false flag
	if left >= len(list) && right >= len(list) {
		return 0, false
	}

	// left is in range but right is not
	if right >= len(list) {
		return left, true
	}

	// otherwise return the index of the bigger value
	if list[left].GetComparable() > list[right].GetComparable() {
		return left, true
	} else {
		return right, true
	}
}

// Helper function returns the index of the child node
// with the smallest value
func getIndexOfMinChild(list []Comparable, pos int) (int, bool) {
	left, right := getChildIndices(pos)

	// if indices are both out of range return 0 with false flag
	if left >= len(list) && right >= len(list) {
		return 0, false
	}

	// left is in range but right is not
	if right >= len(list) {
		return left, true
	}

	// otherwise return the index of the bigger value
	if list[left].GetComparable() < list[right].GetComparable() {
		return left, true
	} else {
		return right, true
	}
}

// helper function returns the parent index of a node
func getParentIndex(pos int) int {
	return (pos - 1) / 2
}

// Function takes a list and converts it into a minimum heap
func MinHeapify(list []Comparable) {
	firstNonLeafNode := int(math.Floor(float64(len(list) / 2)))
	for i := firstNonLeafNode; i >= 0; i-- {

		minHeapifyDown(list, i)
	}
}

// function takes the index of a parent node and recrsively
// moves downward in the heap and adjusts it to be a valid min heap
func minHeapifyDown(list []Comparable, index int) {
	left, right := getChildIndices(index)

	switch {
	case left < len(list) && right < len(list):
		indexOfMinChild, ok := getIndexOfMinChild(list, index)
		if !ok {
			return
		}

		if list[index].GetComparable() > list[indexOfMinChild].GetComparable() {
			list[index], list[indexOfMinChild] = list[indexOfMinChild], list[index]
			minHeapifyDown(list, indexOfMinChild)
		}
	case left < len(list):
		if list[index].GetComparable() > list[left].GetComparable() {
			list[index], list[left] = list[left], list[index]
			minHeapifyDown(list, left)
		}
	default:
		return
	}

}

// function takes the index of a parent node and recrsively
// moves upward in the heap and adjusts it to be a valid min heap
func minHeapifyUp(list []Comparable, index int) {
	parentIndex := getParentIndex(index)

	if parentIndex >= 0 {
		// if child is less than parent
		if list[index].GetComparable() < list[parentIndex].GetComparable() {
			list[index], list[parentIndex] = list[parentIndex], list[index]
			minHeapifyUp(list, parentIndex)
		}
	}
}

// Method to add a new value to the heap
func (h *MinHeap) Add(value Comparable) {
	h.heap = append(h.heap, value)
	minHeapifyUp(h.heap, len(h.heap)-1)
}

// Method to get the top value in the min heap
func (h *MinHeap) Pop() (Comparable, bool) {
	if len(h.heap) <= 0 {
		return nil, false
	}

	popValue := h.heap[0]
	h.heap[0] = h.heap[len(h.heap)-1]
	h.heap = h.heap[:len(h.heap)-1]
	minHeapifyDown(h.heap, 0)

	return popValue, true
}

// Person struct implements the Comparable interface and can be used
// in the heap
type Person struct {
	Name string
	Age  int
}

func (p *Person) GetComparable() int {
	return p.Age
}

func CreatePerson(name string, age int) *Person {
	return &Person{Name: name, Age: age}
}

// Testing the functionality of the Min Heap
func main() {
	michael := CreatePerson("Michael", 23)
	leah := CreatePerson("Leah", 22)
	jake := CreatePerson("jake", 19)
	tim := CreatePerson("tim", 12)
	larry := CreatePerson("larry", 20)
	lenny := CreatePerson("lenny", 21)
	junior := CreatePerson("junior", 10)

	personHeap := CreateMinHeap(10)
	personHeap.Add(michael)
	personHeap.Add(leah)
	personHeap.Add(jake)
	personHeap.Add(tim)
	personHeap.Add(larry)
	personHeap.Add(lenny)
	personHeap.Add(junior)

	value, ok := personHeap.Pop()
	for ok {
		fmt.Printf("Top Value: %v\n", value)
		value, ok = personHeap.Pop()
	}
}
