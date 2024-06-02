package main

// Ternary Search Tree implementation in Go
type TSTNode struct {
    val         byte
    left, mid, right *TSTNode
    isEndOfWord bool
}

type TernarySearchTree struct {
    root *TSTNode
}

func NewTernarySearchTree() *TernarySearchTree {
    return &TernarySearchTree{}
}

func (tst *TernarySearchTree) Insert(word string) {
    if len(word) == 0 {
        return
    }
    tst.root = tst.insertHelper(tst.root, word, 0)
}

func (tst *TernarySearchTree) insertHelper(node *TSTNode, word string, index int) *TSTNode {
    if node == nil {
        node = &TSTNode{val: word[index]}
    }
    if word[index] < node.val {
        node.left = tst.insertHelper(node.left, word, index)
    } else if word[index] > node.val {
        node.right = tst.insertHelper(node.right, word, index)
    } else {
        if index < len(word)-1 {
            node.mid = tst.insertHelper(node.mid, word, index+1)
        } else {
            node.isEndOfWord = true
        }
    }
    return node
}

func (tst *TernarySearchTree) Search(word string) bool {
    node := tst.searchHelper(tst.root, word, 0)
    return node != nil && node.isEndOfWord
}

func (tst *TernarySearchTree) searchHelper(node *TSTNode, word string, index int) *TSTNode {
    if node == nil {
        return nil
    }
    if word[index] < node.val {
        return tst.searchHelper(node.left, word, index)
    } else if word[index] > node.val {
        return tst.searchHelper(node.right, word, index)
    } else {
        if index == len(word)-1 {
            return node
        }
        return tst.searchHelper(node.mid, word, index+1)
    }
}
