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

	hasGrayTop := func(stk ds.Stack[*V]) bool {
		top, ok := stk.Peek()

		if !ok {
			return false
		}

		return fst[top].Color == ColorGray
	}

	// classify the edge that connects a gray vertex being explored
	// to another gray vertex that has also been discovered before
	// TODO: O(???)
	classify := func(e *ds.GraphEdge[V]) {
		// the vertex being reached (Dst) was discovered before
		// the vertex being explored (Src), so Dst is either
		// an ancestor of Src, or they do not have a direct
		// ancestor/descendant relationship.
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
	}

	// build a DFS tree rooted at the given vertex;
	// the tree will be a part of the DFS forest
	// TODO: O(???)
	tree := func(root *V) {
		// only using the ds.Stack interface
		stk := ds.Stack[*V](new(ds.Deque[*V]))

		stk.Push(root)

		for !stk.Empty() {
			vtx, _ := stk.Peek()

			t++

			fst[vtx].Discovery = t
			fst[vtx].Color = ColorGray

			// traversing the adjacency list backwards
			// because we are pushing onto the stack
			for i := len(g.Adj[vtx]) - 1; i >= 0; i-- {
				e := g.Adj[vtx][i]

				if fst[e.Dst].Color != ColorWhite {
					classify(e)
					continue
				}

				// found a tree edge
				fst[e.Dst].Parent = vtx

				stk.Push(e.Dst)
			}

			// we need to pop elements from the stack
			// until it is either empty or has a white
			// vertex on top, for the next iteration
			for hasGrayTop(stk) {
				vtx, _ := stk.Pop()

				t++

				fst[vtx].Finish = t
			}
		}
	}

	// TODO: O(???)
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
