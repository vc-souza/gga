package algo

import (
	"container/heap"
	"fmt"
	"math"
	"sort"

	"github.com/vc-souza/gga/ds"
)

/*
MSTAlgo describes the signature of an algorithm that can build a minimum
spanning tree of an undirected graph with weighted edges. If such an
algorithm is called on a directed graph, then ds.ErrUndefOp error
is returned.
*/
type MSTAlgo[T ds.Item] func(*ds.G[T]) (MST[T], error)

/*
An MST holds the edges of a minimum spanning tree
of an undirected graph with weighted edges.
*/
type MST[T ds.Item] []*ds.GE[T]

/*
MSTKruskal implements Kruskal's algorithm for finding a minimum spanning tree
of an undirected graph with weighted edges.

It is a greedy algorithm that applies the greedy-choice property by first
sorting all edges of the graph in order of non-decreasing edge weights, and
then iterating over the sorted list of edges, always picking the edge of
least weight (greedy choice, locally optimal) that connects previously
unlinked components. A disjoint-set data structure is used to keep track
of the components, and at the end of the algorithm, the list of edges
returned forms an MST of the original graph (globally optimal solution).

Expectations:
	- The graph is correctly built.
	- The graph is undirected.

Complexity:
	- Time:  O(E log V)
	- Space: Θ(V + E).
*/
func MSTKruskal[T ds.Item](g *ds.G[T]) (MST[T], error) {
	if g.Directed() {
		return nil, ds.ErrUndefOp
	}

	edges := make([]*ds.GE[T], g.EdgeCount())

	// By iterating over G.V and adding edges using their original
	// insertion order, we can guarantee that every call of the
	// algorithm on the same graph always yields the same MST,
	// since multiple MSTs might exist for the same graph.
	for i, eIdx := 0, 0; i < len(g.V); i++ {
		es := g.E[g.V[i].Ptr]

		copy(edges[eIdx:], es)

		eIdx += len(es)
	}

	// By using a stable sorting algorithm to sort the sequence
	// of edges in O(E log E) time, we make sure that if a tie
	// happens between edges of same weight, the original insertion
	// order is respected, and a consistent result is achieved.
	sort.Stable(ds.ByEdgeWeight[T](edges))

	d := ds.NewDSet[T]()

	for _, vtx := range g.V {
		d.MakeSet(vtx.Ptr)
	}

	max := g.VertexCount() - 1
	mst := MST[T]{}

	for _, e := range edges {
		if d.FindSet(e.Src) == d.FindSet(e.Dst) {
			continue
		}

		d.Union(e.Src, e.Dst)

		mst = append(mst, e)

		if len(mst) == max {
			break
		}
	}

	return mst, nil
}

// TODO: docs
type primVtx[T ds.Item] struct {
	// Ptr holds a reference to the original vertex.
	Ptr *T

	// BestWt holds the weight of the best edge found so far that can connect the vertex to the MST.
	BestWt float64

	// BestEdge holds the best edge found so far that can connect the vertex to the MST.
	BestEdge *ds.GE[T]
}

func (v primVtx[T]) String() string {
	return fmt.Sprintf(
		"%s edge:[%v] key:<%f>",
		(*v.Ptr).Label(),
		v.BestEdge,
		v.BestWt,
	)
}

// TODO: docs
type primVtxHeap[T ds.Item] []*primVtx[T]

func (h primVtxHeap[T]) Len() int           { return len(h) }
func (h primVtxHeap[T]) Less(i, j int) bool { return h[i].BestWt < h[j].BestWt }
func (h primVtxHeap[T]) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *primVtxHeap[T]) Push(x any) {
	*h = append(*h, x.(*primVtx[T]))
}

func (h *primVtxHeap[T]) Pop() any {
	n := len(*h)
	x := (*h)[n-1]

	(*h)[n-1] = nil
	*h = (*h)[:n-1]

	return x
}

// TODO: docs
func MSTPrim[T ds.Item](g *ds.G[T]) (MST[T], error) {
	if g.Directed() {
		return nil, ds.ErrUndefOp
	}

	vtxHeap := make(primVtxHeap[T], 0, g.VertexCount())
	vtxMap := map[*T]*primVtx[T]{}

	for i, vtx := range g.V {
		var wt float64

		// source
		if i == 0 {
			wt = 0
		} else {
			wt = math.Inf(1)
		}

		pVtx := &primVtx[T]{
			Ptr:    vtx.Ptr,
			BestWt: wt,
		}

		vtxHeap = append(vtxHeap, pVtx)
		vtxMap[pVtx.Ptr] = pVtx
	}

	heap.Init(&vtxHeap)

	mst := MST[T]{}

	// TODO: impl

	return mst, nil
}
