package algo

import (
	"errors"
	"math"

	"github.com/vc-souza/gga/ds"
)

// TODO: docs
type BFSNode[V ds.Item] struct {
	// TODO: docs
	Distance int

	// TODO: docs
	Color int

	// TODO: docs
	Parent *V
}

// TODO: docs
type BFSTree[V ds.Item] map[*V]*BFSNode[V]

// TODO: docs Θ(V +  E)
// TODO: warning: assumes that src exists, and that the graph is correctly built: otherwise, the behavior is undefined
func BFS[V ds.Item](g *ds.Graph[V], src *V) (BFSTree[V], error) {
	tree := BFSTree[V]{}

	// Θ(V)
	for v := range g.Adj {
		tree[v] = &BFSNode[V]{
			Distance: math.MaxInt32,
			Color:    ColorWhite,
			Parent:   nil,
		}
	}

	tree[src].Color = ColorGray
	tree[src].Distance = 0

	// only using the ds.Queue interface
	queue := ds.Queue[*V](new(ds.LLQueue[*V]))

	queue.Enqueue(src)

	// Θ(E)
	for !queue.Empty() {
		curr, ok := queue.Dequeue()

		if !ok {
			return nil, errors.New("could not dequeue")
		}

		for _, edge := range g.Adj[curr] {
			if tree[edge.Dst].Color != ColorWhite {
				continue
			}

			// found a tree edge
			tree[edge.Dst].Distance = tree[curr].Distance + 1
			tree[edge.Dst].Color = ColorGray
			tree[edge.Dst].Parent = curr

			queue.Enqueue(edge.Dst)
		}

		tree[curr].Color = ColorBlack
	}

	return tree, nil
}
