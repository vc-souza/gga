package viz

import (
	"errors"
	"io"
	"math"

	"github.com/vc-souza/gga/algo"
	"github.com/vc-souza/gga/ds"
)

// TODO: docs
type BFSViz[V ds.Item] struct {
	Graph  *ds.Graph[V]
	Tree   algo.BFSTree[V]
	Source *V

	DefaultGraphFmt  ds.FmtAttrs
	DefaultVertexFmt ds.FmtAttrs
	DefaultEdgeFmt   ds.FmtAttrs

	// TODO: docs
	OnUnVertex func(*ds.GraphVertex[V], *algo.BFSNode[V])

	// TODO: docs
	OnSourceVertex func(*ds.GraphVertex[V], *algo.BFSNode[V])

	// TODO: docs
	OnTreeVertex func(*ds.GraphVertex[V], *algo.BFSNode[V])

	// TODO: docs
	OnTreeEdge func(*ds.GraphEdge[V])
}

// TODO: docs
func NewBFSViz[V ds.Item](g *ds.Graph[V], t algo.BFSTree[V], src *V) *BFSViz[V] {
	res := &BFSViz[V]{}

	res.Graph = g
	res.Tree = t
	res.Source = src

	res.OnUnVertex = func(*ds.GraphVertex[V], *algo.BFSNode[V]) {}
	res.OnSourceVertex = func(*ds.GraphVertex[V], *algo.BFSNode[V]) {}
	res.OnTreeVertex = func(*ds.GraphVertex[V], *algo.BFSNode[V]) {}
	res.OnTreeEdge = func(*ds.GraphEdge[V]) {}

	return res
}

// TODO: docs
func (vi *BFSViz[V]) Export(w io.Writer) error {
	ex := NewDotExporter(vi.Graph)

	ex.DefaultGraphFmt = vi.DefaultGraphFmt
	ex.DefaultVertexFmt = vi.DefaultVertexFmt
	ex.DefaultEdgeFmt = vi.DefaultEdgeFmt

	for v, node := range vi.Tree {
		vtx, _, ok := vi.Graph.GetVertex(v)

		if !ok {
			return errors.New("could not find vertex")
		}

		vtx.ResetFmt()

		if node.Distance == math.MaxInt32 {
			vi.OnUnVertex(vtx, node)
			continue
		}

		vi.OnTreeVertex(vtx, node)

		if node.Distance == 0 {
			vi.OnSourceVertex(vtx, node)
			continue
		}

		edg, _, ok := vi.Graph.GetEdge(node.Parent, v)

		if !ok {
			return errors.New("could not find edge")
		}

		edg.ResetFmt()
		vi.OnTreeEdge(edg)

		if vi.Graph.Directed() {
			continue
		}

		rev, _, ok := vi.Graph.GetEdge(v, node.Parent)

		if !ok {
			return errors.New("could not find reverse edge")
		}

		rev.ResetFmt()
		vi.OnTreeEdge(rev)
	}

	vi.Graph.Accept(ex)

	ex.Export(w)

	return nil
}
