package viz

import (
	"github.com/vc-souza/gga/ds"
)

// TODO: docs
type GSCCViz[V ds.Item] struct {
	Graph *ds.Graph[ds.ItemGroup[V]]
	Theme Theme

	// TODO: docs
	OnGSCCVertex func(*ds.GraphVertex[ds.ItemGroup[V]])
}

func (v *GSCCViz[V]) GetGraph() *ds.Graph[ds.ItemGroup[V]] {
	return v.Graph
}

func (v *GSCCViz[V]) GetExtra() []string {
	return nil
}

func (v *GSCCViz[V]) GetTheme() Theme {
	return v.Theme
}

// TODO: docs
func NewGSCCViz[V ds.Item](g *ds.Graph[ds.ItemGroup[V]], t Theme) *GSCCViz[V] {
	res := &GSCCViz[V]{}

	res.Graph = g
	res.Theme = t

	res.OnGSCCVertex = func(*ds.GraphVertex[ds.ItemGroup[V]]) {}

	return res
}

// TODO: docs
func (vi *GSCCViz[V]) Traverse() error {
	for _, vtx := range vi.Graph.Verts {
		vi.OnGSCCVertex(vtx)
	}

	return nil
}
