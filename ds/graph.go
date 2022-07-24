package ds

import (
	"fmt"
	"strings"
)

/*
A GE represents an edge in a graph, which is a connection between two vertices,
with an optional weight. The directed/undirected nature of the edge is given by
the graph that owns it.
*/
type GE struct {
	Formattable

	/*
		Src is the id of the vertex at the source / tail of the edge.
		This edge contributes to the out-degree of that vertex.
	*/
	Src int

	/*
		Dst is the id of the vertex at the destination / head of the edge.
		This edge contributes to the in-degree of that vertex.
	*/
	Dst int

	// Wt is the weight/cost associated with the edge.
	Wt float64

	/*
		Index is the position of this edge in Src's adjacency list.
		This information is useful when passing around copies of the edge.
	*/
	Index int
}

func (e *GE) String() string {
	return fmt.Sprintf(
		"@%d %d -> %d <%.2f>",
		e.Index,
		e.Src,
		e.Dst,
		e.Wt,
	)
}

// A GV represents a vertex in a graph.
type GV struct {
	Formattable

	// Item is the satellite data of the vertex, coming along with it wherever it goes.
	Item Item

	// Index is the index of the vertex in the list of vertices of the graph.
	Index int

	// E is the adjacency list for the vertex, listing all edges that the vertex as their source.
	E []GE
}

func (v *GV) Label() string {
	return v.Item.Label()
}

func (v *GV) String() string {
	b := strings.Builder{}
	es := []string{}

	b.WriteString(fmt.Sprintf("Vertex '%s' @%d adj [", v.Label(), v.Index))

	for i := range v.E {
		es = append(es, v.E[i].String())
	}

	b.WriteString(strings.Join(es, ", "))
	b.WriteString("]\n")

	return b.String()
}

/*
G implements a directed or undirected graph G = (V, E), using adjacency lists
to achieve linear space complexity: Θ(V + E).

Due to the nature of adjacency lists, checking whether an edge (u, v) exists
has O(E) time complexity. For applications that make heavy use of this operation,
an adjacency matrix would be a better fit (O(1) time complexity), with the
trade-off being a worse space complexity of Θ(V²).
*/
type G struct {
	/*
		V is the list of vertices in the graph, ordered by vertex insertion time,
		an ordering that is respected during any internal traversals.
	*/
	V []GV

	sat    map[Item]int
	dir    bool
	eCount int
	vCount int
}

func newG(dir bool) *G {
	g := &G{}

	g.V = make([]GV, 0)

	g.sat = map[Item]int{}
	g.dir = dir

	return g
}

// NewGraph creates a new undirected graph.
func NewGraph() *G {
	return newG(false)
}

// NewDigraph creates a new directed graph.
func NewDigraph() *G {
	return newG(true)
}

func (g *G) String() string {
	b := strings.Builder{}

	if g.Undirected() {
		b.WriteString("Graph")
	} else {
		b.WriteString("Digraph")
	}

	b.WriteString(fmt.Sprintf(" |V| = %d", g.VertexCount()))
	b.WriteString(fmt.Sprintf(" |E| = %d\n", g.EdgeCount()))

	for v := range g.V {
		b.WriteString(g.V[v].String())
	}

	return b.String()
}

// Directed checks whether or not the graph is directed.
func (g *G) Directed() bool {
	return g.dir
}

// Undirected checks whether or not the graph is undirected.
func (g *G) Undirected() bool {
	return !g.dir
}

// VertexCount calculates the size of the set of vertices, |V|, in O(1) time.
func (g *G) VertexCount() int {
	return g.vCount
}

// GetVertexIndex retrieves the index of the vertex associated with the given Item.
func (g *G) GetVertexIndex(i Item) (int, bool) {
	idx, ok := g.sat[i]
	return idx, ok
}

// AddVertex adds a new vertex to the graph, associated with the given Item.
func (g *G) AddVertex(i Item) (int, error) {
	if _, ok := g.sat[i]; ok {
		return 0, ErrExists
	}

	g.V = append(g.V, GV{Item: i})

	idx := len(g.V) - 1

	g.V[idx].Index = idx
	g.sat[i] = idx

	g.vCount++

	return idx, nil
}

// RemoveVertex removes the vertex associated with the given Item, along with any edges incident on it.
func (g *G) RemoveVertex(i Item) error {
	iDel, ok := g.sat[i]

	if !ok {
		return ErrNoVtx
	}

	fixEdges := func() {
		for v := range g.V {
			if v == iDel {
				g.eCount -= len(g.V[v].E)
				continue
			}

			remove := []int{}
			shifts := 0

			for e := range g.V[v].E {
				edge := &g.V[v].E[e]

				// found a edge that arrives at the vertex
				// being removed: schedule it for removal
				// and increase the amount of shits that
				// later edges will be subjected to
				if edge.Dst == iDel {
					remove = append(remove, e)
					shifts++
					continue
				}

				// anticipating the removal of the vertex,
				// decrease any references to vertices that
				// will end up being shifted to the left
				if edge.Src > iDel {
					edge.Src--
				}

				// anticipating the removal of the vertex,
				// decrease any references to vertices that
				// will end up being shifted to the left
				if edge.Dst > iDel {
					edge.Dst--
				}

				// anticipating the removal of edges,
				// decrease the index of the current
				// edge by the number of edges that
				// will be deleted so far
				edge.Index -= shifts
			}

			for _, eIdx := range remove {
				Cut(&g.V[v].E, eIdx)
				g.eCount--
			}
		}
	}

	deleteVertex := func() {
		Cut(&g.V, iDel)
		delete(g.sat, i)
		g.vCount--
	}

	fixVertices := func() {
		for i := iDel; i < len(g.V); i++ {
			g.sat[g.V[i].Item] = i
			g.V[i].Index = i
		}
	}

	fixEdges()
	deleteVertex()
	fixVertices()

	return nil
}

// EdgeCount calculates the size of the set of edges, |E|, in O(1) time.
func (g *G) EdgeCount() int {
	return g.eCount
}

// GetEdgeIndex retrieves the index(es) of the edge associated with the given Items.
func (g *G) GetEdgeIndex(src Item, dst Item) (int, int, bool) {
	iSrc, ok := g.GetVertexIndex(src)

	if !ok {
		return 0, 0, false
	}

	iDst, ok := g.sat[dst]

	if !ok {
		return 0, 0, false
	}

	for e := range g.V[iSrc].E {
		if g.V[iSrc].E[e].Dst == iDst {
			return iSrc, e, true
		}
	}

	return 0, 0, false
}

// AddEdge adds a new weighted edge between the given Items, but only if their vertices have already been added.
func (g *G) AddEdge(src Item, dst Item, wt float64) (int, int, error) {
	if g.Undirected() && src == dst {
		return 0, 0, ErrInvLoop
	}

	iSrc, ok := g.GetVertexIndex(src)

	if !ok {
		return 0, 0, ErrNoVtx
	}

	iDst, ok := g.GetVertexIndex(dst)

	if !ok {
		return 0, 0, ErrNoVtx
	}

	for e := range g.V[iSrc].E {
		if g.V[iSrc].E[e].Dst == iDst {
			return 0, 0, ErrExists
		}
	}

	g.V[iSrc].E = append(
		g.V[iSrc].E,
		GE{
			Index: len(g.V[iSrc].E),
			Src:   iSrc,
			Dst:   iDst,
			Wt:    wt,
		},
	)

	g.eCount++

	return iSrc, len(g.V[iSrc].E) - 1, nil
}

// RemoveEdge removes the edge associated with the given Items.
func (g *G) RemoveEdge(src Item, dst Item) error {
	vIdx, idx, ok := g.GetEdgeIndex(src, dst)

	if !ok {
		return ErrNoEdge
	}

	Cut(&g.V[vIdx].E, idx)

	for i := idx; i < len(g.V[vIdx].E); i++ {
		g.V[vIdx].E[i].Index--
	}

	g.eCount--

	return nil
}

// Accept accepts a graph visitor, and guides its execution using double-dispatching.
func (g G) Accept(vis GraphVisitor) {
	vis.VisitGraphStart(g)

	for v := range g.V {
		vis.VisitVertex(g, g.V[v])

		for e := range g.V[v].E {
			vis.VisitEdge(g, g.V[v].E[e])
		}
	}

	vis.VisitGraphEnd(g)
}

// Transpose creates a transpose graph for a directed graph, where all original edges are reversed.
func Transpose(g *G) (*G, error) {
	if g.Undirected() {
		return nil, ErrUndirected
	}

	res := NewDigraph()

	for v := range g.V {
		res.AddVertex(g.V[v].Item)
	}

	for v := range g.V {
		for _, edge := range g.V[v].E {
			res.AddEdge(
				g.V[edge.Dst].Item,
				g.V[edge.Src].Item,
				edge.Wt,
			)
		}
	}

	return res, nil
}
