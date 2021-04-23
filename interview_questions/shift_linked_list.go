/*
PROMPT
Write a function that takes the head of a linked list and an integer k,
shifts the head of the linked list by k and returns the new head.
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

func ShiftLinkedList(head *LinkedList, k int) *LinkedList {
	if k >= 0 {
		return ShiftListForwardByK(head, k)
	}
	return ShiftListBackwardByK(head, k)
}

func ShiftListForwardByK(head *LinkedList, k int) *LinkedList {
	length := GetLengthOfLinkedList(head)
	if length < 2 {
		return head
	}

	if k%length == 0 {
		return head
	}

	posOfNewHead := length - (k % length)
	lastNode := GetNodeAtPosition(head, length-1)
	newHead := GetNodeAtPosition(head, posOfNewHead)
	nodeBeforeNewHead := GetNodeAtPosition(head, posOfNewHead-1)

	nodeBeforeNewHead.Next = nil
	lastNode.Next = head
	return newHead
}

func ShiftListBackwardByK(head *LinkedList, k int) *LinkedList {
	length := GetLengthOfLinkedList(head)
	if length < 2 {
		return head
	}

	posOfNewHead := (0 - k) % length
	if posOfNewHead == 0 {
		return head
	}

	lastNode := GetNodeAtPosition(head, length-1)
	newHead := GetNodeAtPosition(head, posOfNewHead)
	nodeBeforeNewHead := GetNodeAtPosition(head, posOfNewHead-1)

	nodeBeforeNewHead.Next = nil
	lastNode.Next = head
	return newHead
}

func GetNodeAtPosition(head *LinkedList, pos int) *LinkedList {
	itr := head
	counter := 0
	for itr != nil {
		if counter == pos {
			return itr
		}

		counter += 1
		itr = itr.Next
	}

	return nil
}

func GetLengthOfLinkedList(head *LinkedList) int {
	itr := head
	counter := 0
	for itr != nil {
		counter += 1
		itr = itr.Next
	}

	return counter
}

func main() {
	array := []int{1, 3, 4, 5, 9, 10}
	linkedList := CreateLinkedListFromArray(array)

	fmt.Println("Before Shift:")
	PrintLinkedList(linkedList)

	linkedList = ShiftLinkedList(linkedList, 3)

	fmt.Println("After Shift:")
	PrintLinkedList(linkedList)

	linkedList = ShiftLinkedList(linkedList, -3)

	fmt.Println("After Shifting Back To Initial:")
	PrintLinkedList(linkedList)
}
