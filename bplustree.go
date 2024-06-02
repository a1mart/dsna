package main

import (
	"fmt"
)

// Node represents a node in the B+ tree
type Node struct {
	isLeaf    bool
	keys      []int
	children  []*Node
	nextLeaf  *Node // Pointer to the next leaf node (used for range queries)
}

// BPlusTree represents the B+ tree
type BPlusTree struct {
	root *Node
}

// NewBPlusTree creates a new B+ tree
func NewBPlusTree() *BPlusTree {
	return &BPlusTree{
		root: &Node{isLeaf: true},
	}
}

// Insert inserts a key into the B+ tree
func (tree *BPlusTree) Insert(key int) {
	if len(tree.root.keys) == 0 {
		tree.root.keys = append(tree.root.keys, key)
		return
	}
	if len(tree.root.keys) == 1 {
		if key < tree.root.keys[0] {
			tree.root.keys = append([]int{key}, tree.root.keys...)
		} else {
			tree.root.keys = append(tree.root.keys, key)
		}
		return
	}
	_, _, err := tree.root.insert(key)
	if err != nil {
		fmt.Println("Error:", err)
	}
}

// insert inserts a key into the subtree rooted at the current node
// Returns the child node where the key is inserted and a boolean indicating if the current node needs to split
func (node *Node) insert(key int) (*Node, bool, error) {
	if node.isLeaf {
		// Insert key into the leaf node
		return node, node.insertIntoLeaf(key), nil
	}
	// Find the child node to insert the key
	childIndex := node.findChildIndex(key)
	childNode, _, err := node.children[childIndex].insert(key)
	if err != nil {
		return nil, false, err
	}
	// Check if the child node needs to split
	if childNode != node.children[childIndex] {
		// Split occurred in the child node, update the parent node
		return node, node.insertIntoNode(childIndex, childNode.keys[0], childNode), nil
	}
	return nil, false, nil
}

// insertIntoLeaf inserts a key into the leaf node
// Returns true if the leaf node needs to split
func (node *Node) insertIntoLeaf(key int) bool {
	if len(node.keys) < 2 {
		// Insert key into the leaf node
		if key < node.keys[0] {
			node.keys = append([]int{key}, node.keys...)
		} else {
			node.keys = append(node.keys, key)
		}
		return false
	}
	// Split the leaf node
	return true
}

// insertIntoNode inserts a key into the internal node
// Returns true if the internal node needs to split
func (node *Node) insertIntoNode(childIndex int, key int, childNode *Node) bool {
	if len(node.keys) < 2 {
		// Insert key and child node into the internal node
		node.keys = append(node.keys[:childIndex], append([]int{key}, node.keys[childIndex:]...)...)
		node.children = append(node.children[:childIndex+1], append([]*Node{childNode}, node.children[childIndex+1:]...)...)
		return false
	}
	// Split the internal node
	return true
}

// findChildIndex finds the index of the child node where the key should be inserted
func (node *Node) findChildIndex(key int) int {
	for i, k := range node.keys {
		if key < k {
			return i
		}
	}
	return len(node.keys)
}

// main function to test the B+ tree
func main() {
	tree := NewBPlusTree()

	keys := []int{10, 5, 15, 3, 8, 12, 18}

	for _, key := range keys {
		tree.Insert(key)
	}

	// Print the root node keys
	fmt.Println("Root Node Keys:", tree.root.keys)
}
