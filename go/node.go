package main

import (
	"log"
	"math"
	"sort"
)

type Node struct {
	X float64
	Y float64
	R float64
}

// Nodes implements the sort interface to allow for sorting coordinates.
type Nodes []*Node

func (n Nodes) Len() int      { return len(n) }
func (n Nodes) Swap(i, j int) { n[i], n[j] = n[j], n[i] }
func (n Nodes) Less(i, j int) bool {
	return n[i].X < n[j].X || (n[i].X < n[j].X && n[i].Y < n[j].Y)
}

// handleNodes configures a given graph by adding the provided set of nodes to its collection and
// finding any edges between each.
func handleNodes(g *Graph, ns Nodes) {
	// sort nodes for consistency
	sort.Sort(ns)

	// add nodes to graph
	for _, n := range ns {
		g.AddNode(n)
	}

	// find edges
	for i, bn := range g.Nodes {
		// check base node against remaining nodes
		for _, cn := range g.Nodes[i:] {
			// skip nodes that aren't edges
			if !isEdge(bn, cn) {
				continue
			}

			g.AddEdge(bn, cn)
			break
		}
	}
}

// originConnectedEdges determines whether or not two points are transitively connected by two
// nodes by explicit exact matches of the start and end node origins. For the simplification of the
// current implementation, it assumes that all nodes in graph have unique origin points.
// Required: All arguments, at least two nodes, and at least one edge.
func originConnectedEdges(g *Graph, p1, p2 *Point) bool {
	if g == nil || p1 == nil || p2 == nil || len(g.Nodes) < 2 || len(g.Edges) == 0 {
		return false
	}

	// find nodes matching p1 and p2 points
	var n1, n2 *Node
	for _, n := range g.Nodes {
		// stop searching if we have both matches.
		if n1 != nil && n2 != nil {
			break
		}

		switch {
		case exactOriginMatch(p1, n):
			n1 = n
		case exactOriginMatch(p2, n):
			n2 = n
		}
	}
	if n1 == nil || n2 == nil {
		log.Printf("could not find both points in graph's nodes. found n1? %v n2? %v\n", n1 == nil, n2 == nil)
		return false
	}

	return g.AreTransitivelyConnected(n1, n2)
}

func exactOriginMatch(ref *Point, comp *Node) bool {
	return ref.X == comp.X && ref.Y == comp.Y
}

// isEdge determines whether or not two given nodes are edges. Provided nodes are required to have
// differing origin points in order to be evaluated.
func isEdge(n1, n2 *Node) bool {
	if n1.X == n2.X && n1.Y == n2.Y {
		return false
	}

	dist := math.Sqrt(math.Pow(n2.X-n1.X, 2) + math.Pow(n2.Y-n1.Y, 2))
	radiiSum := n1.R + n2.R

	return radiiSum >= dist
}
