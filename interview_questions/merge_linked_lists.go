/*
Prompt:
Write a function that merges two sorted linked list into a single sorted list
*/

package main

import "fmt"

type LinkedList struct {
	Value int
	Next  *LinkedList
}

func NewLinkedList(value int) *LinkedList {
	return &LinkedList{Value: value}
}

func PrintLinkedList(node *LinkedList) {
	itr := node
	for itr != nil {
		fmt.Printf("%v ", itr.Value)
		itr = itr.Next
	}
	fmt.Printf("\n")
}

func CreateLinkedListFromArray(array []int) *LinkedList {
	if len(array) == 0 {
		return nil
	}

	firstNode := NewLinkedList(array[0])
	itr := firstNode

	for i := 1; i < len(array); i++ {
		newNode := NewLinkedList(array[i])
		itr.Next = newNode
		itr = itr.Next
	}

	return firstNode
}

func MergeLinkedLists(headOne *LinkedList, headTwo *LinkedList) *LinkedList {
	h1 := headOne
	h2 := headTwo
	var prevH1 *LinkedList

	for h1 != nil && h2 != nil {
		// we will merge h2 into h1

		// if h1 is less than h2, we can just move h1
		// forward but keep track of the previous h1
		if h1.Value < h2.Value {
			prevH1 = h1
			h1 = h1.Next
		} else {
			// if h2 is less than h1, then the prevH1.Next will
			// will become the current h2
			if prevH1 != nil {
				prevH1.Next = h2
			}

			prevH1 = h2      // previous will now be h2
			h2 = h2.Next     // increment h2
			prevH1.Next = h1 // newly added h2 node points to the rest of h1
		}
	}

	// if h1 finished first, we need to make sure we get the rest of h1
	if h1 == nil {
		prevH1.Next = h2
	}

	if headOne.Value < headTwo.Value {
		return headOne
	}
	return headTwo

}

func main() {
	array1 := []int{2, 6, 7, 8}
	array2 := []int{1, 3, 4, 5, 9, 10}

	listOne := CreateLinkedListFromArray(array1)
	listTwo := CreateLinkedListFromArray(array2)

	mergedList := MergeLinkedLists(listOne, listTwo)
	PrintLinkedList(mergedList)
}
