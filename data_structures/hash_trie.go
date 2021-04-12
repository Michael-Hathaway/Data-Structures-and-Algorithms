// Map abstract data type implemented as a hash trie
package main

import "fmt"

type TrieNode struct {
	key      rune
	value    interface{}
	children map[rune]*TrieNode
}

type HashTrie struct {
	root *TrieNode
}

/*
func createTrieNode
Desc: helper function for creating a new TrieNode

:param key rune - character key for this trie node
:param value interface{} - value being stored at the given key
*/
func createTrieNode(key rune, value interface{}) *TrieNode {
	return &TrieNode{key: key, value: value, children: map[rune]*TrieNode{}}
}

/*
func NewTrie
Desc: Function creates a new HashTrie and returns a pointer to it

:return *HashTrie - a pointer to a new and initilialzed trie data structure
*/
func NewHashTrie() *HashTrie {
	trie := new(HashTrie)
	trie.root = createTrieNode('0', nil)

	return trie
}

/*
func (tn *TrieNode) getChildWithKey
Desc:

:param key rune - the key for the child to be retrieved

:return *TrieNode - the TrieNode with the given key
										value will be nil if key doesn't exist
*/
func (tn *TrieNode) getChildWithKey(key rune) *TrieNode {
	value, ok := tn.children[key]
	if !ok {
		return nil
	}
	return value
}

/*
func (t *True) Insert
Desc: Method for inserting a key value pair into the trie

:param key string - the key to store the value at
:param value interface{} - the value to store at the given key

:return bool - was the insert successful
*/
func (ht *HashTrie) Add(key string, value interface{}) bool {
	if len(key) == 0 {
		return false
	}

	return addHelper(ht.root, key, value)
}

/*
func (t *True) addHelper
Desc: Helper Method to recursivelt insert a key value pair into the trie

:param node *TrieNode - pointer to the current node in the trie structure
:param key string - the key to store the value at
:param value interface{} - the value to store at the given key

:return bool - was the insert successful
*/
func addHelper(node *TrieNode, key string, value interface{}) bool {
	firstLetterInKey := rune(key[0])
	nextNode := node.getChildWithKey(firstLetterInKey)

	if nextNode == nil {
		nextNode = createTrieNode(firstLetterInKey, nil)
		node.children[firstLetterInKey] = nextNode
	}

	if len(key) == 1 {
		nextNode.value = value
		return true
	}

	key = key[1:]
	return addHelper(nextNode, key, value)
}

/*
func (t *True) Get
Desc: Method to retrieve a value stored at a given key

:param key string - the key whose value you want to retrieve

:return interface{} - the value stored at the given key, if the key is
											invalid the return value will be nil
*/
func (ht *HashTrie) Get(key string) (interface{}, bool) {
	return getHelper(ht.root, key)
}

/*
func (t *True) getHelper
Desc: Helper Method to recrusively retrieve the value stored at
a given key in the trie data structure

:param node *TrieNode - pointer to the current node in the trie structure
:param key string - the key whose value you want to retrieve

:return interface{} - the value stored at the given key, if the key is
invalid the return value will be nil
*/
func getHelper(node *TrieNode, key string) (interface{}, bool) {
	firstLetterInKey := rune(key[0])
	nextNode := node.getChildWithKey(firstLetterInKey)

	if nextNode == nil {
		return nil, false
	}

	if len(key) == 1 {
		if nextNode.value == nil {
			return nil, false
		}

		return nextNode.value, true
	}

	key = key[1:]
	return getHelper(nextNode, key)
}

/*
func (ht *HashTrie) Delete
Desc: Delete a key value pair from the HashTrie

:param key string - the key identifying the record to be deleted

:return bool - true if the delete was successful, false if not
*/
func (ht *HashTrie) Delete(key string) bool {
	if len(key) == 0 {
		return false
	}

	return deleteHelper(ht.root, key)
}

/*
func deleteHelper
Desc: Recursivelt navigates the try structure to find the
key value pair to be deleted

:param node *TrieNode - the current node in the trie
:param key string - the key identifying the record to be deleted

:return bool - true if the delete was successful, false if not
*/
func deleteHelper(node *TrieNode, key string) bool {
	firstLetterInKey := rune(key[0])
	nextNode := node.getChildWithKey(firstLetterInKey)

	if nextNode == nil {
		return false
	}

	if len(key) == 1 {
		nextNode.value = nil
		return true
	}

	key = key[1:]
	return deleteHelper(nextNode, key)
}

/*
func (ht *HashTrie) GetKeysWithPrefix
Desc: Returns a list of all the keys contained in the Trie that
start with the given prefix

:param prefix string - the key prefix being searched for

:return []string - returns a list of all the strings with the
given prefix
*/
func (ht *HashTrie) GetKeysWithPrefix(prefix string) []string {
	nodeWithPrefix := ht.getNodeWithPrefix(prefix)
	var keys []string
	getKeysWithPrefix(nodeWithPrefix, prefix, &keys)

	return keys
}

/*
func getKeysWithPrefix
Desc: recursive helper function to find all the keys that start with
the given prefix

:param node *TrieNode - the current node being visited in the trie
:prefix string - the prefix being searched for
:keys *[]string - pointer to a list of strings that will be filled with
all the keys containging the given prefix
*/
func getKeysWithPrefix(node *TrieNode, prefix string, keys *[]string) {
	if node.value != nil {
		*keys = append(*keys, prefix)
	}

	for key, childNode := range node.children {
		if childNode == nil {
			continue
		}

		newPrefix := prefix + string(key)
		getKeysWithPrefix(childNode, newPrefix, keys)
	}
}

/*
func (ht *HashTrie) getNodeWithPrefix
Desc: Function returns the node with the given prefix. This node
serves as the starting point for finding all keys that start with the
prefix

:param prefix string - the key prefix being searched for

:return *TrieNode - the node in the trie containing the given prefix.
The value is nil if the prefix is not contained in the trie
*/
func (ht *HashTrie) getNodeWithPrefix(prefix string) *TrieNode {
	if len(prefix) == 0 {
		return nil
	}

	return getNodeWithPrefixHelper(ht.root, prefix)
}

/*
func getNodeWithPrefixHelper
Desc: Recursive helper Function that returns the node with the given
prefix. This node serves as the starting point for finding all keys that
start with the prefix.

:param node *TrieNode - the current node being visited in the trie
:param prefix string - the key prefix being searched for

:return *TrieNode - the node in the trie containing the given prefix.
The value is nil if the prefix is not contained in the trie
*/
func getNodeWithPrefixHelper(node *TrieNode, prefix string) *TrieNode {
	firstLetterInKey := rune(prefix[0])
	nextNode := node.getChildWithKey(firstLetterInKey)

	if nextNode == nil {
		return nil
	}

	if len(prefix) == 1 {
		return nextNode
	}

	prefix = prefix[1:]
	return getNodeWithPrefixHelper(nextNode, prefix)
}

// Testing functionality of the HashTrie
func main() {
	trie := NewHashTrie()
	trie.Add("Michael", 23)
	trie.Add("Mike", 28)
	trie.Add("Luna", 25)
	trie.Add("Lee", 20)
	trie.Add("Mark", 21)

	fmt.Println("Retrieving values stored at: Mike")
	value, ok := trie.Get("Mike")
	if !ok {
		fmt.Println("Could not retrieve value stored at key: Michael")
	} else {
		fmt.Println("Value ->", value)
	}

	fmt.Println("Printing all keys that start with M:")
	fmt.Println(trie.GetKeysWithPrefix("M"))
}
