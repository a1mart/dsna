package main

// Fenwick Tree (Binary Indexed Tree) implementation in Go
type FenwickTree []int

func NewFenwickTree(size int) FenwickTree {
    return make(FenwickTree, size+1)
}

func (ft FenwickTree) Update(index, delta int) {
    for i := index + 1; i < len(ft); i += i & -i {
        ft[i] += delta
    }
}

func (ft FenwickTree) Query(index int) int {
    sum := 0
    for i := index + 1; i > 0; i -= i & -i {
        sum += ft[i]
    }
    return sum
}
