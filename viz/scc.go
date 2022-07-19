package viz

import (
	"errors"

	"github.com/vc-souza/gga/algo"
	"github.com/vc-souza/gga/ds"
)

/*
SCCViz formats and exports a graph after an execution of any algorithm that discovers
the strongly connected components of a directed graph. The output of the algorithm
is traversed, and hooks are provided so that custom formatting can be applied to
the graph, its vertices and edges.
*/
type SCCViz[T ds.Item] struct {
	ThemedGraphViz[T]

	SCCs []algo.SCC[T]

	// OnSCCVertex is called for every vertex, along with the index of its SCC.
	OnSCCVertex func(*ds.GV[T], int)

	// OnSCCEdge is called for any edge connecting vertices in the same SCC.
	OnSCCEdge func(*ds.GE[T], int)

	// OnCrossSCCEdge is called for any edge connecting vertices in different SCCs.
	OnCrossSCCEdge func(*ds.GE[T], int, int)
}

// NewSCCViz initializes a new SCCViz with NOOP hooks.
func NewSCCViz[T ds.Item](g *ds.G[T], sccs []algo.SCC[T], t Theme) *SCCViz[T] {
	res := &SCCViz[T]{}

	res.SCCs = sccs

	res.Graph = g
	res.Theme = t

	res.OnSCCVertex = func(*ds.GV[T], int) {}
	res.OnSCCEdge = func(*ds.GE[T], int) {}
	res.OnCrossSCCEdge = func(*ds.GE[T], int, int) {}

	return res
}

// Traverse iterates over the results of any SCC algorithm, calling its hooks when appropriate.
func (vi *SCCViz[T]) Traverse() error {
	sets := map[*T]int{}

	for i, scc := range vi.SCCs {
		for _, v := range scc {
			sets[v] = i

			vtx, _, ok := vi.Graph.GetVertex(v)

			if !ok {
				return errors.New("could not find vertex")
			}

			vi.OnSCCVertex(vtx, i)
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
