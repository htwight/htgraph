package htgraph

//Implemented as an adjacency list
type Graph[T comparable] struct {
	nodes map[*Node[T]]([]*Node[T])
}

type Node[T any] struct {
	Value  T
	Weight float64
}

func (g *Graph[T]) NewGraph() Graph[T] {
	return Graph[T]{
		nodes: make(map[*Node[T]]([]*Node[T])),
	}
}

func (g *Graph[T]) AddNode(n *Node[T]) {
	g.nodes[n] = []*Node[T]{}
}

func (g *Graph[T]) RemoveNode(n *Node[T]) {
	delete(g.nodes, n)
}

func (g *Graph[T]) AddEdge(source *Node[T], target *Node[T]) {
	if !g.Contains(source) {
		g.AddNode(source)
	}

	if !g.Contains(target) {
		g.AddNode(target)
	}

	g.nodes[source] = append(g.nodes[source], target)
}

func (g *Graph[T]) RemoveEdge(source *Node[T], target *Node[T]) {
	edges := g.nodes[source]
	for i, node := range edges {
		if node == target {
			edges[i] = edges[len(edges)-1]
			g.nodes[source] = edges[:len(edges)-1]
		}
	}
}

func (g *Graph[T]) HasEdge(source *Node[T], target *Node[T]) bool {
	edges := g.nodes[source]
	for _, edge := range edges {
		if edge == target {
			return true
		}
	}
	return false
}

func (g *Graph[T]) Contains(n *Node[T]) bool {
	_, ok := g.nodes[n]
	return ok
}
