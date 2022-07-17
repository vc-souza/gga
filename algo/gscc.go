package algo

import (
	"strconv"

	"github.com/vc-souza/gga/ds"
)

// TODO: docs
func GSCC[V ds.Item](g *ds.Graph[V]) (*ds.Graph[ds.ItemGroup[V]], []SCC[V], error) {
	if g.Undirected() {
		return nil, nil, ds.ErrUndefOp
	}

	// Θ(V + E)
	sccs, err := SCCTarjan(g)

	if err != nil {
		return nil, nil, err
	}

	// TODO: explain, Θ(V) space
	vtxSCC := make(map[*V]int)

	for id, scc := range sccs {
		for _, v := range scc {
			vtxSCC[v] = id
		}
	}

	// TODO: explain
	gscc := ds.NewDirectedGraph[ds.ItemGroup[V]]()

	// TODO: explain (cc id == vtx id int he GSCC)
	for id := range sccs {
		gscc.UnsafeAddVertex(
			&ds.ItemGroup[V]{
				Id:    strconv.Itoa(id),
				Items: sccs[id],
			},
		)
	}

	// TODO: explain, O(V) space
	gsccAdj := make([]int, len(sccs)-1)

	// TODO: explain (also, why not 0?)
	for srcId := len(sccs) - 1; srcId > 0; srcId-- {
		// TODO: explain
		for _, v := range sccs[srcId] {
			for _, e := range g.Adj[v] {
				dstId := vtxSCC[e.Dst]

				// same SCC, bail
				if srcId == dstId {
					continue
				}

				// TODO: explain
				// edge already exists, bail
				if gsccAdj[dstId] == srcId {
					continue
				}

				// TODO: explain: topo sort,
				// one edges now are going forwards

				// TODO: explain
				gscc.UnsafeAddWeightedEdge(
					gscc.Verts[srcId].Val,
					gscc.Verts[dstId].Val,
					0,
				)

				// TODO: explain
				gsccAdj[dstId] = srcId
			}
		}

		// TODO: explain
		gsccAdj = gsccAdj[:srcId-1]
	}

	return gscc, sccs, nil
}
