// Implementation of AVL Tree data structure in Go

package main

import (
	"fmt"
	"math"
)

// Represents an individual node within the tree
type Node struct {
	value  int
	height int
	left   *Node
	right  *Node
}

type AVLTree struct {
	root *Node
}

// NODE METHODS AND FUNCTIONS

// function updates a nodes height to 1 + the max height of
// the right or left sub trees
func (n *Node) updateHeight() {
	leftTreeHeight, rightTreeHeight := -1, -1

	if n.left != nil {
		leftTreeHeight = n.left.height
	}

	if n.right != nil {
		rightTreeHeight = n.right.height
	}

	n.height = 1 + int(math.Max(float64(leftTreeHeight), float64(rightTreeHeight)))
}

// method returns the balance factor for the node
func (n *Node) getBalanceFactor() int {
	leftTreeHeight, rightTreeHeight := -1, -1

	if n.left != nil {
		leftTreeHeight = n.left.height
	}

	if n.right != nil {
		rightTreeHeight = n.right.height
	}

	return leftTreeHeight - rightTreeHeight
}

// helper function to rotate a node to the left
func rotateNodeLeft(root *Node) *Node {
	oldRoot := root
	newRoot := root.right

	// perform rotation
	oldRoot.right = newRoot.left
	newRoot.left = oldRoot

	// update heights
	oldRoot.updateHeight()
	newRoot.updateHeight()

	return newRoot
}

// helper function to rotate a node to the left
func rotateNodeRight(root *Node) *Node {
	oldRoot := root
	newRoot := root.left

	// perform rotation
	oldRoot.left = newRoot.right
	newRoot.right = oldRoot

	// update heights
	oldRoot.updateHeight()
	newRoot.updateHeight()

	return newRoot
}

// function to balance a Node. returns a pointer to the balanced Node
func balanceNode(node *Node) *Node {
	balanceFactor := node.getBalanceFactor()

	switch {
	case balanceFactor >= 2:
		// check to see if double rotation is needed
		if node.left.getBalanceFactor() <= -1 {
			node.left = rotateNodeLeft(node.left)
		}
		// right rotate
		node = rotateNodeRight(node)
		return node
	case balanceFactor <= -2:
		// check to see if double rotation is needed
		if node.right.getBalanceFactor() >= 1 {
			node.right = rotateNodeRight(node.right)
		}
		// left rotate
		node = rotateNodeLeft(node)
		return node
	default:
		return node
	}
}

// Helper function for creating a new Node
func createNode(value int) *Node {
	node := new(Node)
	node.value = value
	node.height = 0

	return node
}

// AVL TREE FUNCTIONS AND METHODS

// method adds a new Node to the AVLTree, following the bst rules
func (t *AVLTree) Add(value int) {
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
		node = createNode(value)
	} else if value > node.value {
		node.right = addHelper(node.right, value)
	} else {
		node.left = addHelper(node.left, value)
	}

	node.updateHeight()
	return balanceNode(node)
}

// method to print the values in the tree
// based on an in-order traversal
func (t *AVLTree) PrintInOrder() {
	printInOrderHelper(t.root)
}

// helper function that recursively traverses the tree
// in an in-order fashion and prints the node values
// to the screen
func printInOrderHelper(node *Node) {
	if node != nil {
		printInOrderHelper(node.left)
		// heights are printed to verify that the tree structure
		// is changing to maintain balance
		fmt.Printf("Node Value: %v, Node Height: %v\n", node.value, node.height)
		printInOrderHelper(node.right)
	}
}

// method to check if the tree contains a certain value
// returns true if the value is in the tree and false if not
func (t *AVLTree) Contains(value int) bool {
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
func (t *AVLTree) remove(value int) {
	if t.Contains(value) {
		t.root = removeHelper(t.root, value)
	}
}

// helper function to recursively traverse the tree to find
// where to remove the node
func removeHelper(node *Node, value int) *Node {
	switch {
	case node == nil:
		node = nil
	case value == node.value:
		node = removeNode(node)
	case value > node.value:
		node.right = removeHelper(node.right, value)
	case value < node.value:
		node.left = removeHelper(node.left, value)
	default:
		node = nil
	}

	if node != nil {
		node.updateHeight()
		node = balanceNode(node)
	}

	return node
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

// Code to test the functionality of the AVLTree data structure
func main() {
	tree := AVLTree{}
	tree.Add(7)
	tree.Add(5)
	tree.Add(10)
	tree.Add(8)

	tree.remove(5)

	tree.PrintInOrder()
}
