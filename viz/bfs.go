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
type BFSViz[T ds.Item] struct {
	ThemedGraphViz[T]

	Tree   algo.BFTree[T]
	Source *T

	// OnUnVertex is called when an unreachable vertex is found.
	OnUnVertex func(*ds.GV[T], *algo.BFNode[T])

	// OnSourceVertex is called when the source vertex is found.
	OnSourceVertex func(*ds.GV[T], *algo.BFNode[T])

	// OnTreeVertex is called when any tree vertex is found, including the source vertex.
	OnTreeVertex func(*ds.GV[T], *algo.BFNode[T])

	// OnTreeEdge is called when a tree edge is found.
	OnTreeEdge func(*ds.GE[T])
}

// NewBFSViz initializes a new BFSViz with NOOP hooks.
func NewBFSViz[T ds.Item](g *ds.G[T], tree algo.BFTree[T], src *T, t Theme) *BFSViz[T] {
	res := &BFSViz[T]{}

	res.Tree = tree
	res.Source = src

	res.Graph = g
	res.Theme = t

	res.OnUnVertex = func(*ds.GV[T], *algo.BFNode[T]) {}
	res.OnSourceVertex = func(*ds.GV[T], *algo.BFNode[T]) {}
	res.OnTreeVertex = func(*ds.GV[T], *algo.BFNode[T]) {}
	res.OnTreeEdge = func(*ds.GE[T]) {}

	return res
}

// Traverse iterates over the results of a BFS execution, calling its hooks when appropriate.
func (vi *BFSViz[T]) Traverse() error {
	for v, node := range vi.Tree {
		vtx, _, ok := vi.Graph.GetVertex(v)

		if !ok {
			return ds.ErrNoVtx
		}

		if math.IsInf(node.Distance, 1) {
			vi.OnUnVertex(vtx, node)
			continue
		}

		vi.OnTreeVertex(vtx, node)

		if node.Distance == 0 {
			vi.OnSourceVertex(vtx, node)
			continue
		}

		edge, _, ok := vi.Graph.GetEdge(node.Parent, v)

		if !ok {
			return ds.ErrNoEdge
		}

		vi.OnTreeEdge(edge)

		if vi.Graph.Directed() {
			continue
		}

		rev, _, ok := vi.Graph.GetEdge(v, node.Parent)

		if !ok {
			return ds.ErrNoRevEdge
		}

		vi.OnTreeEdge(rev)
	}

	return nil
}
