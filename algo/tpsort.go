package algo

import (
	"container/list"

	"github.com/vc-souza/gga/ds"
)

// TODO: docs
func TopologicalSort[V ds.Item](g *ds.Graph[V]) (*list.List, error) {
	stk := ds.Stack[*V](new(ds.Deque[*V]))
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
