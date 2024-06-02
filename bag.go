package main

// Multiset (Bag) implementation in Go using a map
type Multiset map[int]int

func NewMultiset() Multiset {
    return make(Multiset)
}

func (m Multiset) Add(val int) {
    m[val]++
}

func (m Multiset) Remove(val int) {
    if m[val] > 0 {
        m[val]--
    }
}

func (m Multiset) Count(val int) int {
    return m[val]
}
