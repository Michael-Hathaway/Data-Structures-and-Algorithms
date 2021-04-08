// Implementation of Binary Search Tree data structure in Go

package main

import "fmt"

// Represents an individual node within the tree
type Node struct {
	value int
	left  *Node
	right *Node
}

// Helper function for creating a new Node
func createNode(value int) *Node {
	node := new(Node)
	node.value = value

	return node
}

type Tree struct {
	root *Node
}

// method adds a new Node to the Tree, following the bst rules
func (t *Tree) Add(value int) {
	if t.root == nil {
		t.root = createNode(value)
	} else {
		t.root = addHelper(t.root, value)
	}
}

// Helper function that recursively traverses the tree
// and inserts a new node in the correct position
func addHelper(node *Node, value int) *Node {
	if node == nil {
		return createNode(value)
	} else if value > node.value {
		node.right = addHelper(node.right, value)
	} else {
		node.left = addHelper(node.left, value)
	}

	return node
}

// method to print the values in the tree
// based on an in-order traversal
func (t *Tree) PrintInOrder() {
	printInOrderHelper(t.root)
}

// helper function that recursively traverses the tree
// in an in-order fashion and prints the node values
// to the screen
func printInOrderHelper(node *Node) {
	if node != nil {
		printInOrderHelper(node.left)
		fmt.Println(node.value)
		printInOrderHelper(node.right)
	}
}

// method to check if the tree contains a certain value
// returns true if the value is in the tree and false if not
func (t *Tree) Contains(value int) bool {
	return containsHelper(t.root, value)
}

// helper function to recursively check if the tree
// contains the given value
func containsHelper(node *Node, value int) bool {
	switch {
	case node == nil:
		return false
	case value == node.value:
		return true
	case value > node.value:
		return containsHelper(node.right, value)
	case value < node.value:
		return containsHelper(node.left, value)
	default:
		return false
	}
}

// method to remove a given value from the tree
func (t *Tree) remove(value int) {
	if t.Contains(value) {
		t.root = removeHelper(t.root, value)
	}
}

// helper function to recursively traverse the tree to find
// where to remove the node
func removeHelper(node *Node, value int) *Node {
	switch {
	case node == nil:
		return node
	case value == node.value:
		node = removeNode(node)
		return node
	case value > node.value:
		node.right = removeHelper(node.right, value)
		return node
	case value < node.value:
		node.left = removeHelper(node.left, value)
		return node
	default:
		return nil
	}
}

// helper function to remove the node and restructure
// the tree
func removeNode(toRemove *Node) *Node {
	switch {
	case toRemove.left == nil && toRemove.right == nil:
		// no left or right child -> return nil
		return nil
	case toRemove.left == nil && toRemove.right != nil:
		// only right child -> return right child
		return toRemove.right
	case toRemove.left != nil && toRemove.right == nil:
		// only left child -> return left child
		return toRemove.left
	case toRemove.left != nil && toRemove.right != nil:
		// if both right and left child, find smallest value
		// in right subtree and replace the value at the current node
		newValue := getValueOfLeftmostNode(toRemove.right)
		toRemove.value = newValue
		toRemove.right = removeHelper(toRemove.right, newValue)
		return toRemove
	default:
		return nil
	}
}

// helper function to find the value of the leftmost
// node in a given subtree
func getValueOfLeftmostNode(node *Node) int {
	currentNode := node

	// traverse as far left into the tree
	// to find the leftmost node
	for currentNode.left != nil {
		currentNode = currentNode.left
	}

	return currentNode.value
}

// Code to test the functionality of the Tree data structure
func main() {
	tree := Tree{}
	tree.Add(7)
	tree.Add(6)
	tree.Add(9)
	tree.Add(8)
	tree.Add(70)
	tree.Add(32)
	tree.Add(-7)

	tree.remove(7)

	tree.PrintInOrder()

	fmt.Printf("Contains the value 10: %v\n", tree.Contains(10))
	fmt.Printf("Contains the value 70: %v\n", tree.Contains(70))
}
