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
type CCViz struct {
	ThemedGraphViz

	CCs []algo.CC

	// OnCCVertex is called for every vertex, along with the index of its CC.
	OnCCVertex func(int, int)

	// OnCCEdge is called for any edge connecting vertices in the same CC.
	OnCCEdge func(int, int, int)
}

// NewCCViz initializes a new CCViz with NOOP hooks.
func NewCCViz(g *ds.G, ccs []algo.CC, t Theme) *CCViz {
	res := &CCViz{}

	res.CCs = ccs

	res.Graph = g
	res.Theme = t

	res.OnCCVertex = func(int, int) {}
	res.OnCCEdge = func(int, int, int) {}

	return res
}

// Traverse iterates over the results of any CC algorithm, calling its hooks when appropriate.
func (vi *CCViz) Traverse() error {
	for i := range vi.CCs {
		for _, v := range vi.CCs[i] {
			vi.OnCCVertex(v, i)

			for j := range vi.Graph.V[v].E {
				vi.OnCCEdge(v, j, i)
			}
		}
	}

	return nil
}
