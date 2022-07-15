package viz

import (
	"errors"
	"math"

	"github.com/vc-souza/gga/algo"
	"github.com/vc-souza/gga/ds"
)

/*
BFSViz formats and exports a graph after an execution of the BFS algorithm.
The output of the algorithm is traversed, and hooks are provided so that
custom formatting can be applied to the graph, its vertices and edges.
*/
type BFSViz[V ds.Item] struct {
	ThemedGraphViz[V]

	Tree   algo.BFTree[V]
	Source *V

	// OnUnVertex is called when an unreachable vertex is found.
	OnUnVertex func(*ds.GraphVertex[V], *algo.BFNode[V])

	// OnSourceVertex is called when the source vertex is found.
	OnSourceVertex func(*ds.GraphVertex[V], *algo.BFNode[V])

	// OnTreeVertex is called when any tree vertex is found, including the source vertex.
	OnTreeVertex func(*ds.GraphVertex[V], *algo.BFNode[V])

	// OnTreeEdge is called when a tree edge is found.
	OnTreeEdge func(*ds.GraphEdge[V])
}

// NewBFSViz initializes a new BFSViz with NOOP hooks.
func NewBFSViz[V ds.Item](g *ds.Graph[V], t algo.BFTree[V], src *V, theme Theme) *BFSViz[V] {
	res := &BFSViz[V]{}

	res.Tree = t
	res.Source = src

	res.Graph = g
	res.Theme = theme

	res.OnUnVertex = func(*ds.GraphVertex[V], *algo.BFNode[V]) {}
	res.OnSourceVertex = func(*ds.GraphVertex[V], *algo.BFNode[V]) {}
	res.OnTreeVertex = func(*ds.GraphVertex[V], *algo.BFNode[V]) {}
	res.OnTreeEdge = func(*ds.GraphEdge[V]) {}

	return res
}

// Traverse iterates over the results of a BFS execution, calling its hooks when appropriate.
func (vi *BFSViz[V]) Traverse() error {
	for v, node := range vi.Tree {
		vtx, _, ok := vi.Graph.GetVertex(v)

		if !ok {
			return errors.New("could not find vertex")
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
			return errors.New("could not find edge")
		}

		vi.OnTreeEdge(edge)
	}

	return nil
}
