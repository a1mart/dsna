package main

import (
	"fmt"
)

// Graph represents the graph structure.
type Graph struct {
	nodes map[string]map[string]int
}

// AddNode adds a node to the graph.
func (g *Graph) AddNode(node string) {
	if g.nodes == nil {
		g.nodes = make(map[string]map[string]int)
	}
	g.nodes[node] = make(map[string]int)
}

// AddEdge adds an edge between two nodes with a given weight.
func (g *Graph) AddEdge(node1, node2 string, weight int) {
	if _, ok := g.nodes[node1]; !ok {
		g.AddNode(node1)
	}
	if _, ok := g.nodes[node2]; !ok {
		g.AddNode(node2)
	}
	g.nodes[node1][node2] = weight
	g.nodes[node2][node1] = weight // for undirected graph
}

// Dijkstra finds the shortest paths from a source node to all other nodes in the graph.
func (g *Graph) Dijkstra(source string) map[string]int {
	distances := make(map[string]int)
	for node := range g.nodes {
		distances[node] = -1 // set initial distances to -1 (infinite)
	}

	distances[source] = 0

	visited := make(map[string]bool)
	for len(visited) < len(g.nodes) {
		u := minDistance(distances, visited)
		visited[u] = true

		for v, weight := range g.nodes[u] {
			if !visited[v] && (distances[v] == -1 || distances[u]+weight < distances[v]) {
				distances[v] = distances[u] + weight
			}
		}
	}
	return distances
}

// minDistance finds the node with the minimum distance value, from the set of nodes not yet visited.
func minDistance(distances map[string]int, visited map[string]bool) string {
	min := -1
	var minNode string
	for node, dist := range distances {
		if !visited[node] && (min == -1 || dist < min) {
			min = dist
			minNode = node
		}
	}
	return minNode
}

func main() {
	graph := Graph{}

	graph.AddEdge("A", "B", 4)
	graph.AddEdge("A", "C", 2)
	graph.AddEdge("B", "C", 5)
	graph.AddEdge("B", "D", 10)
	graph.AddEdge("C", "D", 3)

	distances := graph.Dijkstra("A")

	fmt.Println("Shortest distances from node A:")
	for node, dist := range distances {
		fmt.Printf("Node %s: %d\n", node, dist)
	}
}
