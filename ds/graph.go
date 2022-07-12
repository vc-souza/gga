package ds

/*
A GraphVertex represents a vertex in a graph.
*/
type GraphVertex[V Item] struct {
	Formattable

	/*
		Val holds satellite data for the vertex, which is data that should come along with the vertex everywhere it goes.
	*/
	Val *V
}

// Label provides a label for the vertex, straight from its satellite data.
func (vert *GraphVertex[V]) Label() string {
	return (*vert.Val).Label()
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
		the vertices were inserted (for consistency and presentation purposes),
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

	g.VertMap = make(map[*V]int)
	g.Verts = make([]*GraphVertex[V], 0)
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
func (g *Graph[V]) GetVertex(v *V) (*GraphVertex[V], int, bool) {
	idx, ok := g.VertMap[v]

	if !ok {
		return nil, -1, false
	}

	return g.Verts[idx], idx, true
}

// GetEdge fetches the edge from src to dst, if one exists in the graph.
func (g *Graph[V]) GetEdge(src *V, dst *V) (*GraphEdge[V], int, bool) {
	if src == nil || dst == nil {
		return nil, -1, false
	}

	es, ok := g.Adj[src]

	if !ok {
		return nil, -1, false
	}

	for i, e := range es {
		if e.Dst == dst {
			return e, i, true
		}
	}

	return nil, -1, false
}

/*
UnsafeAddVertex is the unsafe version of AddVertex, used by it after its validity checks.

Unlike the safe versions, no validation is performed, and a graph could easily become invalid:
- The same vertex could be added multiple times.
- A nil vertex could be added.

This method should only be called directly by a client if:
- They fully understand its possible negative implications
- They are positive that the graph being built is correct
- They need optimal performance while building the graph

Vertex checks are not nearly as bad as edge checks (just an additional O(1) check),
but if you're trying to squeeze every ounce of performance out of the graph building
process, this might help a bit.

If you have any doubts about using this version, use the safe one.
*/
func (g *Graph[V]) UnsafeAddVertex(v *V) *GraphVertex[V] {
	res := &GraphVertex[V]{Val: v}

	g.Verts = append(g.Verts, res)
	g.VertMap[v] = len(g.Verts) - 1
	g.Adj[v] = nil

	return res
}

// AddVertex attempts to add whatever vertices are passed to the graph.
func (g *Graph[V]) AddVertex(v *V) (*GraphVertex[V], error) {
	if v == nil {
		return nil, ErrNilArg
	}

	if g.VertexExists(v) {
		return nil, ErrExists
	}

	return g.UnsafeAddVertex(v), nil
}

func (g *Graph[V]) removeVertex(v *V, idx int) {
	// remove the mapping
	delete(g.VertMap, v)

	// remove the actual vertex
	g.Verts = RemoveFromPointersSlice(g.Verts, idx)

	// update the index of all copied vertices
	for i := idx; i < len(g.Verts); i++ {
		item := g.Verts[i].Val
		g.VertMap[item] = i
	}
}

func (g *Graph[V]) removeVertexEdges(v *V) {
	// remove every edge coming out of the vertex
	delete(g.Adj, v)

	// remove every edge arriving at the vertex
	for vert, es := range g.Adj {
		eIdx := -1

		for i, e := range es {
			// only one possible edge:
			// no multigraph support
			if e.Dst == v {
				eIdx = i
				break
			}
		}

		if eIdx == -1 {
			continue
		}

		g.removeEdge(vert, eIdx)
	}
}

// RemoveVertex removes a vertex from the graph, if it exists.
func (g *Graph[V]) RemoveVertex(v *V) error {
	_, idx, ok := g.GetVertex(v)

	if !ok {
		return ErrNotExists
	}

	g.removeVertexEdges(v)
	g.removeVertex(v, idx)

	return nil
}

/*
UnsafeAddWeightedEdge is the unsafe version of AddWeightedEdge/AddUnweightedEdge, used by them after their validity checks.

Unlike the safe versions, no validation is performed, and a graph could easily become invalid:
- Two vertices could have multiple edges connecting them (multigraphs are not supported).
- The src and dst vertices might not exist yet in the graph.
- Self-loops could be added to undirected graphs.
- Either src or dst (or both) could be nil.

This method should only be called directly by a client if:
- They fully understand its possible negative implications
- They are positive that the graph being built is correct
- They need optimal performance while building the graph

Since edge checks have O(E) time complexity, a very large graph that is also very dense could
cause performance issues while being built. If this is your use case and you have a valid
sequence of steps to build a graph, UnsafeAddWeightedEdge might be the method for you.

If you have any doubts about using this version, use the safe ones.
*/
func (g *Graph[V]) UnsafeAddWeightedEdge(src, dst *V, wt float64) *GraphEdge[V] {
	res := &GraphEdge[V]{Src: src, Dst: dst, Wt: wt}

	g.Adj[src] = append(g.Adj[src], res)

	return res
}

/*
AddWeightedEdge attempts to add a new weighted edge to the graph.
Several validity checks are performed, and extra work, like adding
a vertex that does not exist yet, is going to be performed for
the sake of ease of use.

In undirected graphs, if two vertices 'u' and 'v' are connected,
two edges need to be manually added: (u -> v) and (v -> u).

If you are trying to build a large, dense graph, have a sequence of operations
that creates a valid graph, and is running into performance issues, consider
using the UnsafeAddWeightedEdge method directly.
*/
func (g *Graph[V]) AddWeightedEdge(src, dst *V, wt float64) (*GraphEdge[V], error) {
	if src == nil || dst == nil {
		return nil, ErrNilArg
	}

	if g.Undirected() && src == dst {
		return nil, ErrInvalidLoop
	}

	g.AddVertex(src)
	g.AddVertex(dst)

	_, _, ok := g.GetEdge(src, dst)

	if ok {
		return nil, ErrExists
	}

	return g.UnsafeAddWeightedEdge(src, dst, wt), nil
}

/*
AddWeightedEdge attempts to add a new edge to the graph with weight 0.

If you are trying to build a large, dense graph, have a sequence of operations
that creates a valid graph, and is running into performance issues, consider
using the UnsafeAddWeightedEdge method directly.
*/
func (g *Graph[V]) AddUnweightedEdge(src, dst *V) (*GraphEdge[V], error) {
	return g.AddWeightedEdge(src, dst, 0)
}

func (g *Graph[V]) removeEdge(src *V, idx int) {
	g.Adj[src] = RemoveFromPointersSlice(g.Adj[src], idx)
}

// RemoveEdge removes an edge from the graph, if it exists.
func (g *Graph[V]) RemoveEdge(src, dst *V) error {
	_, idx, ok := g.GetEdge(src, dst)

	if !ok {
		return ErrNotExists
	}

	g.removeEdge(src, idx)

	if g.Directed() {
		return nil
	}

	// reverse edge, in undirected graphs
	if _, idx, ok := g.GetEdge(dst, src); ok {
		g.removeEdge(dst, idx)
	}

	return nil
}

// VertexCount calculates |V|, the number of vertices currently in the graph.
func (g *Graph[V]) VertexCount() int {
	return len(g.VertMap)
}

// EdgeCount calculates |E|, the number of edges currently in the graph.
func (g *Graph[V]) EdgeCount() int {
	res := 0

	for _, es := range g.Adj {
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

		es := g.Adj[vert.Val]

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
func (g *Graph[V]) Transpose() (*Graph[V], error) {
	if g.Undirected() {
		return nil, ErrUndefOp
	}

	res := g.EmptyCopy()

	// same order of insertion
	for _, vert := range g.Verts {
		res.AddVertex(vert.Val)
	}

	// reverse the edges
	for _, es := range g.Adj {
		for _, e := range es {
			res.AddWeightedEdge(e.Dst, e.Src, e.Wt)
		}
	}

	return res, nil
}
