package data

/*
A Vertex implementation can be used as a vertex in a gga graph.
No changes are made to the struct that implements this interface.
*/
type Vertex interface {
	VertexId() string
}

/*
An Edge represents a connection between a two vertices, with an optional weight.
The directed/undirected nature of the edge is given by the graph that owns it.
*/
type Edge[V Vertex] struct {
	/*
		Src is the source/tail of the edge (digraph).
		This edge contributes to the out-degree of Src (digraph).
	*/
	Src *V

	/*
		Dst is the destination/head of the edge (digraph).
		This edge contributes to the in-degree of Dst (digraph).
	*/
	Dst *V

	// Wt is the weight/cost associated with the edge.
	Wt float64
}

/*
Graph implements a directed or undirected graph G = (V, E), using an adjacency list.

It has a space complexity of Θ(V + E), making use of pointers to prevent copying expensive
vertices during its operations, so the overall memory footprint should be bearable enough.

Due to the nature of adjacency lists, checking whether an edge (u, v) exists has O(E) time complexity.
For applications (and algorithms) that make heavy use of this operation, an adjacency matrix would be
a better fit (O(1) time complexity), with the trade-off being a worse space complexity of Θ(V²).
*/
type Graph[V Vertex] struct {
	dir bool

	// Adj holds the adjacency lists for all vertices currently in the graph.
	Adj map[*V][]Edge[V]
}

func newGraph[V Vertex](dir bool) *Graph[V] {
	g := Graph[V]{}

	g.Adj = make(map[*V][]Edge[V])
	g.dir = dir

	return &g
}

// NewDirectedGraph creates an empty directed graph.
func NewDirectedGraph[V Vertex]() *Graph[V] {
	return newGraph[V](true)
}

/// NewUndirectedGraph creates an empty undirected graph.
func NewUndirectedGraph[V Vertex]() *Graph[V] {
	return newGraph[V](false)
}

// Directed checks whether or not the graph is directed.
func (g *Graph[V]) Directed() bool {
	return g.dir
}

// Directed checks whether or not the graph is undirected.
func (g *Graph[V]) Undirected() bool {
	return !g.dir
}

// VertexExists checks whether or not a given vertex exists in the graph.
func (g *Graph[V]) VertexExists(v *V) bool {
	_, ok := g.Adj[v]
	return ok
}

// GetEdge fetches an edge from src to dst, if one exists in the graph.
func (g *Graph[V]) GetEdge(src *V, dst *V) (*Edge[V], bool) {
	if src == nil || dst == nil {
		return nil, false
	}

	es, ok := g.Adj[src]

	if !ok {
		return nil, false
	}

	if es == nil {
		return nil, false
	}

	for _, e := range es {
		if e.Dst == dst {
			return &e, true
		}
	}

	return nil, false
}

// AddVertex attempts to add whatever vertices are passed to the graph.
func (g *Graph[V]) AddVertex(v ...*V) {
	for _, ver := range v {
		if ver == nil {
			continue
		}

		if g.VertexExists(ver) {
			continue
		}

		g.Adj[ver] = nil
	}
}

/*
AddWeightedEdge attempts to add a new weighted edge to the graph. If the graph is undirected,
the reverse edge is also added, if it does not already exist.
*/
func (g *Graph[V]) AddWeightedEdge(src, dst *V, wt float64) {
	if src == nil || dst == nil {
		return
	}

	if _, ok := g.GetEdge(src, dst); !ok {
		g.Adj[src] = append(g.Adj[src], Edge[V]{Src: src, Dst: dst, Wt: wt})
	}

	if g.Directed() {
		return
	}

	if _, ok := g.GetEdge(dst, src); !ok {
		g.Adj[dst] = append(g.Adj[dst], Edge[V]{Src: dst, Dst: src, Wt: wt})
	}
}

// AddEdge attempts to add a new unweighted edge to the graph.
func (g *Graph[V]) AddEdge(src, dst *V) {
	g.AddWeightedEdge(src, dst, 0)
}

// VertexCount calculates |V|, the number of vertices currently in the graph.
func (g *Graph[V]) VertexCount() int {
	return len(g.Adj)
}

// EdgeCount calculates |E|, the number of edges currently in the graph.
func (g *Graph[V]) EdgeCount() int {
	res := 0

	for _, es := range g.Adj {
		if es == nil {
			continue
		}

		for range es {
			res++
		}
	}

	if g.Undirected() {
		res = res / 2
	}

	return res
}
