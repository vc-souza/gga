package algo

import (
	"container/heap"
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
type MSTAlgo func(*ds.G) (MST, error)

/*
An MST holds the edges of a minimum spanning tree
of an undirected graph with weighted edges.
*/
type MST []ds.GE

/*
primVtx is an auxiliary type used only by MSTPrim to keep track
of the status of each vertex in the heap.
*/
type primVtx struct {
	// id holds a reference to the original vertex.
	id int

	// wt holds the weight of the best edge found so far that can connect the vertex to the MST.
	wt float64

	// edge holds the best edge found so far that can connect the vertex to the MST.
	edge *ds.GE

	// in tells if the vertex is still in the heap.
	in bool

	// index stores the index of the vertex, in the heap.
	index int
}

// primVtxHeap implements heap.Interface to provide min-heap features for *primVtx[T] values.
type primVtxHeap []*primVtx

func (h primVtxHeap) Len() int           { return len(h) }
func (h primVtxHeap) Less(i, j int) bool { return h[i].wt < h[j].wt }
func (h primVtxHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
	h[i].index = i
	h[j].index = j
}

func (h *primVtxHeap) Push(x any) {
	*h = append(*h, x.(*primVtx))
}

func (h *primVtxHeap) Pop() any {
	n := len(*h)
	x := (*h)[n-1]

	(*h)[n-1] = nil
	*h = (*h)[:n-1]

	x.in = false
	x.index = -1

	return x
}

/*
MSTPrim implements Prim's algorithm for finding a minimum spanning tree
of an undirected graph with weighted edges. One extra restriction is that
the graph needs to be connected: if not, the algorithm will not find the
minimum spanning forest, and an error will be returned.

This is a greedy algorithm that can start from any source vertex (this
particular implementation always starts from the first vertex added to
the graph) and that expands the current MST subset one vertex at a time,
after each iteration.

The greedy-choice property is used by maintaining a min-heap containing
every vertex that has not been added to the MST subset, using the lowest
known edge weight to reach that vertex from the MST subset as the heap key.

At the start of each iteration, the vertex v with the smallest key is
extracted from the heap and added to the MST subset (greedy choice,
locally optimal). Then the adjacency list of v is examined, and every
vertex u that is still in the heap and whose edge (v, u) has a weight
that is smaller than the key of u is updated in the heap, which is
then fixed to keep its heap property.

After the heap is emptied, the algorithm is guaranteed to have computed
an MST of the original graph (globally optimal solution), but only if
the graph is connected.

Expectations:
	- The graph is correctly built.
	- The graph is undirected.

Complexity:
	- Time:  O(E log V)
	- Space: Θ(V).
*/
func MSTPrim(g *ds.G) (MST, error) {
	if g.Directed() {
		return nil, ds.ErrDirected
	}

	vtxHeap := make(primVtxHeap, g.VertexCount())
	att := make([]primVtx, g.VertexCount())

	for v := range g.V {
		var wt float64

		// source
		if v == 0 {
			wt = 0
		} else {
			wt = math.Inf(1)
		}

		att[v].id = v
		att[v].wt = wt
		att[v].in = true
		att[v].index = v

		vtxHeap[v] = &att[v]
	}

	heap.Init(&vtxHeap)

	mst := MST{}

	for len(vtxHeap) != 0 {
		v := heap.Pop(&vtxHeap).(*primVtx)

		if math.IsInf(v.wt, 1) {
			return nil, ds.ErrDisconnected
		}

		if v.edge != nil {
			mst = append(mst, *v.edge)
		}

		for i := range g.V[v.id].E {
			edge := &g.V[v.id].E[i]

			if !att[edge.Dst].in {
				continue
			}

			if edge.Wt >= att[edge.Dst].wt {
				continue
			}

			att[edge.Dst].edge = edge
			att[edge.Dst].wt = edge.Wt

			heap.Fix(&vtxHeap, att[edge.Dst].index)
		}
	}

	return mst, nil
}

/*
MSTKruskal implements Kruskal's algorithm for finding a minimum spanning tree
of an undirected graph with weighted edges.

This is a greedy algorithm that applies the greedy-choice property by first
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
func MSTKruskal(g *ds.G) (MST, error) {
	if g.Directed() {
		return nil, ds.ErrDirected
	}

	edges := make([]ds.GE, g.EdgeCount())

	// By iterating over G.V and adding edges using their original
	// insertion order, we can guarantee that every call of the
	// algorithm on the same graph always yields the same MST,
	// since multiple MSTs might exist for the same graph.
	for i, eIdx := 0, 0; i < len(g.V); i++ {
		copy(edges[eIdx:], g.V[i].E)
		eIdx += len(g.V[i].E)
	}

	// By using a stable sorting algorithm to sort the sequence
	// of edges in O(E log E) time, we make sure that if a tie
	// happens between edges of same weight, the original insertion
	// order is respected, and a consistent result is achieved.
	sort.Stable(ds.ByEdgeWeight(edges))

	d := ds.NewDSet[int]()

	for i := range g.V {
		d.MakeSet(i)
	}

	max := g.VertexCount() - 1
	mst := MST{}

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
