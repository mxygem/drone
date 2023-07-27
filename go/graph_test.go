package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewGraph(t *testing.T) {
	expected := &Graph{Width: 100, Height: 200}

	assert.Equal(t, expected, NewGraph(100, 200))
}

func TestAddNode(t *testing.T) {
	testCases := []struct {
		name     string
		node     *Node
		graph    *Graph
		expected *Graph
	}{
		{
			name:  "no existing nodes",
			node:  &Node{X: 10, Y: 0, R: 5},
			graph: &Graph{},
			expected: &Graph{Nodes: []*Node{
				{X: 10, Y: 0, R: 5},
			}},
		},
		{
			name: "adding node to existing nodes",
			node: &Node{X: 5, Y: 5, R: 1},
			graph: &Graph{Nodes: []*Node{
				{X: 10, Y: 0, R: 5},
			}},
			expected: &Graph{Nodes: []*Node{
				{X: 10, Y: 0, R: 5},
				{X: 5, Y: 5, R: 1},
			}},
		},
		{
			name:     "nil node",
			node:     nil,
			graph:    &Graph{},
			expected: &Graph{},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tc.graph.AddNode(tc.node)

			assert.Equal(t, tc.expected, tc.graph)
		})
	}
}

func TestAddEdge(t *testing.T) {
	testCases := []struct {
		name     string
		node1    *Node
		node2    *Node
		graph    *Graph
		expected []*Edge
	}{
		{
			name:  "no existing nodes",
			node1: &Node{X: 10, Y: 0, R: 5},
			node2: &Node{X: 8, Y: 2, R: 5},
			graph: &Graph{},
			expected: []*Edge{
				{
					N1: &Node{X: 10, Y: 0, R: 5},
					N2: &Node{X: 8, Y: 2, R: 5},
				},
			},
		},
		{
			name:  "adding node to existing nodes",
			node1: &Node{X: 6, Y: 4, R: 5},
			node2: &Node{X: 4, Y: 6, R: 5},
			graph: &Graph{Edges: []*Edge{
				{
					N1: &Node{X: 10, Y: 0, R: 5},
					N2: &Node{X: 8, Y: 2, R: 5},
				},
			}},
			expected: []*Edge{
				{
					N1: &Node{X: 10, Y: 0, R: 5},
					N2: &Node{X: 8, Y: 2, R: 5},
				},
				{
					N1: &Node{X: 6, Y: 4, R: 5},
					N2: &Node{X: 4, Y: 6, R: 5},
				},
			},
		},
		{
			name:     "nil node1",
			node1:    nil,
			node2:    &Node{X: 6, Y: 6, R: 5},
			graph:    &Graph{},
			expected: nil,
		},
		{
			name:     "nil node2",
			node1:    &Node{X: 6, Y: 6, R: 5},
			node2:    nil,
			graph:    &Graph{},
			expected: nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tc.graph.AddEdge(tc.node1, tc.node2)

			assert.Equal(t, tc.expected, tc.graph.Edges)
		})
	}
}

func TestAreTransitivelyConnected(t *testing.T) {
	testCases := []struct {
		name     string
		graph    *Graph
		n1       *Node
		n2       *Node
		expected bool
	}{
		{
			name: "no connection due to no matching edge with beginning point",
			graph: &Graph{
				Edges: []*Edge{
					{N1: &Node{X: 0, Y: 0, R: 3}, N2: &Node{X: 0, Y: 4, R: 3}},
				},
			},
			n1:       &Node{X: 1, Y: 1},
			n2:       &Node{X: 0, Y: 4},
			expected: false,
		},
		{
			name: "single edge connection",
			graph: &Graph{
				Edges: []*Edge{
					{N1: &Node{X: 0, Y: 0, R: 3}, N2: &Node{X: 0, Y: 4, R: 3}},
				},
			},
			n1:       &Node{X: 0, Y: 0},
			n2:       &Node{X: 0, Y: 4},
			expected: true,
		},
		{
			name: "multiple edges match but not fully connected",
			graph: &Graph{
				Edges: []*Edge{
					{N1: &Node{X: 0, Y: 0, R: 2}, N2: &Node{X: 3, Y: 3, R: 2}},
					{N1: &Node{X: 3, Y: 3, R: 2}, N2: &Node{X: 6, Y: 6, R: 2}},
					{N1: &Node{X: 6, Y: 6, R: 2}, N2: &Node{X: 9, Y: 9, R: 2}},
				},
			},
			n1:       &Node{X: 0, Y: 0},
			n2:       &Node{X: 12, Y: 12},
			expected: false,
		},
		{
			name: "edges out of order but connected",
			graph: &Graph{
				Edges: []*Edge{
					{N1: &Node{X: 5, Y: 3, R: 2}, N2: &Node{X: 7, Y: 5, R: 2}},
					{N1: &Node{X: 7, Y: 8, R: 2}, N2: &Node{X: 10, Y: 10, R: 3}},
					{N1: &Node{X: 3, Y: 3, R: 2}, N2: &Node{X: 5, Y: 3, R: 2}},
					{N1: &Node{X: 0, Y: 0, R: 3}, N2: &Node{X: 3, Y: 3, R: 2}},
					{N1: &Node{X: 7, Y: 5, R: 2}, N2: &Node{X: 7, Y: 8, R: 2}},
				},
			},
			n1:       &Node{X: 0, Y: 0},
			n2:       &Node{X: 10, Y: 10},
			expected: true,
		},
		{
			name: "circular edges hit timeout",
			graph: &Graph{
				Edges: []*Edge{
					{N1: &Node{X: 0, Y: 0}, N2: &Node{X: 5, Y: 5}},
					{N1: &Node{X: 5, Y: 5}, N2: &Node{X: 10, Y: 10}},
					{N1: &Node{X: 10, Y: 10}, N2: &Node{X: 0, Y: 0}},
				},
			},
			n1:       &Node{X: 0, Y: 0},
			n2:       &Node{X: 10, Y: 10},
			expected: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.expected, tc.graph.AreTransitivelyConnected(tc.n1, tc.n2))
		})
	}
}

func TestFindEdge(t *testing.T) {
	testCases := []struct {
		name     string
		edges    []*Edge
		node     *Node
		expected *Edge
	}{
		{
			name: "edge with matching n1 not found",
			edges: []*Edge{
				{N1: &Node{X: 0, Y: 0}, N2: &Node{X: 3, Y: 3}},
			},
			node:     &Node{X: 2, Y: 2},
			expected: nil,
		},
		{
			name:     "no edges",
			node:     &Node{X: 2, Y: 2},
			expected: nil,
		},
		{
			name: "match found",
			edges: []*Edge{
				{N1: &Node{X: 0, Y: 0}, N2: &Node{X: 3, Y: 3}},
			},
			node:     &Node{X: 0, Y: 0},
			expected: &Edge{N1: &Node{X: 0, Y: 0}, N2: &Node{X: 3, Y: 3}},
		},
		{
			name: "match found later in collection",
			edges: []*Edge{
				{N1: &Node{X: 0, Y: 0, R: 2}, N2: &Node{X: 3, Y: 3, R: 2}},
				{N1: &Node{X: 3, Y: 3, R: 2}, N2: &Node{X: 5, Y: 3, R: 2}},
				{N1: &Node{X: 5, Y: 3, R: 2}, N2: &Node{X: 7, Y: 5, R: 3}},
				{N1: &Node{X: 7, Y: 5, R: 3}, N2: &Node{X: 7, Y: 8, R: 3}},
				{N1: &Node{X: 7, Y: 8, R: 2}, N2: &Node{X: 10, Y: 10, R: 3}},
			},
			node:     &Node{X: 5, Y: 3, R: 2},
			expected: &Edge{N1: &Node{X: 5, Y: 3, R: 2}, N2: &Node{X: 7, Y: 5, R: 3}},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.expected, findEdge(tc.edges, tc.node))
		})
	}
}
