package ds

import "fmt"

/*
A GV represents a vertex in a graph.
*/
type GV[T Item] struct {
	Formattable

	/*
		Ptr holds a pointer to the satellite data of the vertex, which is data
		that should come along with the vertex wherever it goes.
	*/
	Ptr *T
}

// Label provides a label for the vertex, straight from its satellite data.
func (vert *GV[T]) Label() string {
	return (*vert.Ptr).Label()
}

func (vert GV[T]) String() string {
	return vert.Label()
}

// Accept accepts a graph visitor, and guides its execution using double-dispatching.
func (vert *GV[T]) Accept(v GraphVisitor[T]) {
	v.VisitVertex(vert)
}

/*
A GE represents a connection between two vertices in a graph, with an optional weight.
The directed/undirected nature of the edge is given by the graph that owns it.
*/
type GE[T Item] struct {
	Formattable

	/*
		Src is the source/tail of the edge (digraph).
		This edge contributes to the out-degree of Src (digraph).
	*/
	Src *T

	/*
		Dst is the destination/head of the edge (digraph).
		This edge contributes to the in-degree of Dst (digraph).
	*/
	Dst *T

	// Wt is the weight/cost associated with the edge.
	Wt float64
}

// Accept accepts a graph visitor, and guides its execution using double-dispatching.
func (e *GE[T]) Accept(v GraphVisitor[T]) {
	v.VisitEdge(e)
}

func (e GE[T]) String() string {
	return fmt.Sprintf(
		"%s -> %s <%.2f>",
		(*e.Src).Label(),
		(*e.Dst).Label(),
		e.Wt,
	)
}

/*
G implements a directed or undirected graph G = (V, E), using an adjacency list.

It has a space complexity of Θ(V + E), making use of pointers to prevent copying expensive
vertices during its operations, so the overall memory footprint should be bearable enough.

Due to the nature of adjacency lists, checking whether an edge (u, v) exists has O(E) time complexity.
For applications (and algorithms) that make heavy use of this operation, an adjacency matrix would be
a better fit (O(1) time complexity), with the trade-off being a worse space complexity of Θ(V²).
*/
type G[T Item] struct {
	/*
		V is the list of vertices in the graph. The list is ordered by insertion time,
		and the ordering that it provides is respected during internal traversals.
	*/
	V []*GV[T]

	/*
		E holds the adjacency lists (composed of edges) for the vertices in the graph.
		Note that traversing this map does not guarantee any particular ordering of the
		vertices; if insertion order is desired, iterate over Verts instead.
	*/
	E map[*T][]*GE[T]

	/*
		vMap maps items to the position of their vertex in the vertex list.
		Usually, it would be enough to map items to vertices directly, but
		since it is desirable for the graph to be traversed in the same order
		the vertices were inserted (for consistency and presentation purposes),
		we map the item to the position of their vertex in the ordered list instead.
	*/
	vMap map[*T]int

	// dir indicates whether the graph is directed.
	dir bool

	eCount int
	vCount int
}

func newGraph[T Item](dir bool) *G[T] {
	g := G[T]{}

	g.V = []*GV[T]{}
	g.E = map[*T][]*GE[T]{}

	g.vMap = map[*T]int{}
	g.dir = dir

	return &g
}

// NewDirectedGraph creates an empty directed graph.
func NewDirectedGraph[T Item]() *G[T] {
	return newGraph[T](true)
}

/// NewUndirectedGraph creates an empty undirected graph.
func NewUndirectedGraph[T Item]() *G[T] {
	return newGraph[T](false)
}

// EmptyCopy creates an empty graph of the same kind.
func (g *G[T]) EmptyCopy() *G[T] {
	return newGraph[T](g.dir)
}

// Directed checks whether or not the graph is directed.
func (g *G[T]) Directed() bool {
	return g.dir
}

// Directed checks whether or not the graph is undirected.
func (g *G[T]) Undirected() bool {
	return !g.dir
}

// VertexExists checks whether or not a given vertex exists in the graph.
func (g *G[T]) VertexExists(t *T) bool {
	_, ok := g.vMap[t]
	return ok
}

// GetVertex fetches the vertex for the given data, if one exists in the graph.
func (g *G[T]) GetVertex(t *T) (*GV[T], int, bool) {
	idx, ok := g.vMap[t]

	if !ok {
		return nil, -1, false
	}

	return g.V[idx], idx, true
}

// GetEdge fetches the edge from src to dst, if one exists in the graph.
func (g *G[T]) GetEdge(src *T, dst *T) (*GE[T], int, bool) {
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
func (g *G[T]) UnsafeAddVertex(t *T) *GV[T] {
	res := &GV[T]{Ptr: t}

	g.V = append(g.V, res)
	g.vMap[t] = len(g.V) - 1
	g.E[t] = nil
	g.vCount++

	return res
}

// AddVertex attempts to add whatever vertices are passed to the graph.
func (g *G[T]) AddVertex(t *T) (*GV[T], error) {
	if t == nil {
		return nil, ErrNilArg
	}

	if g.VertexExists(t) {
		return nil, ErrExists
	}

	return g.UnsafeAddVertex(t), nil
}

func (g *G[T]) removeVertex(t *T, idx int) {
	// remove the mapping
	delete(g.vMap, t)

	// remove the actual vertex
	g.V = RemoveFromPointersSlice(g.V, idx)

	// update the index of all copied vertices
	for i := idx; i < len(g.V); i++ {
		item := g.V[i].Ptr
		g.vMap[item] = i
	}

	g.vCount--
}

func (g *G[T]) removeVertexEdges(t *T) {
	g.eCount -= len(g.E[t])

	// remove every edge coming out of the vertex
	delete(g.E, t)

	// remove every edge arriving at the vertex
	for vert, es := range g.E {
		eIdx := -1

		for i, e := range es {
			// only one possible edge:
			// no multigraph support
			if e.Dst == t {
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
func (g *G[T]) RemoveVertex(t *T) error {
	_, idx, ok := g.GetVertex(t)

	if !ok {
		return ErrNotExists
	}

	g.removeVertexEdges(t)
	g.removeVertex(t, idx)

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
func (g *G[T]) UnsafeAddWeightedEdge(src, dst *T, wt float64) *GE[T] {
	res := &GE[T]{Src: src, Dst: dst, Wt: wt}
	g.E[src] = append(g.E[src], res)
	g.eCount++

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
func (g *G[T]) AddWeightedEdge(src, dst *T, wt float64) (*GE[T], error) {
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
func (g *G[T]) AddUnweightedEdge(src, dst *T) (*GE[T], error) {
	return g.AddWeightedEdge(src, dst, 0)
}

func (g *G[T]) removeEdge(src *T, idx int) {
	g.E[src] = RemoveFromPointersSlice(g.E[src], idx)
	g.eCount--
}

// RemoveEdge removes an edge from the graph, if it exists.
func (g *G[T]) RemoveEdge(src, dst *T) error {
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
func (g *G[T]) VertexCount() int {
	return g.vCount
}

// EdgeCount calculates |E|, the number of edges currently in the graph.
func (g *G[T]) EdgeCount() int {
	if g.Undirected() {
		return g.eCount / 2
	} else {
		return g.eCount
	}
}

// Accept accepts a graph visitor, and guides its execution using double-dispatching.
func (g *G[T]) Accept(v GraphVisitor[T]) {
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
func (g *G[T]) Transpose() (*G[T], error) {
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
