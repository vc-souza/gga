package algo

import "github.com/vc-souza/gga/ds"

/*
CCAlgo describes the signature of an algorithm that can discover all
connected components in an undirected graph. If such an algorithm
is called on a directed graph, the ds.ErrUndefOp error is returned.
*/
type CCAlgo func(*ds.G) ([]CC, error)

// A CC holds the vertices in a connected component of an undirected graph.
type CC []int

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
func CCDFS(g *ds.G) ([]CC, error) {
	if g.Directed() {
		return nil, ds.ErrDirected
	}

	var visit func(int)
	var cc *CC

	visited := make([]bool, g.VertexCount())
	ccs := []CC{}

	visit = func(v int) {
		visited[v] = true

		for _, e := range g.V[v].E {
			if !visited[e.Dst] {
				visit(e.Dst)
			}
		}

		*cc = append(*cc, v)
	}

	for v := range g.V {
		if visited[v] {
			continue
		}

		cc = &CC{}

		visit(v)

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
func CCUnionFind(g *ds.G) ([]CC, error) {
	if g.Directed() {
		return nil, ds.ErrDirected
	}

	sets := map[int]CC{}
	d := ds.NewDSet[int]()

	for v := range g.V {
		d.MakeSet(v)
	}

	for v := range g.V {
		for _, e := range g.V[v].E {
			if d.FindSet(e.Src) != d.FindSet(e.Dst) {
				d.Union(e.Src, e.Dst)
			}
		}
	}

	for v := range g.V {
		set := d.FindSet(v)
		sets[set] = append(sets[set], v)
	}

	ccs := make([]CC, 0, len(sets))

	// Instead of iterating over the map directly,
	// we are using the existing vertex order, so
	// that the CC list is the same for every run.
	// If such consistency is not necessary, we
	// could just iterate over the map instead.
	// Asymptotically, it's all O(V) anyway.
	for v := range g.V {
		cc, ok := sets[v]

		if !ok {
			continue
		}

		ccs = append(ccs, cc)
	}

	return ccs, nil
}
