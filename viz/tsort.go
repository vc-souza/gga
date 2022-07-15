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

	/*
		OnOrderEdge is called for every edge that connects two contiguous vertices
		in the final topological ordering. Note that such an edge might not actually
		exist in the graph, which gives the caller the opportunity to either create
		it or take any other action.
	*/
	OnOrderEdge func(*ds.GraphEdge[V], bool)
}

// NewTSortViz initializes a new TSortViz with NOOP hooks.
func NewTSortViz[V ds.Item](g *ds.Graph[V], ord *list.List, t Theme) *TSortViz[V] {
	res := &TSortViz[V]{}

	res.Order = ord

	res.Graph = g
	res.Theme = t

	res.OnVertexRank = func(*ds.GraphVertex[V], int) {}
	res.OnOrderEdge = func(*ds.GraphEdge[V], bool) {}

	return res
}

// Traverse iterates over the results of a Topological Sort execution, calling its hooks when appropriate.
func (vi *TSortViz[V]) Traverse() error {
	var rank int
	var prev *V
	var next *V

	for elem := vi.Order.Front(); elem != nil; elem = elem.Next() {
		rank++

		val, ok := elem.Value.(*V)

		if !ok {
			return ds.ErrInvalidType
		}

		vtx, _, ok := vi.Graph.GetVertex(val)

		if !ok {
			return errors.New("could not find vertex")
		}

		vi.OnVertexRank(vtx, rank)

		next = val

		if prev != nil && next != nil {
			e, _, ok := vi.Graph.GetEdge(prev, next)

			if !ok {
				e = &ds.GraphEdge[V]{Src: prev, Dst: next}
			}

			vi.OnOrderEdge(e, ok)
		}

		prev = next
	}

	return nil
}
