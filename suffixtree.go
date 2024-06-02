package main

import (
	"fmt"
)

type SuffixTreeNode struct {
	children map[rune]*SuffixTreeNode
	start    int
	end      *int // nil if leaf node
}

type SuffixTree struct {
	root *SuffixTreeNode
	text string
}

func NewSuffixTree(text string) *SuffixTree {
	text += "$" // Add sentinel character
	root := &SuffixTreeNode{children: make(map[rune]*SuffixTreeNode)}
	for i := 0; i < len(text); i++ {
		addSuffix(root, text, i)
	}
	return &SuffixTree{root: root, text: text}
}

func addSuffix(node *SuffixTreeNode, text string, i int) {
	current := node
	for j := i; j < len(text); j++ {
		r := rune(text[j])
		if current.children == nil {
			current.children = make(map[rune]*SuffixTreeNode)
		}
		if _, ok := current.children[r]; !ok {
			end := &i // Leaf node's end is the start of the suffix
			current.children[r] = &SuffixTreeNode{make(map[rune]*SuffixTreeNode), i, end}
			return
		}
		current = current.children[r]
	}
}

func (st *SuffixTree) Display() {
	fmt.Println("Suffix Tree:")
	displayHelper(st.root, 0, st.text)
}

func displayHelper(node *SuffixTreeNode, depth int, text string) {
	if node == nil {
		return
	}
	if node.end != nil {
		fmt.Printf("%s\n", text[node.start:])
		return
	}
	for r, child := range node.children {
		fmt.Printf("%s%v\n", string(r), depth)
		displayHelper(child, depth+1, text)
	}
}
/*
func main() {
	text := "banana"
	suffixTree := NewSuffixTree(text)
	suffixTree.Display()
}
*/
