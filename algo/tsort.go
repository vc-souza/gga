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

Link: https://en.wikipedia.org/wiki/Topological_sorting

Expectations:
	- The graph is correctly built.
	- The graph is directed.

Complexity:
	- Time:  Θ(V + E)
	- Space: O(V)
*/
func TSort[V ds.Item](g *ds.G[V]) ([]*V, error) {
	if g.Undirected() {
		return nil, ds.ErrUndefOp
	}

	calls := ds.NewStack[*V]()

	count := g.VertexCount()
	idx := count - 1

	ord := make([]*V, count)

	visited := map[*V]bool{}
	next := map[*V]int{}

	for _, vert := range g.V {
		if visited[vert.Ptr] {
			continue
		}

		calls.Push(vert.Ptr)

		for !calls.Empty() {
			vtx, _ := calls.Peek()
			visited[vtx] = true

			if next[vtx] >= len(g.E[vtx]) {
				calls.Pop()

				ord[idx] = vtx
				idx--

				continue
			}

			for i := next[vtx]; i < len(g.E[vtx]); i++ {
				e := g.E[vtx][i]
				next[vtx]++

				if !visited[e.Dst] {
					calls.Push(e.Dst)

					break
				}
			}
		}
	}

	return ord, nil
}
