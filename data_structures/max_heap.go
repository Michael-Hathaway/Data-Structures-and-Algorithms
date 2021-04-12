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

type MaxHeap struct {
	heap []Comparable
}

// helper function to create a new MaxHeap
func CreateMaxHeap(capacity int) *MaxHeap {
	newHeap := new(MaxHeap)
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

func HeapSort(list []Comparable) {
	// create a max heap out of the input data
	MaxHeapify(list)

	front, end := 0, len(list)-1
	// move max value to the end of the list and heapify
	// up to but not including the last value. This results
	// in a sorted list
	for end > 0 {
		list[front], list[end] = list[end], list[front]
		end -= 1
		maxHeapifyDown(list, front, end)
	}
}

// converts a list of comparables to a max heap
func MaxHeapify(list []Comparable) {
	firstNonLeafNode := int(math.Floor(float64(len(list) / 2)))
	for i := firstNonLeafNode; i >= 0; i-- {

		maxHeapifyDown(list, i, len(list)-1)
	}
}

func maxHeapifyDown(list []Comparable, pos, end int) {
	left, right := getChildIndices(pos)

	switch {
	case left <= end && right <= end:
		indexOfMaxChild, ok := getIndexOfMaxChild(list, pos)
		if !ok {
			return
		}

		if list[pos].GetComparable() < list[indexOfMaxChild].GetComparable() {
			list[pos], list[indexOfMaxChild] = list[indexOfMaxChild], list[pos]
			maxHeapifyDown(list, indexOfMaxChild, end)
		}
	case left <= end:
		if list[pos].GetComparable() < list[left].GetComparable() {
			list[pos], list[left] = list[left], list[pos]
			maxHeapifyDown(list, left, end)
		}
	default:
		return
	}

}

func maxHeapifyUp(list []Comparable, index int) {
	parentIndex := getParentIndex(index)

	if parentIndex >= 0 {
		// if child is less than parent
		if list[index].GetComparable() > list[parentIndex].GetComparable() {
			list[index], list[parentIndex] = list[parentIndex], list[index]
			maxHeapifyUp(list, parentIndex)
		}
	}
}

func (h *MaxHeap) PrintHeap() {
	for i := 0; i < len(h.heap); i++ {
		fmt.Println(h.heap[i])
	}
}

func (h *MaxHeap) Add(value Comparable) {
	h.heap = append(h.heap, value)
	maxHeapifyUp(h.heap, len(h.heap)-1)
}

func (h *MaxHeap) Pop() (Comparable, bool) {
	if len(h.heap) <= 0 {
		return nil, false
	}

	popValue := h.heap[0]
	h.heap[0] = h.heap[len(h.heap)-1]
	h.heap = h.heap[:len(h.heap)-1]
	maxHeapifyDown(h.heap, 0, len(h.heap)-1)

	return popValue, true
}

// Person implements the Comparable interface and can be used
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

// Testing Max Heap and HeapSort implementations
func main() {
	michael := CreatePerson("Michael", 23)
	leah := CreatePerson("Leah", 22)
	jake := CreatePerson("jake", 19)
	tim := CreatePerson("tim", 12)
	larry := CreatePerson("larry", 20)
	lenny := CreatePerson("lenny", 21)
	junior := CreatePerson("junior", 10)

	personList := []Comparable{michael, leah, jake, tim, larry, lenny, junior}

	// HEAPSORT
	fmt.Println("### Testing HeapSort Implementation ###")
	fmt.Println("Before Sorting:")
	for _, value := range personList {
		fmt.Println(value)
	}

	HeapSort(personList)

	fmt.Println("\nAfter Sorting:")
	for _, value := range personList {
		fmt.Println(value)
	}

	fmt.Printf("\n### Constructing Max Heap ###\n")
	personHeap := CreateMaxHeap(10)
	personHeap.Add(michael)
	personHeap.Add(leah)
	personHeap.Add(jake)
	personHeap.Add(tim)
	personHeap.Add(larry)
	personHeap.Add(lenny)
	personHeap.Add(junior)

	fmt.Println("Popping values from top of Max Heap")
	value, ok := personHeap.Pop()
	for ok {
		fmt.Printf("Top Value: %v\n", value)
		value, ok = personHeap.Pop()
	}
}
