package main

// Disjoint Set (Union-Find) implementation in Go using a map
type DisjointSet map[int]int

func NewDisjointSet() DisjointSet {
    return make(DisjointSet)
}

func (ds DisjointSet) Find(x int) int {
    if _, ok := ds[x]; !ok {
        ds[x] = x
    }
    if ds[x] != x {
        ds[x] = ds.Find(ds[x])
    }
    return ds[x]
}

func (ds DisjointSet) Union(x, y int) {
    xRoot := ds.Find(x)
    yRoot := ds.Find(y)
    if xRoot != yRoot {
        ds[xRoot] = yRoot
    }
}
