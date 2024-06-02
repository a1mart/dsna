package main

import (
  "fmt"
  "strings"
)

const (
  AlphabetSize = 26
  FirstCharacter = 'a'
)

// Node
type Node struct{
  children [AlphabetSize]*Node
  isEnd bool
}

// Trie
type Trie struct{
  root *Node
}

func Init() *Trie {
  return &Trie{root: &Node{}}
}

// Insert will take a word and add it to the trie
func (t *Trie) Insert(w string) {
  wordLength := len(w)
  currentNode := t.root
  for i := 0; i < wordLength; i++{
    charIndex := w[i] - FirstCharacter
    if currentNode.children[charIndex] == nil {
      currentNode.children[charIndex] = &Node{}
    }
    currentNode = currentNode.children[charIndex]
  }
  currentNode.isEnd = true
}

// Search will take a word and return true if it exists in the trie
func (t *Trie) Search(w string) bool {
  wordLength := len(w)
  currentNode := t.root
  for i := 0; i < wordLength; i++{
    charIndex := w[i] - FirstCharacter
    if currentNode.children[charIndex] == nil {
      return false
    }
    currentNode = currentNode.children[charIndex]
  }
  if currentNode.isEnd == true {
    return true
  }
  return false
}

func main(){
  tri := Init()
  fmt.Println(tri)
  fmt.Println(tri.root)

  tri.Insert("aragon")
  tri.Insert("oregano")
  tri.Insert("oregan")

  fmt.Println(tri.Search("oregano"))
  fmt.Println(tri.Search("oreganoa"))
  //consider case sensitivity of alphabet... using runes for different languages
  fmt.Println(tri.Search(strings.ToLower("Aragon")))
}
