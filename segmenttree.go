package main

// Segment Tree implementation in Go
type SegmentTree struct {
    tree   []int
    nums   []int
}

func NewSegmentTree(nums []int) *SegmentTree {
    n := len(nums)
    tree := make([]int, 4*n) // Size of the segment tree
    return &SegmentTree{tree: tree, nums: nums}
}

func (st *SegmentTree) buildTree(node, start, end int) {
    if start == end {
        st.tree[node] = st.nums[start]
    } else {
        mid := (start + end) / 2
        leftChild := 2*node + 1
        rightChild := 2*node + 2
        st.buildTree(leftChild, start, mid)
        st.buildTree(rightChild, mid+1, end)
        st.tree[node] = st.tree[leftChild] + st.tree[rightChild] // Example: sum of range
    }
}

func (st *SegmentTree) Query(node, start, end, left, right int) int {
    if right < start || left > end {
        return 0 // Example: for sum query, return 0
    }
    if left <= start && right >= end {
        return st.tree[node]
    }
    mid := (start + end) / 2
    leftChild := 2*node + 1
    rightChild := 2*node + 2
    leftSum := st.Query(leftChild, start, mid, left, right)
    rightSum := st.Query(rightChild, mid+1, end, left, right)
    return leftSum + rightSum // Example: sum of left and right children
}

func (st *SegmentTree) Update(node, start, end, index, value int) {
    if start == end {
        st.nums[index] = value
        st.tree[node] = value
    } else {
        mid := (start + end) / 2
        leftChild := 2*node + 1
        rightChild := 2*node + 2
        if index >= start && index <= mid {
            st.Update(leftChild, start, mid, index, value)
        } else {
            st.Update(rightChild, mid+1, end, index, value)
        }
        st.tree[node] = st.tree[leftChild] + st.tree[rightChild] // Example: update sum
    }
}
