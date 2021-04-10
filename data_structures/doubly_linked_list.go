package main

import "fmt"

type Node struct {
	value interface{}
	prev  *Node
	next  *Node
}

type DoublyLinkedList struct {
	head   *Node
	tail   *Node
	length int
}

// helper function to create a new node with a given value
func createNode(value interface{}) *Node {
	return &Node{value: value}
}

// helper function to create a new doubly linked list
// and initialize the head and tail sentinel nodes
func CreateLinkedList() *DoublyLinkedList {
	ll := new(DoublyLinkedList)

	ll.head = createNode(nil)
	ll.tail = createNode(nil)
	ll.head.next = ll.tail
	ll.tail.prev = ll.head
	ll.length = 0

	return ll
}

// method to print the contents of the linked list on a single line
func (ll *DoublyLinkedList) Print() {
	currentNode := ll.head.next
	for currentNode != ll.tail {
		fmt.Printf("%v ", currentNode.value)
		currentNode = currentNode.next
	}
	fmt.Printf("\n")
}

// method to add new value to the end of the linked list
func (ll *DoublyLinkedList) Append(value interface{}) {
	newNode := createNode(value)

	newNode.prev = ll.tail.prev
	newNode.next = ll.tail

	newNode.prev.next = newNode
	ll.tail.prev = newNode
	ll.length += 1
}

// method to add value to the front of the linked list
func (ll *DoublyLinkedList) Prepend(value interface{}) {
	newNode := createNode(value)

	newNode.prev = ll.head
	newNode.next = ll.head.next

	newNode.next.prev = newNode
	ll.head.next = newNode
	ll.length += 1
}

// method to remove the fisrt value from the list and
// return it along with a boolean indicating whether or not
// the operation was successful
func (ll *DoublyLinkedList) PopFront() (interface{}, bool) {
	if ll.head.next == ll.tail {
		return nil, false
	}
	// get the first node and its value
	firstNode := ll.head.next
	returnValue := firstNode.value

	// remove the first node
	ll.head.next = firstNode.next
	firstNode.next.prev = ll.head
	firstNode.next, firstNode.prev = nil, nil

	ll.length -= 1

	return returnValue, true
}

// method to remove the last value from the list and
// return it along with a boolean indicating whether or not
// the operation was successful
func (ll *DoublyLinkedList) PopBack() (interface{}, bool) {
	if ll.tail.prev == ll.head {
		return nil, false
	}

	// get last node and its value
	lastNode := ll.tail.prev
	returnValue := lastNode.value

	// remove the node
	ll.tail.prev = lastNode.prev
	lastNode.prev.next = ll.tail
	lastNode.next, lastNode.prev = nil, nil

	ll.length -= 1

	return returnValue, true
}

// method to check if the list contains a given value
func (ll *DoublyLinkedList) Contains(value interface{}) bool {
	currentNode := ll.head.next
	for currentNode != ll.tail {
		if currentNode.value == value {
			return true
		}

		currentNode = currentNode.next
	}
	return false
}

// method to reverse the order of the doubly linked list
func (ll *DoublyLinkedList) Reverse() {
	if ll.length <= 1 {
		return
	}

	previousNode := ll.head
	currentNode := ll.head.next
	for currentNode != ll.tail {
		nextNode := currentNode.next

		// switch the pointers
		currentNode.next = previousNode
		currentNode.prev = nextNode

		previousNode = currentNode
		currentNode = nextNode
	}

	// adjust head and tail nodes
	temp := ll.head.next
	ll.head.next = ll.tail.prev
	ll.head.next.prev = ll.head
	ll.tail.prev = temp
	ll.tail.prev.next = ll.tail
}

// testing functionality of DoublyLinkedList
func main() {
	ll := CreateLinkedList()

	ll.Append(100)
	ll.Append("Michael")
	ll.Prepend("Tom")
	ll.Prepend(12.67)
	ll.Prepend("Jake")
	ll.Prepend(-7)

	fmt.Println("Before Reverse:")
	ll.Print()

	ll.Reverse()

	fmt.Println("After Reverse:")
	ll.Print()
}
