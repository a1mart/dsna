package main

import "fmt"

// MaxHeap struct has a slice that holds the array
type MaxHeap struct{
  array []int
}

//insert adds element to the heap
func (h *MaxHeap) Insert(key int){
  h.array = append(h.array, key)
  h.maxHeapifyUp(len(h.array)-1)
}
//extract returns the largest key and removes it from the heap
func (h *MaxHeap) Extract() int{
  extracted := h.array[0]
  l := len(h.array)-1
  if len(h.array) == 0{
    fmt.Println("cannot extract because array length is 0")
    return -1
  }
  //take out last index and put at root
  h.array[0] = h.array[len(h.array)-1]
  h.array = h.array[:l]

  h.maxHeapifyDown(0)
  return extracted
}

//maxHeapifyUp will heapify from bottom up
func (h *MaxHeap) maxHeapifyUp(index int){
  for h.array[parent(index)] < h.array[index]{
    h.swap(parent(index),index)
    index = parent(index)
  }
}
//maxHeapifyUp will heapify from top to bottom
func (h *MaxHeap) maxHeapifyDown(index int){
  lastIndex := len(h.array) -1
  l,r := left(index), right(index)
  child2compare := 0
  // loop while index has at least one child
  for l <= lastIndex{
    if l == lastIndex {   //when left child is the only child
      child2compare = l
    } else if h.array[l] > h.array[r] { //when left child is larger
      child2compare = l
    }else { //right is larger
      child2compare = r
    }
    //compare array value of current index to larger child and swap if smaller
    if h.array[index] < h.array[child2compare]{
      h.swap(index, child2compare)
      index = child2compare
      l, r = left(index), right(index)
    } else {
      return
    }
  }
}

func parent(i int) int{
  return (i-1)/2
}
//left child index
func left(i int) int{
  return 2*i + 1
}
//right child index
func right(i int) int{
  return 2*i + 2
}

// swap keys in the array
func (h *MaxHeap) swap(i1,i2 int) {
  h.array[i1], h.array[i2] = h.array[i2], h.array[i1]
}

func main() {
  m := &MaxHeap{}
  fmt.Println(m)
  buildHeap := []int{10,20,30, 5, 23, 32, 11, 21, 18}
  for _,v := range buildHeap {
    m.Insert(v)
    fmt.Println(m)
  }

  k := 5
  for i := 0; i < k; i++ {
    m.Extract()
    fmt.Println(m)
  }
}
