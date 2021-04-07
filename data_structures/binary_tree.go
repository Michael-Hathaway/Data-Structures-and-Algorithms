package main

import "fmt"

type Node struct {
	value int
	left  *Node
	right *Node
}

func createNode(value int) *Node {
	node := new(Node)
	node.value = value

	return node
}

type Tree struct {
	root *Node
}

func (t *Tree) Add(value int) {
	if t.root == nil {
		t.root = createNode(value)
	} else {
		t.root = addHelper(t.root, value)
	}
}

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

func (t *Tree) PrintInOrder() {
	printInOrderHelper(t.root)
}

func printInOrderHelper(node *Node) {
	if node != nil {
		printInOrderHelper(node.left)
		fmt.Println(node.value)
		printInOrderHelper(node.right)
	}
}

func main() {
	tree := Tree{}
	tree.Add(7)
	tree.Add(8)
	tree.Add(9)

	tree.PrintInOrder()
}
