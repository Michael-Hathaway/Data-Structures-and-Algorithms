// implementation of a Singly Linked List in Go

package main

import "fmt"

type ListNode struct {
	value int
	next  *ListNode
}

// helper function to create a new ListNode
func createListNode(value int) *ListNode {
	return &ListNode{value: value}
}

type SinglyLinkedList struct {
	head   *ListNode
	length int
}

// method to print the node in the linked list
func (ll *SinglyLinkedList) Print() {
	itr := ll.head
	for itr != nil {
		fmt.Printf("%v ", itr.value)
		itr = itr.next
	}

	fmt.Printf("\n")
}

// method add value at the end of the linked list
func (ll *SinglyLinkedList) Append(value int) {
	if ll.head == nil {
		// if head is nil, create first node
		ll.head = createListNode(value)
	} else {
		// otherwise iterate to end of list and insert
		itr := ll.head
		for itr.next != nil {
			itr = itr.next
		}

		itr.next = createListNode(value)
	}

	ll.length += 1
}

// method adds value at the beginning of the linked list
func (ll *SinglyLinkedList) Prepend(value int) {
	newNode := createListNode(value)
	newNode.next = ll.head
	ll.head = newNode
	ll.length += 1
}

// method to insert value at a given position in the linked list
func (ll *SinglyLinkedList) Insert(value, index int) bool {
	if index > ll.length {
		return false
	}

	if index == 0 {
		ll.Prepend(value)
		return true
	}

	if index == ll.length {
		ll.Append(value)
		return true
	}

	position := 1
	previousNode := ll.head
	currentNode := ll.head.next

	for currentNode != nil {
		if position == index {
			newNode := createListNode(value)
			newNode.next = currentNode
			previousNode.next = newNode
			ll.length += 1
			return true
		}

		position += 1
		previousNode = previousNode.next
		currentNode = currentNode.next
	}

	return false
}

// method to check if a given value is contained
// within the linked list, returns true if the value is
// in the list and false if not
func (ll *SinglyLinkedList) Contains(value int) bool {
	currentNode := ll.head
	for currentNode != nil {
		if currentNode.value == value {
			return true
		}

		currentNode = currentNode.next
	}

	return false
}

// method reverses a linked linked list in place
func (ll *SinglyLinkedList) Reverse() {
	if ll.length <= 1 {
		return
	}

	var previousNode *ListNode = nil
	currentNode := ll.head

	for currentNode != nil {
		nextNode := currentNode.next

		// reverse the next pointer
		currentNode.next = previousNode

		previousNode = currentNode
		currentNode = nextNode
	}

	ll.head = previousNode
}

// Testing the functionality of the linked list
func main() {
	ll := SinglyLinkedList{}

	ll.Append(7)
	ll.Append(8)
	ll.Append(9)
	ll.Append(10)

	ll.Prepend(6)
	ll.Prepend(5)

	ll.Insert(4, 0)
	ll.Insert(11, 7)

	fmt.Println("Before calling reverse:")
	ll.Print()

	ll.Reverse()

	fmt.Println("After calling reverse:")
	ll.Print()
}
