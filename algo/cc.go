package algo

import "github.com/vc-souza/gga/ds"

/*
CCAlgo describes the signature of an algorithm that can discover all
connected components in an undirected graph. If such an algorithm
is called on a directed graph, the ds.ErrUndefOp error is returned.
*/
type CCAlgo[V ds.Item] func(*ds.G[V]) ([]CC[V], error)

// A CC holds the vertices in a connected component of an undirected graph.
type CC[V ds.Item] []*V

/*
TODO: docs

Expectations:
	- The graph is correctly built.
	- The graph is undirected.

Complexity:
	- Time:  Θ(V + E)
	- Space: Θ(V)
*/
func CCDFS[V ds.Item](g *ds.G[V]) ([]CC[V], error) {
	if g.Directed() {
		return nil, ds.ErrUndefOp
	}

	ccs := []CC[V]{}
	calls := ds.NewStack[*V]()
	attr := map[*V]*iDFS{}

	for v := range g.E {
		attr[v] = &iDFS{}
	}

	for _, vert := range g.V {
		if attr[vert.Ptr].visited {
			continue
		}

		cc := CC[V]{}

		calls.Push(vert.Ptr)

		for !calls.Empty() {
			vtx, _ := calls.Peek()
			attr[vtx].visited = true

			if attr[vtx].next >= len(g.E[vtx]) {
				calls.Pop()

				cc = append(cc, vtx)

				continue
			}

			for i := attr[vtx].next; i < len(g.E[vtx]); i++ {
				e := g.E[vtx][i]
				attr[vtx].next++

				if !attr[e.Dst].visited {
					calls.Push(e.Dst)
					break
				}
			}
		}

		ccs = append(ccs, cc)
	}

	return ccs, nil
}

/*
TODO: docs

Expectations:
	- The graph is correctly built.
	- The graph is undirected.

Complexity:
	- Time:  O((V + E) α(V)), amortized
	- Space: Θ(V)
*/
func CCUnionFind[V ds.Item](g *ds.G[V]) ([]CC[V], error) {
	if g.Directed() {
		return nil, ds.ErrUndefOp
	}

	sets := map[*V]CC[V]{}
	d := ds.NewDSet[V]()

	for v := range g.E {
		d.MakeSet(v)
	}

	for _, es := range g.E {
		for _, e := range es {
			if d.FindSet(e.Src) != d.FindSet(e.Dst) {
				d.Union(e.Src, e.Dst)
			}
		}
	}

	for v := range g.E {
		sets[d.FindSet(v)] = append(sets[d.FindSet(v)], v)
	}

	ccs := []CC[V]{}

	for _, cc := range sets {
		ccs = append(ccs, cc)
	}

	return ccs, nil
}
