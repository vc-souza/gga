package viz

import (
	"github.com/vc-souza/gga/algo"
	"github.com/vc-souza/gga/ds"
)

/*
MSTViz formats and exports an undirected graph after the execution of any algorithm
that can discover a minimum spanning tree. The output of the algorithm is traversed,
and hooks are provided so that custom formatting can be applied to the graph edges.
*/
type MSTViz[T ds.Item] struct {
	ThemedGraphViz[T]

	MST algo.MST[T]

	// TODO: docs
	OnMSTEdge func(*ds.GE[T])
}

// NewMSTViz initializes a new MSTViz with NOOP hooks.
func NewMSTViz[T ds.Item](g *ds.G[T], mst algo.MST[T], t Theme) *MSTViz[T] {
	res := &MSTViz[T]{}

	res.MST = mst

	res.Graph = g
	res.Theme = t

	res.OnMSTEdge = func(*ds.GE[T]) {}

	return res
}

// Traverse iterates over the results of any MST algorithm, calling its hooks when appropriate.
func (vi *MSTViz[T]) Traverse() error {
	for _, e := range vi.MST {
		vi.OnMSTEdge(e)

		rev, _, ok := vi.Graph.GetEdge(e.Dst, e.Src)

		if !ok {
			return ds.ErrRevEdgeNotExists
		}

		vi.OnMSTEdge(rev)
	}

	return nil
}
