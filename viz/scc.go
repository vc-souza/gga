package viz

import (
	"errors"

	"github.com/vc-souza/gga/algo"
	"github.com/vc-souza/gga/ds"
)

/*
SCCViz formats and exports a graph after an execution of any algorithm that discovers
the strongly connected components of a graph. The output of the algorithm is traversed,
and hooks are provided so that custom formatting can be applied to the graph,
its vertices and edges.
*/
type SCCViz[V ds.Item] struct {
	ThemedGraphViz[V]

	SCCs []algo.SCC[V]

	// OnSCCVertex is called for every vertex, along with the index of its SCC.
	OnSCCVertex func(*ds.GV[V], int)

	// OnSCCEdge is called for any edge connecting vertices in the same SCC.
	OnSCCEdge func(*ds.GE[V], int)

	// OnCrossSCCEdge is called for any edge connecting vertices in different SCCs.
	OnCrossSCCEdge func(*ds.GE[V], int, int)
}

// NewSCCViz initializes a new SCCViz with NOOP hooks.
func NewSCCViz[V ds.Item](g *ds.G[V], sccs []algo.SCC[V], t Theme) *SCCViz[V] {
	res := &SCCViz[V]{}

	res.SCCs = sccs

	res.Graph = g
	res.Theme = t

	res.OnSCCVertex = func(*ds.GV[V], int) {}
	res.OnSCCEdge = func(*ds.GE[V], int) {}
	res.OnCrossSCCEdge = func(*ds.GE[V], int, int) {}

	return res
}

// Traverse iterates over the results of any SCC algorithm, calling its hooks when appropriate.
func (vi *SCCViz[V]) Traverse() error {
	sets := map[*V]int{}

	for i, scc := range vi.SCCs {
		cc := i + 1

		for _, v := range scc {
			sets[v] = cc

			vtx, _, ok := vi.Graph.GetVertex(v)

			if !ok {
				return errors.New("could not find vertex")
			}

			vi.OnSCCVertex(vtx, cc)
		}
	}

	for _, es := range vi.Graph.E {
		for _, e := range es {
			cSrc := sets[e.Src]
			cDst := sets[e.Dst]

			if cSrc == cDst {
				vi.OnSCCEdge(e, cSrc)
			} else {
				vi.OnCrossSCCEdge(e, cSrc, cDst)
			}
		}
	}

	return nil
}
