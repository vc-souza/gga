package viz

import (
	"container/list"
	"errors"

	"github.com/vc-souza/gga/ds"
)

/*
TSortViz formats and exports a graph after an execution of the Topological Sort algorithm.
The output of the algorithm is traversed, and hooks are provided so that custom formatting
can be applied to the graph, its vertices and edges.
*/
type TSortViz[V ds.Item] struct {
	ThemedGraphViz[V]

	Order *list.List

	/*
		OnVertexRank is called for every vertex in the graph, along with the rank
		of the vertex in the final topological ordering of vertices.
	*/
	OnVertexRank func(*ds.GraphVertex[V], int)
}

// NewTSortViz initializes a new TSortViz with NOOP hooks.
func NewTSortViz[V ds.Item](g *ds.Graph[V], ord *list.List, theme Theme) *TSortViz[V] {
	res := &TSortViz[V]{}

	res.Order = ord

	res.Graph = g
	res.Theme = theme

	res.OnVertexRank = func(*ds.GraphVertex[V], int) {}

	return res
}

// Traverse iterates over the results of a Topological Sort execution, calling its hooks when appropriate.
func (vi *TSortViz[V]) Traverse() error {
	rank := 0

	for elem := vi.Order.Front(); elem != nil; elem = elem.Next() {
		rank++

		if val, ok := elem.Value.(*V); ok {

			if vtx, _, ok := vi.Graph.GetVertex(val); ok {
				vi.OnVertexRank(vtx, rank)
			} else {
				return errors.New("could not find vertex")
			}

		} else {
			return ds.ErrInvalidType
		}

	}

	return nil
}
