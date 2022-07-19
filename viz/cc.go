package viz

import (
	"errors"

	"github.com/vc-souza/gga/algo"
	"github.com/vc-souza/gga/ds"
)

/*
CCViz formats and exports a graph after an execution of any algorithm that discovers
the connected components of an undirected graph. The output of the algorithm is
traversed, and hooks are provided so that custom formatting can be applied to
the graph, its vertices and edges.
*/
type CCViz[V ds.Item] struct {
	ThemedGraphViz[V]

	CCs []algo.CC[V]

	// OnCCVertex is called for every vertex, along with the index of its CC.
	OnCCVertex func(*ds.GV[V], int)

	// OnCCEdge is called for any edge connecting vertices in the same CC.
	OnCCEdge func(*ds.GE[V], int)
}

// NewCCViz initializes a new CCViz with NOOP hooks.
func NewCCViz[V ds.Item](g *ds.G[V], ccs []algo.CC[V], t Theme) *CCViz[V] {
	res := &CCViz[V]{}

	res.CCs = ccs

	res.Graph = g
	res.Theme = t

	res.OnCCVertex = func(*ds.GV[V], int) {}
	res.OnCCEdge = func(*ds.GE[V], int) {}

	return res
}

// Traverse iterates over the results of any CC algorithm, calling its hooks when appropriate.
func (vi *CCViz[V]) Traverse() error {
	for i, cc := range vi.CCs {
		for _, v := range cc {
			vtx, _, ok := vi.Graph.GetVertex(v)

			if !ok {
				return errors.New("could not find vertex")
			}

			vi.OnCCVertex(vtx, i)

			for _, e := range vi.Graph.E[v] {
				vi.OnCCEdge(e, i)
			}
		}
	}

	return nil
}
