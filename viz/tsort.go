package viz

import (
	"container/list"
	"errors"

	"github.com/vc-souza/gga/ds"
)

type TSortViz[V ds.Item] struct {
	ThemedGraphViz[V]

	Order *list.List

	// TODO: docs
	OnVertex func(*ds.GraphVertex[V], int)
}

// TODO: docs
func NewTSortViz[V ds.Item](g *ds.Graph[V], ord *list.List, theme Theme) *TSortViz[V] {
	res := &TSortViz[V]{}

	res.Graph = g
	res.Order = ord
	res.Theme = theme

	res.OnVertex = func(*ds.GraphVertex[V], int) {}

	return res
}

// TODO: docs!
func (vi *TSortViz[V]) Traverse() error {
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

	return nil
}
