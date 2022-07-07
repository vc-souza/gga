package ds

/*
TODO: docs
*/
type GraphVertex[V Item] struct {
	Formattable

	// TODO: docs
	Satellite *V
}

func (vert *GraphVertex[V]) Accept(v GraphVisitor[V]) {
	v.VisitVertex(vert)
	v.VisitFormattingAttrs(vert.FmtAttrs)
}

/*
A GraphEdge represents a connection between a two vertices, with an optional weight.
The directed/undirected nature of the edge is given by the graph that owns it.
*/
type GraphEdge[V Item] struct {
	Formattable

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

// TODO: docs
func (e *GraphEdge[V]) Accept(v GraphVisitor[V]) {
	v.VisitEdge(e)
	v.VisitFormattingAttrs(e.FmtAttrs)
}

/*
Graph implements a directed or undirected graph G = (V, E), using an adjacency list.

It has a space complexity of Θ(V + E), making use of pointers to prevent copying expensive
vertices during its operations, so the overall memory footprint should be bearable enough.

Due to the nature of adjacency lists, checking whether an edge (u, v) exists has O(E) time complexity.
For applications (and algorithms) that make heavy use of this operation, an adjacency matrix would be
a better fit (O(1) time complexity), with the trade-off being a worse space complexity of Θ(V²).
*/
type Graph[V Item] struct {
	Formattable

	dir bool

	// TODO: docs
	Verts map[*V]*GraphVertex[V]

	// Adj holds the adjacency lists for all vertices currently in the graph.
	Adj map[*V][]*GraphEdge[V]
}

func newGraph[V Item](dir bool) *Graph[V] {
	g := Graph[V]{}

	g.Verts = make(map[*V]*GraphVertex[V])
	g.Adj = make(map[*V][]*GraphEdge[V])
	g.dir = dir

	return &g
}

// NewDirectedGraph creates an empty directed graph.
func NewDirectedGraph[V Item]() *Graph[V] {
	return newGraph[V](true)
}

/// NewUndirectedGraph creates an empty undirected graph.
func NewUndirectedGraph[V Item]() *Graph[V] {
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
	_, ok := g.Verts[v]
	return ok
}

// GetEdge fetches an edge from src to dst, if one exists in the graph.
func (g *Graph[V]) GetEdge(src *V, dst *V) (*GraphEdge[V], bool) {
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
			return e, true
		}
	}

	return nil, false
}

func (g *Graph[V]) addVertex(v *V) {
	g.Verts[v] = &GraphVertex[V]{Satellite: v}
	g.Adj[v] = nil
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

		g.addVertex(ver)
	}
}

func (g *Graph[V]) addWeightedEdge(src, dst *V, wt float64) {
	if !g.VertexExists(src) {
		g.addVertex(src)
	}

	g.Adj[src] = append(g.Adj[src], &GraphEdge[V]{Src: src, Dst: dst, Wt: wt})
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
		g.addWeightedEdge(src, dst, wt)
	}

	if g.Directed() {
		return
	}

	if _, ok := g.GetEdge(dst, src); !ok {
		g.addWeightedEdge(dst, src, wt)
	}
}

// AddEdge attempts to add a new unweighted edge to the graph.
func (g *Graph[V]) AddEdge(src, dst *V) {
	g.AddWeightedEdge(src, dst, 0)
}

// VertexCount calculates |V|, the number of vertices currently in the graph.
func (g *Graph[V]) VertexCount() int {
	return len(g.Verts)
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

// TODO: docs (no other accepts? no class hierarchy!)
func (g *Graph[V]) Accept(v GraphVisitor[V]) {
	v.VisitGraphStart(g)
	v.VisitFormattingAttrs(g.FmtAttrs)

	for vp, es := range g.Adj {
		v.VisitVertex(g.Verts[vp])

		if es == nil {
			continue
		}

		for _, e := range es {
			v.VisitEdge(e)
		}
	}

	v.VisitGraphEnd(g)
}
