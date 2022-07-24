package viz

import (
	"github.com/vc-souza/gga/algo"
	"github.com/vc-souza/gga/ds"
)

/*
MSTViz formats and exports an undirected graph after the execution of any algorithm
that builds a minimum spanning tree. The output of the algorithm is traversed, and
hooks are provided so that custom formatting can be applied to the graph edges.
*/
type MSTViz struct {
	ThemedGraphViz

	MST algo.MST

	// OnMSTEdge is called for any edge that is a part of an MST.
	OnMSTEdge func(int, int)
}

// NewMSTViz initializes a new MSTViz with NOOP hooks.
func NewMSTViz(g *ds.G, mst algo.MST, t Theme) *MSTViz {
	res := &MSTViz{}

	res.MST = mst

	res.Graph = g
	res.Theme = t

	res.OnMSTEdge = func(int, int) {}

	return res
}

// Traverse iterates over the results of any MST algorithm, calling its hooks when appropriate.
func (vi *MSTViz) Traverse() error {
	for _, e := range vi.MST {
		vi.OnMSTEdge(e.Src, e.Index)

		iV, iE, ok := vi.Graph.GetEdgeIndex(
			vi.Graph.V[e.Dst].Item,
			vi.Graph.V[e.Src].Item,
		)

		if !ok {
			return ds.ErrNoRevEdge
		}

		vi.OnMSTEdge(iV, iE)
	}

	return nil
}
