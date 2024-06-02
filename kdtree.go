package main

import (
	"fmt"
)

type Point struct {
	coordinates []int
}

type KDTreeNode struct {
	point      Point
	leftChild  *KDTreeNode
	rightChild *KDTreeNode
}

type KDTree struct {
	root *KDTreeNode
}

func NewKDTree(points []Point) *KDTree {
	return &KDTree{root: buildKDTree(points, 0)}
}

func buildKDTree(points []Point, depth int) *KDTreeNode {
	if len(points) == 0 {
		return nil
	}

	axis := depth % len(points[0].coordinates)
	median := len(points) / 2

	// Sort points based on the axis
	quickSelect(points, axis, median)

	return &KDTreeNode{
		point:      points[median],
		leftChild:  buildKDTree(points[:median], depth+1),
		rightChild: buildKDTree(points[median+1:], depth+1),
	}
}

func quickSelect(points []Point, axis, k int) {
	start, end := 0, len(points)-1
	for start <= end {
		pivotIndex := partition(points, axis, start, end)
		if pivotIndex == k {
			return
		} else if pivotIndex < k {
			start = pivotIndex + 1
		} else {
			end = pivotIndex - 1
		}
	}
}

func partition(points []Point, axis, start, end int) int {
	pivot := points[end]
	pIndex := start
	for i := start; i < end; i++ {
		if points[i].coordinates[axis] < pivot.coordinates[axis] {
			points[i], points[pIndex] = points[pIndex], points[i]
			pIndex++
		}
	}
	points[end], points[pIndex] = points[pIndex], points[end]
	return pIndex
}

func main() {
	points := []Point{
		{[]int{3, 6}},
		{[]int{17, 15}},
		{[]int{13, 15}},
		{[]int{6, 12}},
		{[]int{9, 1}},
		{[]int{2, 7}},
		{[]int{10, 19}},
	}
	kdTree := NewKDTree(points)
	fmt.Println("KD-Tree:", kdTree)
}
