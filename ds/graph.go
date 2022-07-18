package ds

/*
A GV represents a vertex in a graph.
*/
type GV[V Item] struct {
	Formattable

	/*
		Ptr holds a pointer to the satellite data of the vertex, which is data
		that should come along with the vertex wherever it goes.
	*/
	Ptr *V
}

// Label provides a label for the vertex, straight from its satellite data.
func (vert *GV[V]) Label() string {
	return (*vert.Ptr).Label()
}

// Accept accepts a graph visitor, and guides its execution using double-dispatching.
func (vert *GV[V]) Accept(v GraphVisitor[V]) {
	v.VisitVertex(vert)
}

/*
A GE represents a connection between two vertices in a graph, with an optional weight.
The directed/undirected nature of the edge is given by the graph that owns it.
*/
type GE[V Item] struct {
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
func (e *GE[V]) Accept(v GraphVisitor[V]) {
	v.VisitEdge(e)
}

/*
G implements a directed or undirected graph G = (V, E), using an adjacency list.

It has a space complexity of Θ(V + E), making use of pointers to prevent copying expensive
vertices during its operations, so the overall memory footprint should be bearable enough.

Due to the nature of adjacency lists, checking whether an edge (u, v) exists has O(E) time complexity.
For applications (and algorithms) that make heavy use of this operation, an adjacency matrix would be
a better fit (O(1) time complexity), with the trade-off being a worse space complexity of Θ(V²).
*/
type G[V Item] struct {
	/*
		V is the list of vertices in the graph. The list is ordered by insertion time,
		and the ordering that it provides is respected during internal traversals.
	*/
	V []*GV[V]

	/*
		E holds the adjacency lists (composed of edges) for the vertices in the graph.
		Note that traversing this map does not guarantee any particular ordering of the
		vertices; if insertion order is desired, iterate over Verts instead.
	*/
	E map[*V][]*GE[V]

	/*
		vMap maps items to the position of their vertex in the vertex list.
		Usually, it would be enough to map items to vertices directly, but
		since it is desirable for the graph to be traversed in the same order
		the vertices were inserted (for consistency and presentation purposes),
		we map the item to the position of their vertex in the ordered list instead.
	*/
	vMap map[*V]int

	// dir indicates whether the graph is directed.
	dir bool
}

func newGraph[V Item](dir bool) *G[V] {
	g := G[V]{}

	g.V = make([]*GV[V], 0)
	g.E = make(map[*V][]*GE[V])

	g.vMap = make(map[*V]int)
	g.dir = dir

	return &g
}

// NewDirectedGraph creates an empty directed graph.
func NewDirectedGraph[V Item]() *G[V] {
	return newGraph[V](true)
}

/// NewUndirectedGraph creates an empty undirected graph.
func NewUndirectedGraph[V Item]() *G[V] {
	return newGraph[V](false)
}

// EmptyCopy creates an empty graph of the same kind.
func (g *G[V]) EmptyCopy() *G[V] {
	return newGraph[V](g.dir)
}

// Directed checks whether or not the graph is directed.
func (g *G[V]) Directed() bool {
	return g.dir
}

// Directed checks whether or not the graph is undirected.
func (g *G[V]) Undirected() bool {
	return !g.dir
}

// VertexExists checks whether or not a given vertex exists in the graph.
func (g *G[V]) VertexExists(v *V) bool {
	_, ok := g.vMap[v]
	return ok
}

// GetVertex fetches the vertex for the given data, if one exists in the graph.
func (g *G[V]) GetVertex(v *V) (*GV[V], int, bool) {
	idx, ok := g.vMap[v]

	if !ok {
		return nil, -1, false
	}

	return g.V[idx], idx, true
}

// GetEdge fetches the edge from src to dst, if one exists in the graph.
func (g *G[V]) GetEdge(src *V, dst *V) (*GE[V], int, bool) {
	if src == nil || dst == nil {
		return nil, -1, false
	}

	es, ok := g.E[src]

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
func (g *G[V]) UnsafeAddVertex(v *V) *GV[V] {
	res := &GV[V]{Ptr: v}

	g.V = append(g.V, res)
	g.vMap[v] = len(g.V) - 1
	g.E[v] = nil

	return res
}

// AddVertex attempts to add whatever vertices are passed to the graph.
func (g *G[V]) AddVertex(v *V) (*GV[V], error) {
	if v == nil {
		return nil, ErrNilArg
	}

	if g.VertexExists(v) {
		return nil, ErrExists
	}

	return g.UnsafeAddVertex(v), nil
}

func (g *G[V]) removeVertex(v *V, idx int) {
	// remove the mapping
	delete(g.vMap, v)

	// remove the actual vertex
	g.V = RemoveFromPointersSlice(g.V, idx)

	// update the index of all copied vertices
	for i := idx; i < len(g.V); i++ {
		item := g.V[i].Ptr
		g.vMap[item] = i
	}
}

func (g *G[V]) removeVertexEdges(v *V) {
	// remove every edge coming out of the vertex
	delete(g.E, v)

	// remove every edge arriving at the vertex
	for vert, es := range g.E {
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
func (g *G[V]) RemoveVertex(v *V) error {
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
func (g *G[V]) UnsafeAddWeightedEdge(src, dst *V, wt float64) *GE[V] {
	res := &GE[V]{Src: src, Dst: dst, Wt: wt}

	g.E[src] = append(g.E[src], res)

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
func (g *G[V]) AddWeightedEdge(src, dst *V, wt float64) (*GE[V], error) {
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
func (g *G[V]) AddUnweightedEdge(src, dst *V) (*GE[V], error) {
	return g.AddWeightedEdge(src, dst, 0)
}

func (g *G[V]) removeEdge(src *V, idx int) {
	g.E[src] = RemoveFromPointersSlice(g.E[src], idx)
}

// RemoveEdge removes an edge from the graph, if it exists.
func (g *G[V]) RemoveEdge(src, dst *V) error {
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
func (g *G[V]) VertexCount() int {
	return len(g.vMap)
}

// EdgeCount calculates |E|, the number of edges currently in the graph.
func (g *G[V]) EdgeCount() int {
	res := 0

	for _, es := range g.E {
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
func (g *G[V]) Accept(v GraphVisitor[V]) {
	v.VisitGraphStart(g)

	for _, vert := range g.V {
		vert.Accept(v)

		es := g.E[vert.Ptr]

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
func (g *G[V]) Transpose() (*G[V], error) {
	if g.Undirected() {
		return nil, ErrUndefOp
	}

	res := g.EmptyCopy()

	// same order of insertion
	for _, vert := range g.V {
		res.AddVertex(vert.Ptr)
	}

	// reverse the edges
	for _, es := range g.E {
		for _, e := range es {
			res.AddWeightedEdge(e.Dst, e.Src, e.Wt)
		}
	}

	return res, nil
}
