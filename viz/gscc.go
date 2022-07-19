package viz

import (
	"github.com/vc-souza/gga/ds"
)

/*
GSCCViz formats and exports a GSCC graph after it has been calculated. Its vertices
are traversed, and hooks are provided so that custom formatting can be applied.
*/
type GSCCViz[V ds.Item] struct {
	Graph *ds.G[ds.Group[V]]
	Theme Theme

	// OnGSCCVertex is called for every vertex of the GSCC.
	OnGSCCVertex func(*ds.GV[ds.Group[V]])
}

func (v *GSCCViz[V]) GetGraph() *ds.G[ds.Group[V]] {
	return v.Graph
}

func (v *GSCCViz[V]) GetExtra() []string {
	return nil
}

func (v *GSCCViz[V]) GetTheme() Theme {
	return v.Theme
}

// NewGSCCViz initializes a new GSCCViz with NOOP hooks.
func NewGSCCViz[V ds.Item](g *ds.G[ds.Group[V]], t Theme) *GSCCViz[V] {
	res := &GSCCViz[V]{}

	res.Graph = g
	res.Theme = t

	res.OnGSCCVertex = func(*ds.GV[ds.Group[V]]) {}

	return res
}

// Traverse iterates over the vertices of a GSCC graph, calling its hooks when appropriate.
func (vi *GSCCViz[V]) Traverse() error {
	for _, vtx := range vi.Graph.V {
		vi.OnGSCCVertex(vtx)
	}

	return nil
}
