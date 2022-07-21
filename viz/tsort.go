package viz

import (
	"github.com/vc-souza/gga/ds"
)

/*
TSortViz formats and exports a directed graph after the execution
of the Topological Sort algorithm. The output of the algorithm is
traversed, and hooks are provided so that custom formatting can be
applied to the graph, its vertices and edges.
*/
type TSortViz[T ds.Item] struct {
	ThemedGraphViz[T]

	Order []*T

	/*
		OnVertexRank is called for every vertex in the graph, along with the rank
		of the vertex in the final topological ordering of vertices.
	*/
	OnVertexRank func(*ds.GV[T], int)

	/*
		OnOrderEdge is called for every edge that connects two contiguous vertices
		in the final topological ordering. Note that such an edge might not actually
		exist in the graph, which gives the caller the opportunity to either create
		it or take any other action.
	*/
	OnOrderEdge func(*ds.GE[T], bool)
}

// NewTSortViz initializes a new TSortViz with NOOP hooks.
func NewTSortViz[T ds.Item](g *ds.G[T], ord []*T, t Theme) *TSortViz[T] {
	res := &TSortViz[T]{}

	res.Order = ord

	res.Graph = g
	res.Theme = t

	res.OnVertexRank = func(*ds.GV[T], int) {}
	res.OnOrderEdge = func(*ds.GE[T], bool) {}

	return res
}

// Traverse iterates over the results of a Topological Sort execution, calling its hooks when appropriate.
func (vi *TSortViz[T]) Traverse() error {
	var rank int
	var prev *T
	var next *T

	for _, v := range vi.Order {
		rank++

		vtx, _, ok := vi.Graph.GetVertex(v)

		if !ok {
			return ds.ErrVtxNotExists
		}

		vi.OnVertexRank(vtx, rank)

		next = v

		if prev != nil && next != nil {
			e, _, ok := vi.Graph.GetEdge(prev, next)

			if !ok {
				e = &ds.GE[T]{Src: prev, Dst: next}
			}

			vi.OnOrderEdge(e, ok)
		}

		prev = next
	}

	return nil
}
