package algo

import (
	"container/heap"
	"errors"
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
	// ptr holds a reference to the original vertex.
	ptr *T

	// wt holds the weight of the best edge found so far that can connect the vertex to the MST.
	wt float64

	// edge holds the best edge found so far that can connect the vertex to the MST.
	edge *ds.GE[T]

	// in tells if the vertex is still in the heap.
	in bool

	// index stores the index of the vertex, in the heap.
	index int
}

func (v primVtx[T]) String() string {
	return fmt.Sprintf(
		"%s in:%t i:<%d> wt:<%f> edge:[%v]",
		(*v.ptr).Label(),
		v.in,
		v.index,
		v.wt,
		v.edge,
	)
}

// TODO: docs
type primVtxHeap[T ds.Item] []*primVtx[T]

func (h primVtxHeap[T]) Len() int           { return len(h) }
func (h primVtxHeap[T]) Less(i, j int) bool { return h[i].wt < h[j].wt }
func (h primVtxHeap[T]) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]

	h[i].index = i
	h[j].index = j
}

func (h *primVtxHeap[T]) Push(x any) {
	*h = append(*h, x.(*primVtx[T]))
}

func (h *primVtxHeap[T]) Pop() any {
	n := len(*h)
	x := (*h)[n-1]

	(*h)[n-1] = nil
	*h = (*h)[:n-1]

	x.in = false
	x.index = -1

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
			ptr:   vtx.Ptr,
			wt:    wt,
			in:    true,
			index: i,
		}

		vtxHeap = append(vtxHeap, pVtx)
		vtxMap[pVtx.ptr] = pVtx
	}

	heap.Init(&vtxHeap)

	mst := MST[T]{}

	for len(vtxHeap) != 0 {
		vtx := heap.Pop(&vtxHeap).(*primVtx[T])

		if math.IsInf(vtx.wt, 1) {
			return nil, errors.New("input graph is not connected")
		}

		if vtx.edge != nil {
			mst = append(mst, vtx.edge)
		}

		for _, e := range g.E[vtx.ptr] {
			dstVtx := vtxMap[e.Dst]

			if !dstVtx.in {
				continue
			}

			if e.Wt >= dstVtx.wt {
				continue
			}

			dstVtx.edge = e
			dstVtx.wt = e.Wt

			heap.Fix(&vtxHeap, dstVtx.index)
		}
	}

	return mst, nil
}
