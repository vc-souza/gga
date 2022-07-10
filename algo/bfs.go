package algo

import (
	"errors"
	"math"

	"github.com/vc-souza/gga/ds"
)

// TODO: docs
type VertexAttrs[V ds.Item] struct {
	// TODO: docs
	Distance int

	// TODO: docs
	Color int

	// TODO: docs
	Parent *V
}

// TODO: docs
type BFSResult[V ds.Item] map[*V]*VertexAttrs[V]

// TODO: docs Θ(V +  E)
// TODO: warning: assumes src exists, assumes the graph is correctly built: otherwise, the behavior is undefined
func BFS[V ds.Item](g *ds.Graph[V], src *V) (BFSResult[V], error) {
	attrs := BFSResult[V]{}

	// Θ(V)
	for _, vert := range g.Verts {
		v := vert.Sat

		attrs[v] = &VertexAttrs[V]{
			Distance: math.MaxInt32,
			Color:    ColorWhite,
			Parent:   nil,
		}
	}

	attrs[src].Distance = 0
	attrs[src].Color = ColorGray

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
			if attrs[edge.Dst].Color != ColorWhite {
				continue
			}

			// found a tree edge
			attrs[edge.Dst].Distance = attrs[curr].Distance + 1
			attrs[edge.Dst].Color = ColorGray
			attrs[edge.Dst].Parent = curr

			queue.Enqueue(edge.Dst)
		}

		attrs[curr].Color = ColorBlack
	}

	return attrs, nil
}
