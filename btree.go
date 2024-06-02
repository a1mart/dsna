package main

import "fmt"

const t = 2 // Minimum degree

type BTreeNode struct {
    keys     []int
    children []*BTreeNode
    leaf     bool
}

func NewBTreeNode(leaf bool) *BTreeNode {
    return &BTreeNode{
        keys:     make([]int, 0),
        children: make([]*BTreeNode, 0),
        leaf:     leaf,
    }
}

func (node *BTreeNode) traverse() {
    var i int
    for i = 0; i < len(node.keys); i++ {
        if !node.leaf {
            node.children[i].traverse()
        }
        fmt.Printf("%d ", node.keys[i])
    }
    if !node.leaf {
        node.children[i].traverse()
    }
}

func (node *BTreeNode) search(key int) bool {
    i := 0
    for i < len(node.keys) && key > node.keys[i] {
        i++
    }
    if i < len(node.keys) && node.keys[i] == key {
        return true
    }
    if node.leaf {
        return false
    }
    return node.children[i].search(key)
}

func (node *BTreeNode) insertNonFull(key int) {
    i := len(node.keys) - 1
    if node.leaf {
        for i >= 0 && key < node.keys[i] {
            i--
        }
        node.keys = append(node.keys, 0)
        copy(node.keys[i+2:], node.keys[i+1:])
        node.keys[i+1] = key
    } else {
        for i >= 0 && key < node.keys[i] {
            i--
        }
        i++
        if len(node.children[i].keys) == 2*t-1 {
            node.splitChild(i, node.children[i])
            if key > node.keys[i] {
                i++
            }
        }
        node.children[i].insertNonFull(key)
    }
}

func (node *BTreeNode) splitChild(i int, y *BTreeNode) {
    z := NewBTreeNode(y.leaf)
    z.keys = append(z.keys, y.keys[t:]...)
    y.keys = y.keys[:t]

    if !y.leaf {
        z.children = append(z.children, y.children[t:]...)
        y.children = y.children[:t]
    }

    node.children = append(node.children, nil)
    copy(node.children[i+2:], node.children[i+1:])
    node.children[i+1] = z

    node.keys = append(node.keys, 0)
    copy(node.keys[i+1:], node.keys[i:])
    node.keys[i] = y.keys[t-1]
}

type BTree struct {
    root *BTreeNode
}

func NewBTree() *BTree {
    return &BTree{
        root: NewBTreeNode(true),
    }
}

func (tree *BTree) traverse() {
    if tree.root != nil {
        tree.root.traverse()
    }
}

func (tree *BTree) search(key int) bool {
    if tree.root == nil {
        return false
    }
    return tree.root.search(key)
}

func (tree *BTree) insert(key int) {
    root := tree.root
    if len(root.keys) == 2*t-1 {
        s := NewBTreeNode(false)
        s.children = append(s.children, root)
        s.splitChild(0, root)
        tree.root = s
        s.insertNonFull(key)
    } else {
        root.insertNonFull(key)
    }
}

func main() {
    tree := NewBTree()

    keys := []int{10, 20, 5, 6, 12, 30, 7, 17}

    for _, key := range keys {
        tree.insert(key)
    }

    fmt.Println("Traversal of the constructed B-tree is:")
    tree.traverse()

    searchKey := 6
    if tree.search(searchKey) {
        fmt.Printf("\nKey %d found in the B-tree", searchKey)
    } else {
        fmt.Printf("\nKey %d not found in the B-tree", searchKey)
    }
}
