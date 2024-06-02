package main

import "fmt"

//node
type Node struct{
  Key int
  Left *Node
  Right *Node
}

//insert
func (n *Node) Insert(k int){
  if n.Key < k{
    //move right
    if n.Right == nil {
      n.Right = &Node{Key: k}
    } else {
      n.Right.Insert(k)
    }
  } else if n.Key > k {
    //move left
    if n.Left == nil {
      n.Left = &Node{Key: k}
    } else {
      n.Left.Insert(k)
    }
  }
}

//search
func (n *Node) Search(k int) bool{
  if n == nil{
    return false
  }
  if n.Key < k{
    //move right
    return n.Right.Search(k)
  } else if n.Key > k {
    //move left
    return n.Left.Search(k)
  }
  return true
}


func main() {
  tree := &Node{Key: 100}   //tree needs an address
  tree.Insert(50)
  tree.Insert(200)
  tree.Insert(220)
  tree.Insert(47)
  tree.Insert(143)
  tree.Insert(101)
  tree.Insert(11)


  fmt.Println(tree)
  fmt.Println(tree.Search(200))
  fmt.Println(tree.Search(530))
  fmt.Println(tree.Search(11))
}
