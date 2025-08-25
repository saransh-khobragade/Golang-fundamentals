// Graph implementation in Go
// Supports addVertex, addEdge, removeEdge, removeVertex, depthFirstSearch, breadthFirstSearch, and utility methods

package datastructures

import "fmt"

type Graph struct {
	adjacencyList map[string][]string
}

func NewGraph() *Graph {
	return &Graph{
		adjacencyList: make(map[string][]string),
	}
}

func (g *Graph) AddVertex(vertex string) {
	if _, exists := g.adjacencyList[vertex]; !exists {
		g.adjacencyList[vertex] = make([]string, 0)
	}
}

func (g *Graph) AddEdge(vertex1, vertex2 string) {
	if _, exists := g.adjacencyList[vertex1]; !exists {
		g.AddVertex(vertex1)
	}
	if _, exists := g.adjacencyList[vertex2]; !exists {
		g.AddVertex(vertex2)
	}
	g.adjacencyList[vertex1] = append(g.adjacencyList[vertex1], vertex2)
	g.adjacencyList[vertex2] = append(g.adjacencyList[vertex2], vertex1)
}

func (g *Graph) RemoveEdge(vertex1, vertex2 string) {
	if neighbors, exists := g.adjacencyList[vertex1]; exists {
		for i, neighbor := range neighbors {
			if neighbor == vertex2 {
				g.adjacencyList[vertex1] = append(neighbors[:i], neighbors[i+1:]...)
				break
			}
		}
	}
	if neighbors, exists := g.adjacencyList[vertex2]; exists {
		for i, neighbor := range neighbors {
			if neighbor == vertex1 {
				g.adjacencyList[vertex2] = append(neighbors[:i], neighbors[i+1:]...)
				break
			}
		}
	}
}

func (g *Graph) RemoveVertex(vertex string) {
	if neighbors, exists := g.adjacencyList[vertex]; exists {
		for len(neighbors) > 0 {
			adjacentVertex := neighbors[0]
			g.RemoveEdge(vertex, adjacentVertex)
			neighbors = g.adjacencyList[vertex]
		}
		delete(g.adjacencyList, vertex)
	}
}

func (g *Graph) DepthFirstSearch(start string) []string {
	result := make([]string, 0)
	visited := make(map[string]bool)
	g.dfsHelper(start, visited, &result)
	return result
}

func (g *Graph) dfsHelper(vertex string, visited map[string]bool, result *[]string) {
	if vertex == "" || visited[vertex] {
		return
	}

	visited[vertex] = true
	*result = append(*result, vertex)

	for _, neighbor := range g.adjacencyList[vertex] {
		if !visited[neighbor] {
			g.dfsHelper(neighbor, visited, result)
		}
	}
}

func (g *Graph) BreadthFirstSearch(start string) []string {
	result := make([]string, 0)
	visited := make(map[string]bool)
	queue := []string{start}
	visited[start] = true

	for len(queue) > 0 {
		currentVertex := queue[0]
		queue = queue[1:]
		result = append(result, currentVertex)

		for _, neighbor := range g.adjacencyList[currentVertex] {
			if !visited[neighbor] {
				visited[neighbor] = true
				queue = append(queue, neighbor)
			}
		}
	}
	return result
}

func (g *Graph) IsEmpty() bool {
	return len(g.adjacencyList) == 0
}

func (g *Graph) Size() int {
	return len(g.adjacencyList)
}

func (g *Graph) GetVertices() []string {
	vertices := make([]string, 0, len(g.adjacencyList))
	for vertex := range g.adjacencyList {
		vertices = append(vertices, vertex)
	}
	return vertices
}

func (g *Graph) GetNeighbors(vertex string) []string {
	if neighbors, exists := g.adjacencyList[vertex]; exists {
		result := make([]string, len(neighbors))
		copy(result, neighbors)
		return result
	}
	return make([]string, 0)
}

func (g *Graph) HasVertex(vertex string) bool {
	_, exists := g.adjacencyList[vertex]
	return exists
}

func (g *Graph) HasEdge(vertex1, vertex2 string) bool {
	if neighbors, exists := g.adjacencyList[vertex1]; exists {
		for _, neighbor := range neighbors {
			if neighbor == vertex2 {
				return true
			}
		}
	}
	return false
}

func (g *Graph) Clear() {
	g.adjacencyList = make(map[string][]string)
}

func (g *Graph) Display() {
	fmt.Printf("Graph: %v\n", g.adjacencyList)
} 