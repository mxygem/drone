package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHandleNodes(t *testing.T) {
	testCases := []struct {
		name     string
		nodes    []*Node
		expected *Graph
	}{
		{
			name: "test unsorted node collection",
			nodes: []*Node{
				{X: 7, Y: 5, R: 2},
				{X: 3, Y: 3, R: 2},
				{X: 10, Y: 10, R: 3},
				{X: 5, Y: 3, R: 2},
				{X: 7, Y: 8, R: 2},
				{X: 0, Y: 0, R: 3},
			},
			expected: &Graph{
				Nodes: []*Node{
					{X: 0, Y: 0, R: 3},
					{X: 3, Y: 3, R: 2},
					{X: 5, Y: 3, R: 2},
					{X: 7, Y: 5, R: 2},
					{X: 7, Y: 8, R: 2},
					{X: 10, Y: 10, R: 3},
				},
				Edges: []*Edge{
					{N1: &Node{X: 0, Y: 0, R: 3}, N2: &Node{X: 3, Y: 3, R: 2}},
					{N1: &Node{X: 3, Y: 3, R: 2}, N2: &Node{X: 5, Y: 3, R: 2}},
					{N1: &Node{X: 5, Y: 3, R: 2}, N2: &Node{X: 7, Y: 5, R: 2}},
					{N1: &Node{X: 7, Y: 5, R: 2}, N2: &Node{X: 7, Y: 8, R: 2}},
					{N1: &Node{X: 7, Y: 8, R: 2}, N2: &Node{X: 10, Y: 10, R: 3}},
				},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			graph := &Graph{}

			handleNodes(graph, tc.nodes)

			assert.ElementsMatch(t, tc.expected.Nodes, graph.Nodes, "unexpected nodes found")
			assert.ElementsMatch(t, tc.expected.Edges, graph.Edges, "unexpected edges found")
		})
	}
}

func TestIsEdge(t *testing.T) {
	testCases := []struct {
		name     string
		node1    *Node
		node2    *Node
		expected bool
	}{
		{
			name:     "nodes must have different origin points to be evaluated",
			node1:    &Node{X: 5, Y: 5, R: 2},
			node2:    &Node{X: 5, Y: 5, R: 5},
			expected: false,
		},
		{
			name:     "not edges",
			node1:    &Node{X: 5, Y: 0, R: 1},
			node2:    &Node{X: 5, Y: 10, R: 1},
			expected: false,
		},
		{
			name:     "edges same radii",
			node1:    &Node{X: 5, Y: 4, R: 5},
			node2:    &Node{X: 5, Y: 6, R: 5},
			expected: true,
		},
		{
			name:     "edges different radii",
			node1:    &Node{X: 10, Y: 0, R: 25},
			node2:    &Node{X: 0, Y: 10, R: 1},
			expected: true,
		},
		{
			name:     "various 0",
			node1:    &Node{X: 0, Y: 0, R: 3},
			node2:    &Node{X: 3, Y: 3, R: 2},
			expected: true,
		},
		{
			name:     "various 1",
			node1:    &Node{X: 3, Y: 3, R: 2},
			node2:    &Node{X: 5, Y: 3, R: 2},
			expected: true,
		},
		{
			name:     "various 2",
			node1:    &Node{X: 5, Y: 3, R: 2},
			node2:    &Node{X: 7, Y: 5, R: 3},
			expected: true,
		},
		{
			name:     "various 3",
			node1:    &Node{X: 7, Y: 5, R: 3},
			node2:    &Node{X: 7, Y: 8, R: 3},
			expected: true,
		},
		{
			name:     "various 4",
			node1:    &Node{X: 7, Y: 8, R: 3},
			node2:    &Node{X: 10, Y: 10, R: 3},
			expected: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.expected, isEdge(tc.node1, tc.node2))
		})
	}
}

func TestOriginConnectedEdges(t *testing.T) {
	testCases := []struct {
		name     string
		graph    *Graph
		point1   *Point
		point2   *Point
		expected bool
	}{
		{
			name:     "no graph provided",
			expected: false,
		},
		{
			name:     "no point1 provided",
			graph:    &Graph{},
			expected: false,
		},
		{
			name:     "no point2 provided",
			graph:    &Graph{},
			point1:   &Point{X: 2, Y: 3},
			expected: false,
		},
		{
			name:     "no nodes in graph",
			graph:    &Graph{},
			point1:   &Point{X: 2, Y: 3},
			point2:   &Point{X: 6, Y: 3},
			expected: false,
		},
		{
			name: "at least two nodes required",
			graph: &Graph{
				Nodes: []*Node{
					{X: 0, Y: 10},
				},
			},
			point1:   &Point{X: 2, Y: 3},
			point2:   &Point{X: 6, Y: 3},
			expected: false,
		},
		{
			name: "no edges in graph",
			graph: &Graph{
				Nodes: []*Node{
					{X: 0, Y: 10},
					{X: 10, Y: 0},
				},
			},
			point1:   &Point{X: 2, Y: 3},
			point2:   &Point{X: 6, Y: 3},
			expected: false,
		},
		{
			name:   "no nodes match points",
			point1: &Point{X: 2, Y: 3},
			point2: &Point{X: 6, Y: 3},
			graph: &Graph{
				Nodes: []*Node{
					{X: 0, Y: 0, R: 3},
					{X: 0, Y: 5, R: 3},
				},
				Edges: []*Edge{{N1: &Node{X: 0, Y: 0, R: 3}, N2: &Node{X: 0, Y: 0, R: 3}}},
			},
			expected: false,
		},
		{
			name: "first two nodes are connected but end point is not",
			graph: &Graph{
				Nodes: []*Node{
					{X: 0, Y: 0, R: 2},
					{X: 0, Y: 3, R: 2},
					{X: 0, Y: 10, R: 2},
				},
				Edges: []*Edge{
					{N1: &Node{X: 0, Y: 0, R: 2}, N2: &Node{X: 0, Y: 3, R: 2}},
				},
			},
			point1:   &Point{X: 0, Y: 0},
			point2:   &Point{X: 0, Y: 10},
			expected: false,
		},
		{
			name: "out of order nodes - full connection",
			graph: &Graph{
				Nodes: []*Node{
					{X: 10, Y: 10, R: 3},
					{X: 3, Y: 3, R: 2},
					{X: 7, Y: 5, R: 2},
					{X: 5, Y: 3, R: 2},
					{X: 0, Y: 0, R: 3},
					{X: 7, Y: 8, R: 2},
				},
				Edges: []*Edge{
					{N1: &Node{X: 0, Y: 0, R: 3}, N2: &Node{X: 3, Y: 3, R: 2}},
					{N1: &Node{X: 3, Y: 3, R: 2}, N2: &Node{X: 5, Y: 3, R: 2}},
					{N1: &Node{X: 5, Y: 3, R: 2}, N2: &Node{X: 7, Y: 5, R: 2}},
					{N1: &Node{X: 7, Y: 5, R: 2}, N2: &Node{X: 7, Y: 8, R: 2}},
					{N1: &Node{X: 7, Y: 8, R: 2}, N2: &Node{X: 10, Y: 10, R: 3}},
				},
			},
			point1:   &Point{X: 0, Y: 0},
			point2:   &Point{X: 10, Y: 10},
			expected: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.expected, originConnectedEdges(tc.graph, tc.point1, tc.point2))
		})
	}
}

func TestExactOriginMatch(t *testing.T) {
	testCases := []struct {
		name     string
		ref      *Point
		comp     *Node
		expected bool
	}{
		{
			name:     "match",
			ref:      &Point{X: 200, Y: 200},
			comp:     &Node{X: 200, Y: 200, R: 5},
			expected: true,
		},
		{
			name:     "no match",
			ref:      &Point{X: 404, Y: 404},
			comp:     &Node{X: 0, Y: 0, R: 5},
			expected: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.expected, exactOriginMatch(tc.ref, tc.comp))
		})
	}
}
