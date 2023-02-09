package main

import "fmt"

type Graph struct {
	table map[int][]int
}

// type graph struct{
// 	starting *graphNode
// }

func main() {
	graph := &Graph{table: make(map[int][]int)}

	graph.insert(3, 5, true)
	graph.insert(3, 4, true)
	graph.insert(5, 6, false)

	graph.display()

}

func (g *Graph) display() {
	for vertex, edges := range g.table {
		fmt.Println(vertex, "  : ", edges)
	}
}

func (g *Graph) insert(vertex, edge int, bidirectional bool) {
	if !g.exist(vertex) {
		g.addVertex(vertex)
	}
	if !g.exist(edge) {
		g.addVertex(edge)
	}

	g.table[vertex] = append(g.table[vertex], edge)

	if bidirectional {
		g.table[edge] = append(g.table[edge], vertex)
	}

}

func (g *Graph) addVertex(vertex int) {
	// edges := new([]int)
	g.table[vertex] = []int{}
}

func (g *Graph) exist(vertex int) bool {
	_, exist := g.table[vertex]
	return exist
}
