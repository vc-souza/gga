package algo

import (
	"sort"

	"github.com/vc-souza/gga/ds"
)

// TODO: docs
type MSTAlgo[T ds.Item] func(*ds.G[T]) (MST[T], error)

// TODO: docs
type MST[T ds.Item] []*ds.GE[T]

// TODO: docs
func MSTKruskal[T ds.Item](g *ds.G[T]) (MST[T], error) {
	if g.Directed() {
		return nil, ds.ErrUndefOp
	}

	edges := make([]*ds.GE[T], g.EdgeCount())

	for i, eIdx := 0, 0; i < len(g.V); i++ {
		es := g.E[g.V[i].Ptr]

		copy(edges[eIdx:], es)

		eIdx += len(es)
	}

	// O(E log E)
	sort.Stable(ds.ByEdgeWeight[T](edges))

	d := ds.NewDSet[T]()

	for _, vtx := range g.V {
		d.MakeSet(vtx.Ptr)
	}

	max := g.VertexCount() - 1
	mst := MST[T]{}

	for _, e := range edges {
		if d.FindSet(e.Src) == d.FindSet(e.Dst) {
			continue
		}

		d.Union(e.Src, e.Dst)

		mst = append(mst, e)

		if len(mst) == max {
			break
		}
	}

	return mst, nil
}
