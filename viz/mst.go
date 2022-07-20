package viz

import (
	"errors"

	"github.com/vc-souza/gga/algo"
	"github.com/vc-souza/gga/ds"
)

// TODO: docs
type MSTViz[T ds.Item] struct {
	ThemedGraphViz[T]

	MST algo.MST[T]

	// TODO: docs
	OnMSTEdge func(*ds.GE[T])
}

// TODO: docs
func NewMSTViz[T ds.Item](g *ds.G[T], mst algo.MST[T], t Theme) *MSTViz[T] {
	res := &MSTViz[T]{}

	res.MST = mst

	res.Graph = g
	res.Theme = t

	res.OnMSTEdge = func(*ds.GE[T]) {}

	return res
}

// TODO: docs
func (vi *MSTViz[T]) Traverse() error {
	for _, e := range vi.MST {
		vi.OnMSTEdge(e)

		rev, _, ok := vi.Graph.GetEdge(e.Dst, e.Src)

		if !ok {
			return errors.New("could not find reverse edge")
		}

		vi.OnMSTEdge(rev)
	}

	return nil
}
