package main

import "fmt"

// Graph represents an adjacency list graph
type Graph struct{
  verticies []*Vertex
}
// Vertex represents a graph vertex
type Vertex struct{
  key int
  adjacent []*Vertex
}

// Add Vertex
func (g *Graph) AddVertex(k int){
  if contains(g.verticies, k){
    err := fmt.Errorf("Vertex %v already exists", k)
    fmt.Println(err.Error())
  } else{
    g.verticies = append(g.verticies, &Vertex{key: k})
  }
}

// Add Edge
func (g *Graph) AddEdge(from, to int) {
  // get vertex
  fromVertex := g.getVertex(from)
  toVertex := g.getVertex(to)
  // check error
  if fromVertex == nil || toVertex == nil {
    err := fmt.Errorf("Invalid Edge (%v-->%v)", from, to)
    fmt.Println(err.Error())
  } else if contains(fromVertex.adjacent, to) {
    err := fmt.Errorf("Existing Edge (%v-->%v)", from, to)
    fmt.Println(err.Error())
  } else {
    // add edge
    fromVertex.adjacent = append(fromVertex.adjacent, toVertex)
  }
}

//getVertex
func (g *Graph) getVertex(k int) *Vertex{
  for i,v := range g.verticies {
    if v.key == k {
      return g.verticies[i]
    }
  }
  return nil
}
//contains
func contains(s []*Vertex, k int) bool {
  for _,v := range s {
    if k == v.key {
      return true
    }
  }
  return false
}

// Print will print out the adjacent list for each vertex of the graph
func (g *Graph) Print() {
  for _,v := range g.verticies {
    fmt.Printf("\nVertex %v: ", v.key)
    for _,v := range v.adjacent {
      fmt.Printf("%v ", v.key)
    }
  }
  fmt.Println()
}


func main(){
  a := &Graph{}
  for i:=0; i<5; i++{
    a.AddVertex(i)
  }
  a.AddVertex(0)
  a.AddEdge(0,1)
  a.AddEdge(0,2)

  a.AddEdge(1,2)
  a.AddEdge(1,2)
  a.AddEdge(6,2)


  a.Print()
}
