package main

// Deque (Double-ended Queue) implementation in Go using a slice
type Deque []int

func NewDeque() Deque {
    return make(Deque, 0)
}

func (d *Deque) PushFront(val int) {
    *d = append([]int{val}, *d...)
}

func (d *Deque) PushBack(val int) {
    *d = append(*d, val)
}

func (d *Deque) PopFront() int {
    if len(*d) == 0 {
        return -1 // Placeholder for empty deque
    }
    val := (*d)[0]
    *d = (*d)[1:]
    return val
}

func (d *Deque) PopBack() int {
    if len(*d) == 0 {
        return -1 // Placeholder for empty deque
    }
    val := (*d)[len(*d)-1]
    *d = (*d)[:len(*d)-1]
    return val
}
