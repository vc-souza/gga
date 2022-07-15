package viz

import (
	"container/list"
	"errors"
	"io"

	"github.com/vc-souza/gga/ds"
)

type TSortViz[V ds.Item] struct {
	Graph *ds.Graph[V]
	Order *list.List

	Theme Theme

	// TODO: docs
	OnVertex func(*ds.GraphVertex[V], int)
}

// TODO: docs
func NewTSortViz[V ds.Item](g *ds.Graph[V], ord *list.List) *TSortViz[V] {
	res := &TSortViz[V]{}

	res.Graph = g
	res.Order = ord

	res.OnVertex = func(*ds.GraphVertex[V], int) {}

	return res
}

func (vi *TSortViz[V]) Export(w io.Writer) error {
	ex := NewExporter(vi.Graph)

	ResetGraphFmt(vi.Graph)
	SetTheme(ex, vi.Theme)

	for rank, elem := 1, vi.Order.Front(); elem != nil; rank, elem = rank+1, elem.Next() {
		if val, ok := elem.Value.(*V); ok {
			if vtx, _, ok := vi.Graph.GetVertex(val); ok {
				vi.OnVertex(vtx, rank)
			} else {
				return errors.New("could not find vertex")
			}
		} else {
			return ds.ErrInvalidType
		}

	}

	ex.Export(w)

	return nil
}
