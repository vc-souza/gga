package viz

import (
	"errors"
	"io"

	"github.com/vc-souza/gga/algo"
	"github.com/vc-souza/gga/ds"
)

// TODO: docs
type DFSViz[V ds.Item] struct {
	AlgoViz

	Forest algo.DFSForest[V]
	Edges  *algo.EdgeTypes[V]
	Graph  *ds.Graph[V]

	// TODO: docs
	OnTreeVertex func(*ds.GraphVertex[V], *algo.DFSNode[V])

	// TODO: docs
	OnRootVertex func(*ds.GraphVertex[V], *algo.DFSNode[V])

	// TODO: docs
	OnTreeEdge func(*ds.GraphEdge[V])

	// TODO: docs
	OnForwardEdge func(*ds.GraphEdge[V])

	// TODO: docs
	OnBackEdge func(*ds.GraphEdge[V])

	// TODO: docs
	OnCrossEdge func(*ds.GraphEdge[V])
}

// TODO: docs
func NewDFSViz[V ds.Item](g *ds.Graph[V], f algo.DFSForest[V], e *algo.EdgeTypes[V]) *DFSViz[V] {
	res := &DFSViz[V]{}

	res.Forest = f
	res.Edges = e
	res.Graph = g

	res.OnTreeVertex = func(*ds.GraphVertex[V], *algo.DFSNode[V]) {}
	res.OnRootVertex = func(*ds.GraphVertex[V], *algo.DFSNode[V]) {}

	res.OnTreeEdge = func(*ds.GraphEdge[V]) {}
	res.OnForwardEdge = func(*ds.GraphEdge[V]) {}
	res.OnBackEdge = func(*ds.GraphEdge[V]) {}
	res.OnCrossEdge = func(*ds.GraphEdge[V]) {}

	return res
}

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

		// TODO: reverse edge
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
