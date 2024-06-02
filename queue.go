package main

import "fmt"

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
  myQueue := Queue{}
  fmt.Println(myQueue)
  myQueue.Enqueue(100)
  myQueue.Enqueue(200)
  myQueue.Enqueue(400)
  fmt.Println(myQueue)
  myQueue.Dequeue()
  fmt.Println(myQueue)
}
