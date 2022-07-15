package algo

import (
	"container/list"

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

Complexity:
	- Time:  Θ(V + E)
	- Space: O(V)
*/
func TSort[V ds.Item](g *ds.Graph[V]) (*list.List, error) {
	if g.Undirected() {
		return nil, ds.ErrUndefOp
	}

	stk := ds.NewStack[*V]()
	ord := list.New()

	visited := map[*V]bool{}
	next := map[*V]int{}

	for _, vert := range g.Verts {
		if visited[vert.Val] {
			continue
		}

		stk.Push(vert.Val)

		for !stk.Empty() {
			vtx, _ := stk.Peek()
			visited[vtx] = true

			if next[vtx] >= len(g.Adj[vtx]) {
				stk.Pop()
				ord.PushFront(vtx)
				continue
			}

			for i := next[vtx]; i < len(g.Adj[vtx]); i++ {
				e := g.Adj[vtx][i]
				next[vtx]++

				if !visited[e.Dst] {
					stk.Push(e.Dst)
					break
				}
			}
		}
	}

	return ord, nil
}
