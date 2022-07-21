package algo

import "github.com/vc-souza/gga/ds"

/*
CCAlgo describes the signature of an algorithm that can discover all
connected components in an undirected graph. If such an algorithm
is called on a directed graph, the ds.ErrUndefOp error is returned.
*/
type CCAlgo[T ds.Item] func(*ds.G[T]) ([]CC[T], error)

// A CC holds the vertices in a connected component of an undirected graph.
type CC[T ds.Item] []*T

/*
CCDFS implements an algorithm for finding the connected components of an undirected graph
by using a single DFS that returns each DF tree in the DF forest as a connected component.

The DFS approach is better suited for static graphs: when the sets of vertices and edges
do not change over time. In this scenario, CCDFS has an asymptotically better time
complexity (linear) than the disjoint-set implementation, CCUnionFind (superlinear).

However, if the graph is dynamic, CCUnionFind will do a better job over time, since CCDFS
would need to be executed every time the graph changes - Θ(V + E) -, while CCUnionFind
would only need to be executed once - O((V + E) α(V)), amortized -, and its disjoint-set
data structure would need to be updated after every graph change, with each disjoint-set
operation taking O(α(V)) amortized time.

Expectations:
	- The graph is correctly built.
	- The graph is undirected.

Complexity:
	- Time:  Θ(V + E)
	- Space: Θ(V)
*/
func CCDFS[T ds.Item](g *ds.G[T]) ([]CC[T], error) {
	if g.Directed() {
		return nil, ds.ErrUndefOp
	}

	var visit func(*T)
	var cc *CC[T]

	visited := map[*T]bool{}
	ccs := []CC[T]{}

	for v := range g.E {
		visited[v] = false
	}

	visit = func(vtx *T) {
		visited[vtx] = true

		for _, e := range g.E[vtx] {
			if !visited[e.Dst] {
				visit(e.Dst)
			}
		}

		*cc = append(*cc, vtx)
	}

	for _, vert := range g.V {
		if visited[vert.Ptr] {
			continue
		}

		cc = &CC[T]{}

		visit(vert.Ptr)

		// DF tree == connected component
		ccs = append(ccs, *cc)
	}

	return ccs, nil
}

/*
CCUnionFind implements an algorithm for finding the connected components of an undirected graph
by manipulating a disjoint-set data structure while traversing the input graph. One disjoint set
is initially created for each vertex, and then for every edge that connects vertices in different
disjoint sets, the sets are merged.

The "Union-Find" approach is better suited for dynamic graphs, where the sets of vertices and edges
change over time, with the disjoint-set data structure being manipulated to update already calculated
connected components - O(α(V)) amortized time per operation -, which makes it possible for the algorithm
to be executed only once per graph.

For static graphs, however, CCUnionFind has an asymptotically worse time complexity than CCDFS:
CCUnionFind is O((V + E) α(V)), amortized (superlinear), while CCDFS is Θ(V + E) (linear).

Expectations:
	- The graph is correctly built.
	- The graph is undirected.

Complexity:
	- Time:  O((V + E) α(V)), amortized
	- Space: Θ(V)
*/
func CCUnionFind[T ds.Item](g *ds.G[T]) ([]CC[T], error) {
	if g.Directed() {
		return nil, ds.ErrUndefOp
	}

	sets := map[*T]CC[T]{}
	d := ds.NewDSet[T]()

	for v := range g.E {
		d.MakeSet(v)
	}

	for _, vert := range g.V {
		for _, e := range g.E[vert.Ptr] {
			if d.FindSet(e.Src) != d.FindSet(e.Dst) {
				d.Union(e.Src, e.Dst)
			}
		}
	}

	for _, vert := range g.V {
		set := d.FindSet(vert.Ptr)
		sets[set] = append(sets[set], vert.Ptr)
	}

	ccs := make([]CC[T], 0, len(sets))

	// Instead of iterating over the map directly,
	// we are using the existing vertex order, so
	// that the CC list is the same for every run.
	// If such consistency is not necessary, we
	// could just iterate over the map instead.
	// Asymptotically, it's all O(V) anyway.
	for _, vert := range g.V {
		cc, ok := sets[vert.Ptr]

		if !ok {
			continue
		}

		ccs = append(ccs, cc)
	}

	return ccs, nil
}
