package main

import (
	"log"
)

var (
	// W & H represent the dimensions of a given field. **Not currently implemented.**
	W = 10.0
	H = 10.0
	// InputNodes is a collection of node locations and their respective radii. The following nodes
	// **are** all transitively connected between the first and last node.
	InputNodes = Nodes{
		{X: 4, Y: 6, R: 2},
		{X: 0, Y: 10, R: 2},
		{X: 6, Y: 4, R: 2},
		{X: 2, Y: 8, R: 2},
		{X: 10, Y: 0, R: 2},
		{X: 8, Y: 2, R: 2},
	}
	// EdgePointStart/End represent points along the edge of the field for use in checking
	// transitive connectivity of nodes. Each point is currently required to have an exact match
	// within the provided InputNodes.
	EdgePointStart = &Point{X: 0, Y: 10}
	EdgePointEnd   = &Point{X: 10, Y: 0}
)

func main() {
	log.Printf("Checking transitive connectivity from %v to %v\n", EdgePointStart, EdgePointEnd)
	g := NewGraph(W, H)

	handleNodes(g, InputNodes)
	log.Printf("Added %d nodes and %d edges\n", len(g.Nodes), len(g.Edges))

	connected := originConnectedEdges(g, EdgePointStart, EdgePointEnd)
	log.Printf("Connected? %v", connected)
}
