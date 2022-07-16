package algo

import (
	"github.com/vc-souza/gga/ds"
)

// TODO: docs
func Condensation[V ds.Item](g *ds.Graph[V], f SCCAlgorithm[V]) (*ds.Graph[ds.ItemList[V]], error) {
	if g.Undirected() {
		return nil, ds.ErrUndefOp
	}

	if f == nil {
		f = SCCTarjan[V]
	}

	// Θ(V + E)
	sccs, err := f(g)

	if err != nil {
		return nil, err
	}

	vtxSCC := make(map[*V]int)

	for id, scc := range sccs {
		for _, v := range scc {
			vtxSCC[v] = id
		}
	}

	vCount := g.VertexCount()
	adjMtx := make([][]int, vCount)

	for i := range adjMtx {
		adjMtx[i] = make([]int, vCount)
	}

	gscc := ds.NewDirectedGraph[ds.ItemList[V]]()

	sccAddr := map[int]*ds.ItemList[V]{}

	for i := len(sccs) - 1; i >= 0; i-- {
		ls := ds.ItemList[V](sccs[i])

		sccAddr[i] = &ls

		gscc.UnsafeAddVertex(&ls)
	}

	for v, es := range g.Adj {
		srcSCC := vtxSCC[v]

		for _, e := range es {
			dstSCC := vtxSCC[e.Dst]

			// same SCC, bail
			if srcSCC == dstSCC {
				continue
			}

			// edge already exists, bail
			if adjMtx[srcSCC][dstSCC] == 1 {
				continue
			}

			gscc.UnsafeAddWeightedEdge(
				sccAddr[srcSCC],
				sccAddr[dstSCC],
				0,
			)

			adjMtx[srcSCC][dstSCC] = 1
		}
	}

	return gscc, nil
}
