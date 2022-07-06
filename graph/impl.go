package graph

// An Identifiable
type Identifiable interface {
	Id() string
}

// TODO: docs
type Edge[T Identifiable] struct {
	// TODO: docs
	Src *T
	// TODO: docs
	Dst *T
	// TODO: docs
	Wt float64
}

// TODO: docs
type Graph[T Identifiable] struct {
	dir bool

	// TODO: docs
	Adj map[*T][]Edge[T]
}

func newGraph[T Identifiable](dir bool) *Graph[T] {
	g := Graph[T]{}

	g.Adj = make(map[*T][]Edge[T])
	g.dir = dir

	return &g
}

// TODO: docs
func NewDirectedGraph[T Identifiable]() *Graph[T] {
	return newGraph[T](true)
}

// TODO: docs
func NewUndirectedGraph[T Identifiable]() *Graph[T] {
	return newGraph[T](false)
}

// TODO: docs
func (g *Graph[T]) Directed() bool {
	return g.dir
}

// TODO: docs
func (g *Graph[T]) VertexExists(v *T) bool {
	_, ok := g.Adj[v]
	return ok
}

// TODO: docs
func (g *Graph[T]) AddVertex(v ...*T) {
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

// TODO: docs
func (g *Graph[T]) EdgeExists(src *T, dst *T) bool {
	es, ok := g.Adj[src]

	if !ok {
		return false
	}

	if es == nil {
		return false
	}

	for _, e := range es {
		if e.Dst == dst {
			return true
		}
	}

	return false
}

// TODO: docs
func (g *Graph[T]) AddWeightedEdge(src, dst *T, wt float64) {
	if !g.EdgeExists(src, dst) {
		g.Adj[src] = append(g.Adj[src], Edge[T]{Src: src, Dst: dst, Wt: wt})
	}

	if !g.Directed() && !g.EdgeExists(dst, src) {
		g.Adj[dst] = append(g.Adj[dst], Edge[T]{Src: dst, Dst: src, Wt: wt})
	}
}

// TODO: docs
func (g *Graph[T]) AddEdge(src, dst *T) {
	g.AddWeightedEdge(src, dst, 0)
}

// TODO: docs
func (g *Graph[T]) CountVertices() int {
	return len(g.Adj)
}

// TODO: docs
func (g *Graph[T]) CountEdges() int {
	res := 0

	for _, es := range g.Adj {
		if es == nil {
			continue
		}

		for range es {
			res++
		}
	}

	if !g.Directed() {
		res = res / 2
	}

	return res
}
