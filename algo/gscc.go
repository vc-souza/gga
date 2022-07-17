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
	vtxToId := make(map[*V]int)

	for id, scc := range sccs {
		for _, v := range scc {
			vtxToId[v] = id
		}
	}

	// TODO: explain
	gscc := ds.NewDirectedGraph[ds.ItemGroup[V]]()

	// TODO: explain, O(V) space
	idToPtr := make([]*ds.ItemGroup[V], len(sccs))

	// TODO: explain, O(V) space
	gsccAdj := make([]int, len(sccs)-1)

	// TODO: explain
	newGSCCVtx := func(id int) {
		if idToPtr[id] != nil {
			return
		}

		idToPtr[id] = &ds.ItemGroup[V]{
			Id:    strconv.Itoa(id),
			Items: sccs[id],
		}

		gscc.UnsafeAddVertex(idToPtr[id])
	}

	// TODO: explain
	newGSCCEdge := func(srcId, dstId int) {
		newGSCCVtx(dstId)

		gscc.UnsafeAddWeightedEdge(
			idToPtr[srcId],
			idToPtr[dstId],
			0,
		)

		// TODO: explain
		gsccAdj[dstId] = srcId
	}

	// TODO: explain
	for srcId := len(sccs) - 1; srcId >= 0; srcId-- {
		newGSCCVtx(srcId)

		// TODO: explain
		if srcId == 0 {
			break
		}

		// TODO: explain
		for _, v := range sccs[srcId] {
			for _, e := range g.Adj[v] {
				dstId := vtxToId[e.Dst]

				// same SCC, bail
				if srcId == dstId {
					continue
				}

				// TODO: explain
				// edge already exists, bail
				if gsccAdj[dstId] == srcId {
					continue
				}

				newGSCCEdge(srcId, dstId)
			}
		}

		// TODO: explain
		gsccAdj = gsccAdj[: srcId-1 : srcId]
	}

	return gscc, sccs, nil
}
