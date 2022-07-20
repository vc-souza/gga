package viz

import (
	"github.com/vc-souza/gga/algo"
	"github.com/vc-souza/gga/ds"
)

/*
DFSViz formats and exports a graph after an execution of the DFS algorithm.
The output of the algorithm is traversed, and hooks are provided so that
custom formatting can be applied to the graph, its vertices and edges.
*/
type DFSViz[T ds.Item] struct {
	ThemedGraphViz[T]

	Forest algo.DFForest[T]
	Edges  *algo.EdgeTypes[T]

	// OnTreeVertex is called for every vertex in the graph.
	OnTreeVertex func(*ds.GV[T], *algo.DFNode[T])

	// OnRootVertex is called when the root of a DF tree is found.
	OnRootVertex func(*ds.GV[T], *algo.DFNode[T])

	// OnTreeEdge is called when a tree edge is found.
	OnTreeEdge func(*ds.GE[T])

	// OnForwardEdge is called when a forward edge is found.
	OnForwardEdge func(*ds.GE[T])

	// OnBackEdge is called when a back edge is found.
	OnBackEdge func(*ds.GE[T])

	// OnCrossEdge is called when a cross edge is found.
	OnCrossEdge func(*ds.GE[T])
}

// NewDFSViz initializes a new DFSViz with NOOP hooks.
func NewDFSViz[T ds.Item](g *ds.G[T], f algo.DFForest[T], e *algo.EdgeTypes[T], t Theme) *DFSViz[T] {
	res := &DFSViz[T]{}

	res.Forest = f
	res.Edges = e

	res.Graph = g
	res.Theme = t

	res.OnTreeVertex = func(*ds.GV[T], *algo.DFNode[T]) {}
	res.OnRootVertex = func(*ds.GV[T], *algo.DFNode[T]) {}

	res.OnTreeEdge = func(*ds.GE[T]) {}
	res.OnForwardEdge = func(*ds.GE[T]) {}
	res.OnBackEdge = func(*ds.GE[T]) {}
	res.OnCrossEdge = func(*ds.GE[T]) {}

	return res
}

// Traverse iterates over the results of a DFS execution, calling its hooks when appropriate.
func (vi *DFSViz[T]) Traverse() error {
	for v, node := range vi.Forest {
		vtx, _, ok := vi.Graph.GetVertex(v)

		if !ok {
			return ds.ErrVtxNotExists
		}

		vi.OnTreeVertex(vtx, node)

		if node.Parent == nil {
			vi.OnRootVertex(vtx, node)
			continue
		}

		edge, _, ok := vi.Graph.GetEdge(node.Parent, v)

		if !ok {
			return ds.ErrEdgeNotExists
		}

		vi.OnTreeEdge(edge)

		if vi.Graph.Directed() {
			continue
		}

		rev, _, ok := vi.Graph.GetEdge(v, node.Parent)

		if !ok {
			return ds.ErrRevEdgeNotExists
		}

		vi.OnTreeEdge(rev)
	}

	for _, e := range vi.Edges.Forward {
		vi.OnForwardEdge(e)
	}

	for _, e := range vi.Edges.Back {
		vi.OnBackEdge(e)
	}

	for _, e := range vi.Edges.Cross {
		vi.OnCrossEdge(e)
	}

	return nil
}
