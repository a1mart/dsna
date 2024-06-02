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

//queue FIFO
type Queue struct{
  items []int //hold a slice
}
//enqueue at end of line
func (s *Queue) Enqueue(i int) {
  s.items = append(s.items, i)
}
//dequeue from front of line
func (s *Queue) Dequeue() int{
  toRemove := s.items[0]
  s.items = s.items[1:]
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

  myQueue := Queue{}
  fmt.Println(myQueue)
  myQueue.Enqueue(100)
  myQueue.Enqueue(200)
  myQueue.Enqueue(400)
  fmt.Println(myQueue)
  myQueue.Dequeue()
  fmt.Println(myQueue)
}
