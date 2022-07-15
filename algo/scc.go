package algo

import (
	"github.com/vc-souza/gga/ds"
)

// TODO: docs
type SCC[V ds.Item] []*V

// TODO: docs
func SCCKosaraju[V ds.Item](g *ds.Graph[V]) ([]SCC[V], error) {
	if g.Undirected() {
		return nil, ds.ErrUndefOp
	}

	stk := ds.NewStack[*V]()
	sccs := []SCC[V]{}

	visited := map[*V]bool{}
	next := map[*V]int{}

	// Θ(V + E)
	ord, err := TSort(g)

	if err != nil {
		return nil, err
	}

	// Θ(V + E)
	tg, err := g.Transpose()

	if err != nil {
		return nil, err
	}

	for e := ord.Front(); e != nil; e = e.Next() {
		v, ok := e.Value.(*V)

		if !ok {
			return nil, ds.ErrInvalidType
		}

		if visited[v] {
			continue
		}

		scc := SCC[V]{}

		stk.Push(v)

		for !stk.Empty() {
			vtx, _ := stk.Peek()
			visited[vtx] = true

			if next[vtx] >= len(tg.Adj[vtx]) {
				stk.Pop()
				scc = append(scc, vtx)
				continue
			}

			for i := next[vtx]; i < len(tg.Adj[vtx]); i++ {
				e := tg.Adj[vtx][i]
				next[vtx]++

				if !visited[e.Dst] {
					stk.Push(e.Dst)
					break
				}
			}
		}

		sccs = append(sccs, scc)
	}

	return sccs, nil
}
