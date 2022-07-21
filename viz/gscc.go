package viz

import (
	"github.com/vc-souza/gga/ds"
)

/*
GSCCViz formats and exports a GSCC graph after it has been calculated. Its vertices
are traversed, and hooks are provided so that custom formatting can be applied.
*/
type GSCCViz[T ds.Item] struct {
	Graph *ds.G[ds.Group[T]]
	Theme Theme

	// OnGSCCVertex is called for every vertex of the GSCC.
	OnGSCCVertex func(*ds.GV[ds.Group[T]])
}

func (v *GSCCViz[T]) GetGraph() *ds.G[ds.Group[T]] {
	return v.Graph
}

func (v *GSCCViz[T]) GetExtra() []string {
	return nil
}

func (v *GSCCViz[T]) GetTheme() Theme {
	return v.Theme
}

// NewGSCCViz initializes a new GSCCViz with NOOP hooks.
func NewGSCCViz[T ds.Item](g *ds.G[ds.Group[T]], t Theme) *GSCCViz[T] {
	res := &GSCCViz[T]{}

	res.Graph = g
	res.Theme = t

	res.OnGSCCVertex = func(*ds.GV[ds.Group[T]]) {}

	return res
}

// Traverse iterates over the vertices of a GSCC graph, calling its hooks when appropriate.
func (v *GSCCViz[T]) Traverse() error {
	for _, vtx := range v.Graph.V {
		v.OnGSCCVertex(vtx)
	}

	return nil
}
