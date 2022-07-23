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
func TSort(g *ds.G) ([]int, error) {
	if g.Undirected() {
		return nil, ds.ErrUndirected
	}

	var visit func(int)

	count := g.VertexCount()
	ordIdx := count - 1

	ord := make([]int, count)
	visited := map[int]bool{}

	for i := range g.V {
		visited[i] = false
	}

	visit = func(v int) {
		visited[v] = true

		for _, e := range g.V[v].E {
			if !visited[e.Dst] {
				visit(e.Dst)
			}
		}

		ord[ordIdx] = v
		ordIdx--
	}

	for v := range g.V {
		if visited[v] {
			continue
		}

		visit(v)
	}

	return ord, nil
}
