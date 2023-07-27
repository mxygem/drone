package main

import (
	"log"
)

const (
	_safetyLimit int = 10
)

type Graph struct {
	Width  float64
	Height float64
	Nodes  Nodes
	Edges  []*Edge
}

type Edge struct {
	N1 *Node
	N2 *Node
}

type Point struct {
	X float64
	Y float64
}

func NewGraph(w, h float64) *Graph {
	return &Graph{
		Width:  w,
		Height: h,
	}
}

func (g *Graph) AddNode(n *Node) {
	if n == nil {
		return
	}

	g.Nodes = append(g.Nodes, n)
}

func (g *Graph) AddEdge(n1, n2 *Node) {
	if n1 == nil || n2 == nil {
		return
	}

	g.Edges = append(g.Edges, &Edge{N1: n1, N2: n2})
}

// AreTransitivelyConnected determines whether or not a pair of nodes are connected to one another
// via edges. It assumes that the node arguments have been provided and are valid nodes in the
// graph.
// NOTES on current implementation:
//   - It is **required** that the provided starting node **must** be the leading node in a found
//     edge. (aka: argument n1 must match edge n1)
//   - Additionally, it is assumed that each set of nodes will only have a single linear path if
//     connected. aka no edges will contain duplicate n1 nodes
func (g *Graph) AreTransitivelyConnected(n1, n2 *Node) bool {
	// look through edges for one with matching n1 nodes and to collection.
	var edges []*Edge
	for _, e := range g.Edges {
		if n1.X != e.N1.X && n1.Y != e.N1.Y {
			continue
		}

		edges = append(edges, e)
		break
	}
	if len(edges) == 0 {
		return false
	}

	for {
		if len(edges) >= _safetyLimit {
			log.Println("SAFETY LIMIT REACHED WHEN CHECKING TRANSITIVE CONNECTIVITY")
			return false
		}

		// look for next edge using the n2 from the last found edge
		e := findEdge(g.Edges, edges[len(edges)-1].N2)
		if e == nil {
			break
		}

		edges = append(edges, e)
	}

	lastEdge := edges[len(edges)-1]

	return lastEdge.N2.X == n2.X && lastEdge.N2.Y == n2.Y
}

func findEdge(es []*Edge, n *Node) *Edge {
	for _, e := range es {
		if e.N1.X != n.X || e.N1.Y != n.Y {
			continue
		}

		return e
	}

	return nil
}
