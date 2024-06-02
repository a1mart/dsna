package main

import "fmt"

//stack FILO
type Stack struct{
  items []int //hold a slice
}
//push onto top
func (s *Stack) Push(i int) {
  s.items = append(s.items, i)
}
//pop from top
func (s *Stack) Pop() int{
  l := len(s.items)-1
  toRemove := s.items[l]
  s.items = s.items[:l]
  return toRemove
}

func main(){
  myStack := Stack{}
  fmt.Println(myStack)
  myStack.Push(100)
  myStack.Push(200)
  myStack.Push(400)
  fmt.Println(myStack)
  myStack.Pop()
  fmt.Println(myStack)
}
