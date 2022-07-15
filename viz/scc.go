package viz

import (
	"errors"

	"github.com/vc-souza/gga/algo"
	"github.com/vc-souza/gga/ds"
)

// TODO: docs
type SCCViz[V ds.Item] struct {
	ThemedGraphViz[V]

	SCCs []algo.SCC[V]

	// TODO: docs
	OnVertex func(*ds.GraphVertex[V], int)

	// TODO: docs
	OnEdge func(*ds.GraphEdge[V], int)

	// TODO: docs
	OnCrossEdge func(*ds.GraphEdge[V], int, int)
}

// TODO: docs
func NewSCCViz[V ds.Item](g *ds.Graph[V], sccs []algo.SCC[V], t Theme) *SCCViz[V] {
	res := &SCCViz[V]{}

	res.SCCs = sccs

	res.Graph = g
	res.Theme = t

	res.OnVertex = func(*ds.GraphVertex[V], int) {}
	res.OnEdge = func(*ds.GraphEdge[V], int) {}
	res.OnCrossEdge = func(*ds.GraphEdge[V], int, int) {}

	return res
}

// TODO: docs
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

			vi.OnVertex(vtx, cc)
		}
	}

	for _, es := range vi.Graph.Adj {
		for _, e := range es {
			cSrc := sets[e.Src]
			cDst := sets[e.Dst]

			if cSrc == cDst {
				vi.OnEdge(e, cSrc)
			} else {
				vi.OnCrossEdge(e, cSrc, cDst)
			}
		}
	}

	return nil
}
