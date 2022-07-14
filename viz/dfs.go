package viz

import (
	"errors"
	"io"

	"github.com/vc-souza/gga/algo"
	"github.com/vc-souza/gga/ds"
)

/*
DFSViz formats and exports a graph after an execution of the DFS algorithm.
The output of the algorithm is traversed, and hooks are provided so that
custom formatting can be applied to the graph, its vertices and edges.
*/
type DFSViz[V ds.Item] struct {
	AlgoViz

	Forest algo.DFForest[V]
	Edges  *algo.EdgeTypes[V]
	Graph  *ds.Graph[V]

	// OnTreeVertex is called for every vertex in the graph.
	OnTreeVertex func(*ds.GraphVertex[V], *algo.DFNode[V])

	// OnRootVertex is called when the root of a DF tree is found.
	OnRootVertex func(*ds.GraphVertex[V], *algo.DFNode[V])

	// OnTreeEdge is called when a tree edge is found.
	OnTreeEdge func(*ds.GraphEdge[V])

	// OnForwardEdge is called when a forward edge is found.
	OnForwardEdge func(*ds.GraphEdge[V])

	// OnBackEdge is called when a back edge is found.
	OnBackEdge func(*ds.GraphEdge[V])

	// OnCrossEdge is called when a cross edge is found.
	OnCrossEdge func(*ds.GraphEdge[V])
}

// NewDFSViz initializes a new DFSViz with NOOP hooks and no custom formatting.
func NewDFSViz[V ds.Item](g *ds.Graph[V], f algo.DFForest[V], e *algo.EdgeTypes[V]) *DFSViz[V] {
	res := &DFSViz[V]{}

	res.Forest = f
	res.Edges = e
	res.Graph = g

	res.OnTreeVertex = func(*ds.GraphVertex[V], *algo.DFNode[V]) {}
	res.OnRootVertex = func(*ds.GraphVertex[V], *algo.DFNode[V]) {}

	res.OnTreeEdge = func(*ds.GraphEdge[V]) {}
	res.OnForwardEdge = func(*ds.GraphEdge[V]) {}
	res.OnBackEdge = func(*ds.GraphEdge[V]) {}
	res.OnCrossEdge = func(*ds.GraphEdge[V]) {}

	return res
}

/*
Export traverses the results of a DFS execution, calling its hooks when appropriate.
The graph is then exported to the given io.Writer, using the standard viz.Exporter.
*/
func (vi *DFSViz[V]) Export(w io.Writer) error {
	ex := NewExporter(vi.Graph)

	ex.DefaultGraphFmt = vi.DefaultGraphFmt
	ex.DefaultVertexFmt = vi.DefaultVertexFmt
	ex.DefaultEdgeFmt = vi.DefaultEdgeFmt

	ResetGraphFmt(vi.Graph)

	for v, node := range vi.Forest {
		vtx, _, ok := vi.Graph.GetVertex(v)

		if !ok {
			return errors.New("could not find vertex")
		}

		vi.OnTreeVertex(vtx, node)

		if node.Parent == nil {
			vi.OnRootVertex(vtx, node)
			continue
		}

		edge, _, ok := vi.Graph.GetEdge(node.Parent, v)

		if !ok {
			return errors.New("could not find edge")
		}

		vi.OnTreeEdge(edge)
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

	ex.Export(w)

	return nil
}
