package algo

import (
	"github.com/vc-souza/gga/ds"
)

// TODO: docs
func Condensation[V ds.Item](g *ds.Graph[V]) (*ds.Graph[ds.Items[V]], []SCC[V], error) {
	if g.Undirected() {
		return nil, nil, ds.ErrUndefOp
	}

	// Θ(V + E)
	sccs, err := SCCTarjan(g)

	if err != nil {
		return nil, nil, err
	}

	vtxSCC := make(map[*V]int)

	for id, scc := range sccs {
		for _, v := range scc {
			vtxSCC[v] = id
		}
	}

	gscc := ds.NewDirectedGraph[ds.Items[V]]()

	// TODO: explain
	ptrs := map[int]*ds.Items[V]{}

	// TODO: explain
	for sccId := len(sccs) - 1; sccId >= 0; sccId-- {
		ls := ds.Items[V](sccs[sccId])
		ptrs[sccId] = &ls

		gscc.UnsafeAddVertex(&ls)
	}

	// TODO: explain
	gsccAdj := make([]int, len(sccs)-1)

	// TODO: explain
	for sccId := len(sccs) - 1; sccId > 0; sccId-- {
		scc := sccs[sccId]

		// TODO: explain
		for _, v := range scc {
			srcSCC := vtxSCC[v]

			for _, e := range g.Adj[v] {
				dstSCC := vtxSCC[e.Dst]

				// same SCC, bail
				if srcSCC == dstSCC {
					continue
				}

				// TODO: explain
				// edge already exists, bail
				if gsccAdj[dstSCC] == sccId {
					continue
				}

				gscc.UnsafeAddWeightedEdge(
					ptrs[srcSCC],
					ptrs[dstSCC],
					0,
				)

				// TODO: explain
				gsccAdj[dstSCC] = sccId
			}
		}

		// TODO: explain
		gsccAdj = gsccAdj[: sccId-1 : sccId]
	}

	return gscc, sccs, nil
}
