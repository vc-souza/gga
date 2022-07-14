package algo

import (
	"github.com/vc-souza/gga/ds"
)

// TODO: docs
func TopologicalSort[V ds.Item](g *ds.Graph[V]) ([]*V, error) {
	stk := ds.Stack[*V](new(ds.Deque[*V]))
	ord := make([]*V, 0, g.VertexCount())

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

				// TODO: linked list?
				ord = append([]*V{vtx}, ord...)

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
