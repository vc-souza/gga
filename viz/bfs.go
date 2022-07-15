package viz

import (
	"errors"
	"io"
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
	Tree   algo.BFTree[V]
	Graph  *ds.Graph[V]
	Source *V

	Theme Theme

	// OnUnVertex is called when an unreachable vertex is found.
	OnUnVertex func(*ds.GraphVertex[V], *algo.BFNode[V])

	// OnSourceVertex is called when the source vertex is found.
	OnSourceVertex func(*ds.GraphVertex[V], *algo.BFNode[V])

	// OnTreeVertex is called when any tree vertex is found, including the source vertex.
	OnTreeVertex func(*ds.GraphVertex[V], *algo.BFNode[V])

	// OnTreeEdge is called when a tree edge is found.
	OnTreeEdge func(*ds.GraphEdge[V])
}

// NewBFSViz initializes a new BFSViz with NOOP hooks and no custom formatting.
func NewBFSViz[V ds.Item](g *ds.Graph[V], t algo.BFTree[V], src *V) *BFSViz[V] {
	res := &BFSViz[V]{}

	res.Tree = t
	res.Graph = g
	res.Source = src

	res.OnUnVertex = func(*ds.GraphVertex[V], *algo.BFNode[V]) {}
	res.OnSourceVertex = func(*ds.GraphVertex[V], *algo.BFNode[V]) {}
	res.OnTreeVertex = func(*ds.GraphVertex[V], *algo.BFNode[V]) {}
	res.OnTreeEdge = func(*ds.GraphEdge[V]) {}

	return res
}

/*
Export traverses the results of a BFS execution, calling its hooks when appropriate.
The graph is then exported to the given io.Writer, using the standard viz.Exporter.
*/
func (vi *BFSViz[V]) Export(w io.Writer) error {
	ex := NewExporter(vi.Graph)

	ResetGraphFmt(vi.Graph)
	SetTheme(ex, vi.Theme)

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

	ex.Export(w)

	return nil
}
