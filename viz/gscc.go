package viz

import (
	"github.com/vc-souza/gga/ds"
)

/*
GSCCViz formats and exports a GSCC graph after it has been calculated. Its vertices
are traversed, and hooks are provided so that custom formatting can be applied.
*/
type GSCCViz struct {
	ThemedGraphViz

	// OnGSCCVertex is called for every vertex of the GSCC.
	OnGSCCVertex func(v int)
}

// NewGSCCViz initializes a new GSCCViz with NOOP hooks.
func NewGSCCViz(g *ds.G, t Theme) *GSCCViz {
	res := &GSCCViz{}

	res.Graph = g
	res.Theme = t

	res.OnGSCCVertex = func(v int) {}

	return res
}

// Traverse iterates over the vertices of a GSCC graph, calling its hooks when appropriate.
func (vi *GSCCViz) Traverse() error {
	for v := range vi.Graph.V {
		vi.OnGSCCVertex(v)
	}

	return nil
}
