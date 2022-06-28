package htgraph

import (
	"testing"
)

func TestAddNode(t *testing.T) {
	var graph Graph[int]
	two := &Node[int]{Value: 2}
	graph = graph.NewGraph()
	graph.AddNode(two)
	if !graph.Contains(two) {
		t.Fail()
	}
}

func TestRemoveNode(t *testing.T) {
	var graph Graph[int]
	two := &Node[int]{Value: 2}
	dos := &Node[int]{Value: 2}
	graph = graph.NewGraph()
	graph.AddNode(two)
	graph.AddNode(dos)
	graph.RemoveNode(two)
	if !graph.Contains(dos) || graph.Contains(two) {
		t.Fail()
	}
}

func TestAddEdge(t *testing.T) {
	var graph Graph[int]
	one := &Node[int]{Value: 1}
	two := &Node[int]{Value: 2}
	graph = graph.NewGraph()
	graph.AddEdge(one, two)
	if !graph.HasEdge(one, two) {
		t.Fail()
	}
}

func TestRemoveEdge(t *testing.T) {
	var graph Graph[int]
	one := &Node[int]{Value: 1}
	two := &Node[int]{Value: 2}
	graph = graph.NewGraph()
	graph.AddEdge(one, two)
	graph.RemoveEdge(one, two)
	if graph.HasEdge(one, two) {
		t.Fail()
	}
}
