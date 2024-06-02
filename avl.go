package main

import (
    "fmt"
)

type Node struct {
    key    int
    left   *Node
    right  *Node
    height int
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}

func height(node *Node) int {
    if node == nil {
        return -1
    }
    return node.height
}

func updateHeight(node *Node) {
    node.height = 1 + max(height(node.left), height(node.right))
}

func rotateRight(y *Node) *Node {
    x := y.left
    T2 := x.right

    x.right = y
    y.left = T2

    updateHeight(y)
    updateHeight(x)

    return x
}

func rotateLeft(x *Node) *Node {
    y := x.right
    T2 := y.left

    y.left = x
    x.right = T2

    updateHeight(x)
    updateHeight(y)

    return y
}

func getBalance(node *Node) int {
    if node == nil {
        return 0
    }
    return height(node.left) - height(node.right)
}

func insert(root *Node, key int) *Node {
    if root == nil {
        return &Node{key: key, height: 0}
    }

    if key < root.key {
        root.left = insert(root.left, key)
    } else if key > root.key {
        root.right = insert(root.right, key)
    } else {
        // Duplicate keys are not allowed in AVL trees
        return root
    }

    updateHeight(root)

    balance := getBalance(root)

    // Left Left Case
    if balance > 1 && key < root.left.key {
        return rotateRight(root)
    }

    // Right Right Case
    if balance < -1 && key > root.right.key {
        return rotateLeft(root)
    }

    // Left Right Case
    if balance > 1 && key > root.left.key {
        root.left = rotateLeft(root.left)
        return rotateRight(root)
    }

    // Right Left Case
    if balance < -1 && key < root.right.key {
        root.right = rotateRight(root.right)
        return rotateLeft(root)
    }

    return root
}

func preOrder(root *Node) {
    if root != nil {
        fmt.Printf("%d ", root.key)
        preOrder(root.left)
        preOrder(root.right)
    }
}
/*
func main() {
    var root *Node
    keys := []int{10, 20, 30, 40, 50, 25}

    for _, key := range keys {
        root = insert(root, key)
    }

    fmt.Println("Preorder traversal of constructed AVL tree is:")
    preOrder(root)
}
*/
