package viz

import (
	"github.com/vc-souza/gga/algo"
	"github.com/vc-souza/gga/ds"
)

/*
SCCViz formats and exports a directed graph after the execution of any algorithm
that discovers strongly connected components. The output of the algorithm is
traversed, and hooks are provided so that custom formatting can be applied to
the graph, its vertices and edges.
*/
type SCCViz struct {
	ThemedGraphViz

	SCCs []algo.SCC

	// OnSCCVertex is called for every vertex, along with the index of its SCC.
	OnSCCVertex func(int, int)

	// OnSCCEdge is called for any edge connecting vertices in the same SCC.
	OnSCCEdge func(int, int, int)

	// OnCrossSCCEdge is called for any edge connecting vertices in different SCCs.
	OnCrossSCCEdge func(int, int, int, int)
}

// NewSCCViz initializes a new SCCViz with NOOP hooks.
func NewSCCViz(g *ds.G, sccs []algo.SCC, t Theme) *SCCViz {
	res := &SCCViz{}

	res.SCCs = sccs

	res.Graph = g
	res.Theme = t

	res.OnSCCVertex = func(int, int) {}
	res.OnSCCEdge = func(int, int, int) {}
	res.OnCrossSCCEdge = func(int, int, int, int) {}

	return res
}

// Traverse iterates over the results of any SCC algorithm, calling its hooks when appropriate.
func (vi *SCCViz) Traverse() error {
	sets := make([]int, vi.Graph.VertexCount())

	for i := range vi.SCCs {
		for _, v := range vi.SCCs[i] {
			vi.OnSCCVertex(v, i)
			sets[v] = i
		}
	}

	for v := range vi.Graph.V {
		for e := range vi.Graph.V[v].E {
			cSrc := sets[vi.Graph.V[v].E[e].Src]
			cDst := sets[vi.Graph.V[v].E[e].Dst]

			if cSrc == cDst {
				vi.OnSCCEdge(v, e, cSrc)
			} else {
				vi.OnCrossSCCEdge(v, e, cSrc, cDst)
			}
		}
	}

	return nil
}
