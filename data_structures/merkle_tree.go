package main

import (
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"strconv"
	"time"
)

// any type that implements the fmt.Stringer interface can be stored
// in the merkle tree
type Hashable interface {
	String() string
}

type MerkleNode struct {
	hash  string
	left  *MerkleNode
	right *MerkleNode
	data  fmt.Stringer
}

type MerkleTree struct {
	merkleRoot *MerkleNode
}

// method returns true if the node is a leaf node
func (mn *MerkleNode) isLeafNode() bool {
	return mn.data != nil && mn.left == nil && mn.right == nil
}

// method calculates and returns the hash of the merkle node
func (mn *MerkleNode) calcNodeHash() string {
	if mn.isLeafNode() {
		hashableString := mn.data.String()
		return hashSHA256(hashableString)
	}

	leftNodeHash := mn.left.hash
	var rightNodeHash string
	if mn.right != nil {
		rightNodeHash = mn.right.hash
	}

	return hashSHA256(leftNodeHash + rightNodeHash)
}

// function hashes a string using SHA256 and returns
// the Hex digest as a string
func hashSHA256(input string) string {
	h := sha1.New()
	h.Write([]byte(input))
	return hex.EncodeToString(h.Sum(nil))
}

// Function takes an array of stringers and builds
// a merkle tree from the data
func NewMerkleTree(data []fmt.Stringer) *MerkleTree {
	leafNodes := createLeafNodes(data)
	return buildMerkleTree(leafNodes)
}

// function takes an array of MerkleNode pointers and builds
// a builds a merkle tree by combining nodes until a
// root node is created
func buildMerkleTree(leafNodes []*MerkleNode) *MerkleTree {
	for len(leafNodes) > 1 {
		leafNodes = combineAdjacentMerkleNodes(leafNodes)
	}

	return &MerkleTree{merkleRoot: leafNodes[0]}
}

// function takes an array of MerkleNode pointers and returns
// new array where adjacent nodes have been combined
func combineAdjacentMerkleNodes(nodes []*MerkleNode) []*MerkleNode {
	combinedNodes := []*MerkleNode{}

	for i := 0; i < len(nodes); i += 2 {
		leftNode := nodes[i]
		var rightNode *MerkleNode
		if i+1 < len(nodes) {
			rightNode = nodes[i+1]
		}

		combinedNode := createInternalMerkleNode(leftNode, rightNode)
		combinedNodes = append(combinedNodes, combinedNode)
	}

	return combinedNodes
}

// Function to create an internal MerkleNode. Function takes two
// MerkleNodes which will be used as left and right children
func createInternalMerkleNode(left, right *MerkleNode) *MerkleNode {
	node := new(MerkleNode)
	node.left, node.right = left, right
	node.hash = node.calcNodeHash()

	return node
}

// function creates a leaf MerkleNode with the given data
func createLeafMerkleNode(data fmt.Stringer) *MerkleNode {
	node := new(MerkleNode)
	node.data = data
	node.hash = node.calcNodeHash()

	return node
}

// function takes an array of stringable data and creates
// a list of leaf MerkleNodes pointers
func createLeafNodes(data []fmt.Stringer) []*MerkleNode {
	var leafNodes []*MerkleNode

	for _, value := range data {
		node := createLeafMerkleNode(value)
		leafNodes = append(leafNodes, node)
	}

	return leafNodes
}

// function verifies that a merkle tree is valid
func VerifyMerkleTree(tree *MerkleTree) bool {
	return verifyMerkleTreeHelper(tree.merkleRoot)
}

// function verifies MerkleTree by recursively calculating node
// hashes
func verifyMerkleTreeHelper(node *MerkleNode) bool {
	if node == nil {
		return true
	}

	// verify hash of data at leaf node
	if node.isLeafNode() {
		hashOfData := hashSHA256(node.data.String())
		if hashOfData == node.hash {
			return true
		}
		return false
	}

	// verify internal nodes
	leftNode, rightNode := node.left, node.right
	leftHash := leftNode.hash
	var rightHash string
	if rightNode != nil {
		rightHash = rightNode.hash
	}

	hashOfChildNode := hashSHA256(leftHash + rightHash)
	if hashOfChildNode == node.hash {
		return verifyMerkleTreeHelper(leftNode) && verifyMerkleTreeHelper(rightNode)
	}

	return false
}

// Transaction struct implements the Stringer interface
type Transaction struct {
	senderId   string
	receiverId string
	datetime   string
	amount     float64
}

func (t *Transaction) String() string {
	amountAsString := strconv.FormatFloat(t.amount, 'f', 2, 64)
	transactionString := t.senderId + t.receiverId + t.datetime + amountAsString

	return transactionString
}

func createTransaction(sId, rId string, amt float64) *Transaction {
	t := new(Transaction)
	t.senderId, t.receiverId = sId, rId
	t.datetime = time.Now().String()
	t.amount = amt

	return t
}

func main() {
	// create some transactions
	t1 := createTransaction("Michael", "John", 12.3)
	t2 := createTransaction("Ryan", "Michael", 10.67)
	t3 := createTransaction("Ben", "Ryan", 14.53)
	t4 := createTransaction("John", "Ben", 13.3)
	t5 := createTransaction("Kaei", "John", 15.3)

	// store the transaction data in a merkle tree
	transactions := []fmt.Stringer{t1, t2, t3, t4, t5}
	tree := NewMerkleTree(transactions)

	fmt.Println("### Verifying the merkel tree ###")
	isValid := VerifyMerkleTree(tree)
	fmt.Printf("Is the tree valid -> %v\n", isValid)

	// modify some of the data
	fmt.Println("Modifying data at leaf merkle nodes")
	t3.senderId = "Ryan"
	t3.receiverId = "Ben"

	fmt.Println("### Verifying the merkel tree ###")
	isValid = VerifyMerkleTree(tree)
	fmt.Printf("Is the tree valid -> %v\n", isValid)
}
