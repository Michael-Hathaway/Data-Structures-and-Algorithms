// PROMPT
// Given a non-empty array of integers representing the pre-order
// traversal of a binary search tree, write a function that
// reconstructs the bst and returns the root node
//

package main

import "fmt"

type BST struct {
	Value int

	Left  *BST
	Right *BST
}

// helper function to create a new BST
func CreateBST(value int) *BST {
	return &BST{Value: value}
}

// helper function to return the index of the first item
// in an array that is greater than or equal to n
func GetIndexOfFirstValueGreaterOrEqualToN(n int, array []int) int {
	for index, value := range array {
		if value >= n {
			return index
		}
	}

	return -1
}

// Function to recreate BST from the list of pre order traversl values
func ReconstructBst(preOrderTraversalValues []int) *BST {
	// If we run out of values we return nil
	if len(preOrderTraversalValues) <= 0 {
		return nil
	}

	// value for new node will be first in the list
	currentValue := preOrderTraversalValues[0]
	// update the rest of the array
	preOrderTraversalValues = preOrderTraversalValues[1:]

	// create new bst with currentValue
	bst := CreateBST(currentValue)

	// to reconstruct the left and right sub trees we need to find
	// all the values less than the current one value the left tree
	// and all values greater than the current value for the right tree
	n := GetIndexOfFirstValueGreaterOrEqualToN(currentValue, preOrderTraversalValues)

	var leftSubTreeValues []int
	var rightSubTreeValues []int
	if n == -1 {
		// no value greater than current, so right is empty array
		leftSubTreeValues = preOrderTraversalValues
		rightSubTreeValues = []int{}
	} else {
		// split the traversl values for left and right trees
		leftSubTreeValues = preOrderTraversalValues[:n]
		rightSubTreeValues = preOrderTraversalValues[n:]
	}

	// recursively reconstruct the left and right subtrees
	bst.Left = ReconstructBst(leftSubTreeValues)
	bst.Right = ReconstructBst(rightSubTreeValues)

	return bst
}

// function prints the in order traversal of the bst
func PrintInOrder(bst *BST) {
	if bst != nil {
		PrintInOrder(bst.Left)
		fmt.Println(bst.Value)
		PrintInOrder(bst.Right)
	}
}

func main() {
	preOrderTraversalValues := []int{10, 4, 2, 1, 5, 17, 19, 18}
	bst := ReconstructBst(preOrderTraversalValues)

	// use inorder print to validate
	PrintInOrder(bst)
}
