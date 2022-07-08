package ds

import "errors"

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
	/*
		VertMap maps items to the position of their vertex in the vertex list.
		Usually, it would be enough to map items to vertices directly, but
		since it is desireable for the graph to be traversed in the same order
		the vertices were inserted (for consistency  and presentation purposes),
		we map the item to the position of their vertex in the ordered list instead.
	*/
	VertMap map[*V]int

	/*
		Verts is the list of vertices in the graph. The list is ordered by insertion time,
		and the ordering that it provides is respected during internal traversals.
	*/
	Verts []*GraphVertex[V]

	/*
		Adj holds the adjacency lists (composed of edges) for the vertices in the graph.
		Note that traversing this map does not guarantee any particular ordering of the
		vertices; if insertion order is desired, iterate over Verts instead.
	*/
	Adj map[*V][]*GraphEdge[V]

	dir bool
}

func newGraph[V Item](dir bool) *Graph[V] {
	g := Graph[V]{}

	g.Verts = make([]*GraphVertex[V], 0)
	g.VertMap = make(map[*V]int)
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

// EmptyCopy creates an empty graph of the same kind.
func (g *Graph[V]) EmptyCopy() *Graph[V] {
	return newGraph[V](g.dir)
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
	_, ok := g.VertMap[v]
	return ok
}

// GetVertex fetches the vertex for the given data, if one exists in the graph.
func (g *Graph[V]) GetVertex(v *V) (*GraphVertex[V], bool) {
	idx, ok := g.VertMap[v]

	if !ok {
		return nil, false
	}

	return g.Verts[idx], true
}

// GetEdge fetches the edge from src to dst, if one exists in the graph.
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
	g.Verts = append(g.Verts, &GraphVertex[V]{Sat: v})
	g.VertMap[v] = len(g.Verts) - 1
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
	g.Adj[src] = append(g.Adj[src], &GraphEdge[V]{Src: src, Dst: dst, Wt: wt})
}

/*
AddWeightedEdge attempts to add a new weighted edge to the graph. If the graph is undirected,
the reverse edge is also added, if it does not already exist.
*/
func (g *Graph[V]) AddWeightedEdge(src, dst *V, wt float64) error {
	if src == nil || dst == nil {
		return errors.New("edge has nil components")
	}

	if g.Undirected() && src == dst {
		return errors.New("adding self-loop to undirected graph")
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
		return nil
	}

	if _, ok := g.GetEdge(dst, src); !ok {
		g.addWeightedEdge(dst, src, wt)
	}

	return nil
}

// AddEdge attempts to add a new unweighted edge to the graph.
func (g *Graph[V]) AddEdge(src, dst *V) error {
	return g.AddWeightedEdge(src, dst, 0)
}

// VertexCount calculates |V|, the number of vertices currently in the graph.
func (g *Graph[V]) VertexCount() int {
	return len(g.VertMap)
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

// Accept accepts a graph visitor, and guides its execution using double-dispatching.
func (g *Graph[V]) Accept(v GraphVisitor[V]) {
	v.VisitGraphStart(g)

	for _, vert := range g.Verts {
		vert.Accept(v)

		es, ok := g.Adj[vert.Sat]

		if !ok {
			continue
		}

		if es == nil {
			continue
		}

		for _, e := range es {
			e.Accept(v)
		}
	}

	v.VisitGraphEnd(g)
}

/*
Transpose creates a transpose of the graph: a new graph where all edges are reversed.
This is only true for directed graphs: undirected graphs will get a deep copy instead.
*/
func (g *Graph[V]) Transpose() *Graph[V] {
	res := g.EmptyCopy()

	// same order of insertion
	for _, vert := range g.Verts {
		res.AddVertex(vert.Sat)
	}

	// reverse the edges
	for _, es := range g.Adj {
		if es == nil {
			continue
		}

		for _, e := range es {
			res.AddWeightedEdge(e.Dst, e.Src, e.Wt)
		}
	}

	return res
}
