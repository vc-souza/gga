package viz

import (
	"math"

	"github.com/vc-souza/gga/algo"
	"github.com/vc-souza/gga/ds"
)

/*
BFSViz formats and exports a graph after the execution of the BFS algorithm.
The output of the algorithm is traversed, and hooks are provided so that
custom formatting can be applied to the graph, its vertices and edges.
*/
type BFSViz struct {
	ThemedGraphViz

	Tree   algo.BFTree
	Source int

	// OnUnVertex is called when an unreachable vertex is found.
	OnUnVertex func(int, algo.BFNode)

	// OnSourceVertex is called when the source vertex is found.
	OnSourceVertex func(int, algo.BFNode)

	// OnTreeVertex is called when any tree vertex is found, including the source vertex.
	OnTreeVertex func(int, algo.BFNode)

	// OnTreeEdge is called when a tree edge is found.
	OnTreeEdge func(int, int)
}

// NewBFSViz initializes a new BFSViz with NOOP hooks.
func NewBFSViz(g *ds.G, tree algo.BFTree, src int, t Theme) *BFSViz {
	res := &BFSViz{}

	res.Tree = tree
	res.Source = src

	res.Graph = g
	res.Theme = t

	res.OnUnVertex = func(int, algo.BFNode) {}
	res.OnSourceVertex = func(int, algo.BFNode) {}
	res.OnTreeVertex = func(int, algo.BFNode) {}
	res.OnTreeEdge = func(int, int) {}

	return res
}

// Traverse iterates over the results of a BFS execution, calling its hooks when appropriate.
func (vi *BFSViz) Traverse() error {
	for v, node := range vi.Tree {
		if math.IsInf(node.Distance, 1) {
			vi.OnUnVertex(v, node)
			continue
		}

		vi.OnTreeVertex(v, node)

		if node.Distance == 0 {
			vi.OnSourceVertex(v, node)
			continue
		}

		iV, iE, ok := vi.Graph.GetEdgeIndex(
			vi.Graph.V[node.Parent].Item,
			vi.Graph.V[v].Item,
		)

		if !ok {
			return ds.ErrNoEdge
		}

		vi.OnTreeEdge(iV, iE)

		if vi.Graph.Directed() {
			continue
		}

		iV, iE, ok = vi.Graph.GetEdgeIndex(
			vi.Graph.V[v].Item,
			vi.Graph.V[node.Parent].Item,
		)

		if !ok {
			return ds.ErrNoRevEdge
		}

		vi.OnTreeEdge(iV, iE)
	}

	return nil
}
