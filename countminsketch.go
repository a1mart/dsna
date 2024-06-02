package main

import (
    "hash/fnv"
    "math"
)

type CountMinSketch struct {
    table [][]int
    width int
    depth int
}

func NewCountMinSketch(width, depth int) *CountMinSketch {
    table := make([][]int, depth)
    for i := range table {
        table[i] = make([]int, width)
    }
    return &CountMinSketch{
        table: table,
        width: width,
        depth: depth,
    }
}

func (cms *CountMinSketch) hash(item string) int {
    hash := fnv.New64a()
    hash.Write([]byte(item))
    return int(hash.Sum64()) & (cms.width - 1) // Ensure non-negativity and keep within range
}

func (cms *CountMinSketch) Add(item string) {
    for i := 0; i < cms.depth; i++ {
        index := cms.hash(item) % cms.width
        cms.table[i][index]++
    }
}

func (cms *CountMinSketch) EstimateFrequency(item string) int {
    minCount := math.MaxInt32
    for i := 0; i < cms.depth; i++ {
        index := cms.hash(item) % cms.width
        if cms.table[i][index] < minCount {
            minCount = cms.table[i][index]
        }
    }
    return minCount
}

func main() {
    width := 1000
    depth := 5
    cms := NewCountMinSketch(width, depth)

    items := []string{"apple", "banana", "apple", "orange", "apple", "banana", "banana"}
    for _, item := range items {
        cms.Add(item)
    }

    frequencies := map[string]int{}
    for _, item := range items {
        frequencies[item] = cms.EstimateFrequency(item)
    }

    for item, count := range frequencies {
        println("Estimated frequency of", item, ":", count)
    }
}
