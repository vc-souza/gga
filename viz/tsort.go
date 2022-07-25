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
type TSortViz struct {
	ThemedGraphViz

	Order []int

	/*
		OnVertexRank is called for every vertex in the graph, along with the rank
		of the vertex in the final topological ordering of vertices.
	*/
	OnVertexRank func(int, int)

	/*
		OnOrderEdge is called for every edge that connects two contiguous vertices
		in the final topological ordering. Note that such an edge might not actually
		exist in the graph, which gives the caller the opportunity to either create
		it or take any other action.
	*/
	OnOrderEdge func(int, int, int, bool)
}

// NewTSortViz initializes a new TSortViz with NOOP hooks.
func NewTSortViz(g *ds.G, ord []int, t Theme) *TSortViz {
	res := &TSortViz{}

	res.Order = ord

	res.Graph = g
	res.Theme = t

	res.OnVertexRank = func(int, int) {}
	res.OnOrderEdge = func(int, int, int, bool) {}

	return res
}

// Traverse iterates over the results of a Topological Sort execution, calling its hooks when appropriate.
func (vi *TSortViz) Traverse() error {
	var rank int
	prev := -1
	next := -1

	for _, v := range vi.Order {
		rank++

		vi.OnVertexRank(v, rank)

		next = v

		if prev != -1 {
			_, idx, ok := vi.Graph.EdgeIndex(
				vi.Graph.V[prev].Item,
				vi.Graph.V[next].Item,
			)

			vi.OnOrderEdge(prev, idx, next, ok)
		}

		prev = next
	}

	return nil
}
