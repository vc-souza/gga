package viz

import (
	"github.com/vc-souza/gga/algo"
	"github.com/vc-souza/gga/ds"
)

/*
CCViz formats and exports an undirected graph after the execution of any algorithm
that discovers connected components. The output of the algorithm is traversed, and
hooks are provided so that custom formatting can be applied to the graph, its
vertices and edges.
*/
type CCViz[T ds.Item] struct {
	ThemedGraphViz[T]

	CCs []algo.CC[T]

	// OnCCVertex is called for every vertex, along with the index of its CC.
	OnCCVertex func(*ds.GV[T], int)

	// OnCCEdge is called for any edge connecting vertices in the same CC.
	OnCCEdge func(*ds.GE[T], int)
}

// NewCCViz initializes a new CCViz with NOOP hooks.
func NewCCViz[T ds.Item](g *ds.G[T], ccs []algo.CC[T], t Theme) *CCViz[T] {
	res := &CCViz[T]{}

	res.CCs = ccs

	res.Graph = g
	res.Theme = t

	res.OnCCVertex = func(*ds.GV[T], int) {}
	res.OnCCEdge = func(*ds.GE[T], int) {}

	return res
}

// Traverse iterates over the results of any CC algorithm, calling its hooks when appropriate.
func (vi *CCViz[T]) Traverse() error {
	for i, cc := range vi.CCs {
		for _, v := range cc {
			vtx, _, ok := vi.Graph.GetVertex(v)

			if !ok {
				return ds.ErrNoVtx
			}

			vi.OnCCVertex(vtx, i)

			for _, e := range vi.Graph.E[v] {
				vi.OnCCEdge(e, i)
			}
		}
	}

	return nil
}
