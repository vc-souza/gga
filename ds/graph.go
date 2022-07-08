package ds

/*
A GraphVertex represents a vertex in a graph.
*/
type GraphVertex[V Item] struct {
	Formattable

	/*
		Sat holds satellite data for the vertex, which is data that should come along with the vertex everywhere it goes.
	*/
	Sat *V
}

// Label provides a label for the vertex, straight from its satellite data.
func (vert *GraphVertex[V]) Label() string {
	return (*vert.Sat).Label()
}

// Accept accepts a graph visitor, and guides its execution using double-dispatching.
func (vert *GraphVertex[V]) Accept(v GraphVisitor[V]) {
	v.VisitVertex(vert)
}

/*
A GraphEdge represents a connection between two vertices in a graph, with an optional weight.
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

// Accept accepts a graph visitor, and guides its execution using double-dispatching.
func (e *GraphEdge[V]) Accept(v GraphVisitor[V]) {
	v.VisitEdge(e)
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

	verts map[*V]*GraphVertex[V]
	adj   map[*V][]*GraphEdge[V]
	dir   bool
}

func newGraph[V Item](dir bool) *Graph[V] {
	g := Graph[V]{}

	g.verts = make(map[*V]*GraphVertex[V])
	g.adj = make(map[*V][]*GraphEdge[V])
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
	_, ok := g.verts[v]
	return ok
}

// GetVertex fetches the vertex for the given data, if one exists in the graph.
func (g *Graph[V]) GetVertex(v *V) (*GraphVertex[V], bool) {
	vert, ok := g.verts[v]
	return vert, ok
}

// GetEdge fetches the edge from src to dst, if one exists in the graph.
func (g *Graph[V]) GetEdge(src *V, dst *V) (*GraphEdge[V], bool) {
	if src == nil || dst == nil {
		return nil, false
	}

	es, ok := g.adj[src]

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
	g.verts[v] = &GraphVertex[V]{Sat: v}
	g.adj[v] = nil
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
	g.adj[src] = append(g.adj[src], &GraphEdge[V]{Src: src, Dst: dst, Wt: wt})
}

/*
AddWeightedEdge attempts to add a new weighted edge to the graph. If the graph is undirected,
the reverse edge is also added, if it does not already exist.
*/
func (g *Graph[V]) AddWeightedEdge(src, dst *V, wt float64) {
	if src == nil || dst == nil {
		return
	}

	if !g.VertexExists(src) {
		g.addVertex(src)
	}

	if !g.VertexExists(dst) {
		g.addVertex(dst)
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
	return len(g.verts)
}

// EdgeCount calculates |E|, the number of edges currently in the graph.
func (g *Graph[V]) EdgeCount() int {
	res := 0

	for _, es := range g.adj {
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

// Accept accepts a graph visitor, and guides its execution using double-dispatching.
func (g *Graph[V]) Accept(v GraphVisitor[V]) {
	v.VisitGraphStart(g)

	for vp, es := range g.adj {
		g.verts[vp].Accept(v)

		if es == nil {
			continue
		}

		for _, e := range es {
			e.Accept(v)
		}
	}

	v.VisitGraphEnd(g)
}
