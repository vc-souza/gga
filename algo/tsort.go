package algo

import (
	"github.com/vc-souza/gga/ds"
)

/*
TSort implements an algorithm for Topological Sorting.

Given a directed graph, TSort eventually produces an ordering of the vertices in the original graph
such that, for every edge (u,v) vertex u appears before vertex v in the final ordering.

This algorithm is a simplified version of a DFS, with the addition of a linked list that stores the
ordering of the vertices, which is appended to - at the head - after each vertex is fully explored.

Expectations:
	- The graph is correctly built.
	- The graph is directed.

Complexity:
	- Time:  Θ(V + E)
	- Space: Θ(V)
*/
func TSort[T ds.Item](g *ds.G[T]) ([]*T, error) {
	if g.Undirected() {
		return nil, ds.ErrUndirected
	}

	var visit func(*T)

	count := g.VertexCount()
	ordIdx := count - 1

	ord := make([]*T, count)
	visited := map[*T]bool{}

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

		ord[ordIdx] = vtx
		ordIdx--
	}

	for _, vert := range g.V {
		if visited[vert.Ptr] {
			continue
		}

		visit(vert.Ptr)
	}

	return ord, nil
}
