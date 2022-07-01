package htgraph

import "sync"

//Implemented as an adjacency list
type Graph[T comparable] struct {
	nodes map[*Node[T]]([]Edge[T])
	mu    sync.Mutex
}

type Edge[T any] struct {
	target *Node[T]
	source *Node[T]
	weight float64
}

type Node[T any] struct {
	Value T
}

func (g *Graph[T]) NewGraph() Graph[T] {
	return Graph[T]{
		nodes: make(map[*Node[T]]([]Edge[T])),
	}
}

func (g *Graph[T]) AddNode(n *Node[T]) {
	g.mu.Lock()
	g.nodes[n] = []Edge[T]{}
	g.mu.Unlock()
}

func (g *Graph[T]) RemoveNode(n *Node[T]) {
	g.mu.Lock()
	delete(g.nodes, n)
	g.mu.Unlock()
}

func (g *Graph[T]) AddEdge(source *Node[T], target *Node[T]) {
	g.AddWeightedEdge(source, target, 0)
}

func (g *Graph[T]) AddWeightedEdge(source *Node[T], target *Node[T], weight float64) {
	if !g.Contains(source) {
		g.AddNode(source)
	}

	if !g.Contains(target) {
		g.AddNode(target)
	}

	edge := Edge[T]{source: source, target: target, weight: weight}

	g.mu.Lock()
	g.nodes[source] = append(g.nodes[source], edge)
	g.mu.Unlock()
}

func (g *Graph[T]) RemoveEdge(source *Node[T], target *Node[T]) {
	g.mu.Lock()
	defer g.mu.Unlock()
	edges := g.nodes[source]
	for i, node := range edges {
		if node.target == target {
			edges[i] = edges[len(edges)-1]
			g.nodes[source] = edges[:len(edges)-1]
		}
	}
}

func (g *Graph[T]) HasEdge(source *Node[T], target *Node[T]) bool {
	g.mu.Lock()
	defer g.mu.Unlock()
	edges := g.nodes[source]
	for _, edge := range edges {
		if edge.target == target {
			return true
		}
	}
	return false
}

func (g *Graph[T]) Contains(n *Node[T]) bool {
	g.mu.Lock()
	_, ok := g.nodes[n]
	g.mu.Unlock()
	return ok
}
