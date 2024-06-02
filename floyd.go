package main

import (
    "fmt"
    "math"
    "sync"
)

// Function to initialize the distance matrix
func initializeDist(graph [][]float64) [][]float64 {
    V := len(graph)
    dist := make([][]float64, V)
    for i := range dist {
        dist[i] = make([]float64, V)
        for j := range dist[i] {
            if i == j {
                dist[i][j] = 0
            } else if graph[i][j] != 0 {
                dist[i][j] = graph[i][j]
            } else {
                dist[i][j] = math.Inf(1)
            }
        }
    }
    return dist
}

// Function to update distances for a given k
func updateDistances(dist [][]float64, V int, k int, wg *sync.WaitGroup) {
    defer wg.Done()
    for i := 0; i < V; i++ {
        for j := 0; j < V; j++ {
            if dist[i][j] > dist[i][k]+dist[k][j] {
                dist[i][j] = dist[i][k] + dist[k][j]
            }
        }
    }
}

// Parallel Floyd-Warshall algorithm
func parallelFloydWarshall(graph [][]float64) [][]float64 {
    V := len(graph)
    dist := initializeDist(graph)

    var wg sync.WaitGroup

    for k := 0; k < V; k++ {
        wg.Add(1)
        go updateDistances(dist, V, k, &wg)
    }

    wg.Wait()
    return dist
}

// Helper function to print the distance matrix
func printMatrix(matrix [][]float64) {
    for _, row := range matrix {
        for _, val := range row {
            if val == math.Inf(1) {
                fmt.Print("Inf ")
            } else {
                fmt.Printf("%4.1f ", val)
            }
        }
        fmt.Println()
    }
}

func main() {
    graph := [][]float64{
        {0, 3, math.Inf(1), 5},
        {2, 0, math.Inf(1), 4},
        {math.Inf(1), 1, 0, math.Inf(1)},
        {math.Inf(1), math.Inf(1), 2, 0},
    }

    shortestPaths := parallelFloydWarshall(graph)
    fmt.Println("Shortest path matrix:")
    printMatrix(shortestPaths)
}
