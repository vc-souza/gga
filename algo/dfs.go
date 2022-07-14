package algo

import (
	"github.com/vc-souza/gga/ds"
)

// TODO: docs
type DFSNode[V ds.Item] struct {
	// TODO: docs
	Discovery int

	// TODO: docs
	Finish int

	// TODO: docs
	Color int

	// TODO: docs
	Parent *V

	next int
}

// TODO: docs
type DFSForest[V ds.Item] map[*V]*DFSNode[V]

// TODO: docs
func DFS[V ds.Item](g *ds.Graph[V]) (DFSForest[V], *EdgeTypes[V], error) {
	fst := DFSForest[V]{}
	tps := &EdgeTypes[V]{}
	t := 0

	// Θ(V)
	for v := range g.Adj {
		fst[v] = &DFSNode[V]{
			Color: ColorWhite,
		}
	}

	// classify the edge that connects a gray vertex being explored
	// to another gray vertex that has also been discovered before
	classify := func(e *ds.GraphEdge[V]) {
		if g.Directed() {
			// the vertex being reached (Dst) was discovered before
			// the vertex being explored (Src), so Dst is either
			// an ancestor of Src, or they do not have a direct
			// ancestor/descendant relationship
			if fst[e.Src].Discovery >= fst[e.Dst].Discovery {
				// ancestor/descendant relationship,
				// self-loops included here
				if fst[e.Dst].Finish == 0 {
					tps.Back = append(tps.Back, e)
				} else {
					tps.Cross = append(tps.Cross, e)
				}
			} else {
				// Src is an ancestor of Dst, and since Dst has
				// been discovered before, this is a Forward edge
				tps.Forward = append(tps.Forward, e)
			}
		} else {
			// due to how adjacency lists work, undirected
			// graphs represent the same edge twice, so
			// if we're dealing with the reverse of a tree
			// edge, then do not flag the reverse edge as
			// being a back edge
			if fst[e.Src].Parent == e.Dst {
				return
			}

			// undirected graphs only have tree and back edges
			// even if this looks like a forward edge from one
			// side, it will be classified as a back edge
			// when the reverse edge gets classified
			tps.Back = append(tps.Back, e)
		}
	}

	// build a DFS tree rooted at the given vertex;
	// the tree will be a part of the DFS forest
	tree := func(root *V) {
		// only using the ds.Stack interface
		stk := ds.Stack[*V](new(ds.Deque[*V]))

		stk.Push(root)

		for !stk.Empty() {
			vtx, _ := stk.Peek()

			if fst[vtx].Color == ColorWhite {
				t++

				fst[vtx].Discovery = t
				fst[vtx].Color = ColorGray
			}

			if fst[vtx].next >= len(g.Adj[vtx]) {
				vtx, _ := stk.Pop()

				t++

				fst[vtx].Finish = t

				continue
			}

			for i := fst[vtx].next; i < len(g.Adj[vtx]); i++ {
				e := g.Adj[vtx][i]

				if fst[e.Dst].Color != ColorWhite {
					classify(e)
					fst[vtx].next++
					continue
				}

				// found a tree edge
				fst[e.Dst].Parent = vtx

				stk.Push(e.Dst)

				fst[vtx].next++
				break
			}
		}
	}

	for _, vert := range g.Verts {
		root := vert.Val

		// already part of another tree
		if fst[root].Color != ColorWhite {
			continue
		}

		tree(root)
	}

	return fst, tps, nil
}
