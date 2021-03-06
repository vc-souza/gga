package algo

import (
	"github.com/vc-souza/gga/ds"
)

/*
GSCC implements an algorithm for building a condensation graph GSCC from a directed graph.

A condensation graph is the result of contracting each strongly connected component
(SCC) of the original graph to a vertex in the GSCC. Each edge (u, v) in the GSCC
exists because there exists at least one edge (x, y) in the original graph such that
x is in the SCC u and y is in the SCC v. One important property of the condensation
graph GSCC is that it is guaranteed to be a DAG (Directed Acyclic Graph).

Tarjan's SCC algorithm is used to calculate the SCCs because of its side-effect
of returning the SCCs in reverse topological order of the GSCC, which means
that if the SCC list is traversed in reverse we will never examine an edge
connecting the current SCC to an SCC that was examined earlier. In theory,
the space required to keep track of possible adjacencies between SCCs
is reduced by one after each SCC that is examined, which would not change
the worst-case space complexity of this step, O(V), but might end up yielding
better constant factors in the end.

Expectations:
	- The graph is correctly built.
	- The graph is directed.

Complexity:
	- Time:  Θ(V + E)
	- Space: Θ(V)
*/
func GSCC(g *ds.G) (*ds.G, []SCC, error) {
	if g.Undirected() {
		return nil, nil, ds.ErrUndirected
	}

	// Θ(V + E)
	sccs, err := SCCTarjan(g)

	if err != nil {
		return nil, nil, err
	}

	// Using Θ(V) extra space to achieve O(1) amortized
	// query time later, when the SCC that a vertex
	// belongs to will need to be queried Θ(E) times,
	// when building the adjacency list of the GSCC.
	vtxSCC := make([]int, g.VertexCount())

	for id, scc := range sccs {
		for _, v := range scc {
			vtxSCC[v] = id
		}
	}

	gscc := ds.NewDigraph()

	// By aligning the SCC id with the id of their
	// vertex in the GSCC we can get the Group
	// in O(1) amortized time later, without having
	// to use any extra space to map SCC ids to their
	// respective Group.
	for id := range sccs {
		items := make([]ds.Item, len(sccs[id]))

		for i, v := range sccs[id] {
			items[i] = g.V[v].Item
		}

		gscc.AddVertex(&ds.Group{
			Items: items,
			Id:    id,
		})
	}

	// Using O(V) space to keep trap of the adjacency
	// relationships between the SCCs in the GSCC.
	gsccAdj := make([]int, len(sccs)-1)

	// Traversing the SCC list in reverse to capitalize
	// on the fact that Tarjan's SCC algorithm returns
	// the list in reverse topological order of the GSCC.
	// Using this property, the SCC with index 0 can be
	// skipped, since it is the last one in that order.
	for srcId := len(sccs) - 1; srcId > 0; srcId-- {
		for _, v := range sccs[srcId] {
			for _, e := range g.V[v].E {
				dstId := vtxSCC[e.Dst]

				// vertices in the same SCC, skip.
				if srcId == dstId {
					continue
				}

				// Since the same slice is used by all components
				// to keep track of their adjacencies, and we
				// do not want to zero the whole thing after
				// each SCC is processed, we are using the id
				// of the current SCC to keep track of edges.
				// Which means that, while looking at a position
				// in the slice, if any other number that is not
				// the id of the current SCC is found, we treat
				// it as a 0, and the respective edge as not
				// currently existing.

				// edge between SCCs already exists, skip.
				if gsccAdj[dstId] == srcId {
					continue
				}

				// Since the SCC list and the GSCC vertex list are aligned,
				// we can find the Group assigned to SCC x by looking
				// at the vertex of GSCC at index x.
				gscc.AddEdge(
					gscc.V[srcId].Item,
					gscc.V[dstId].Item,
					0,
				)

				// marking the edge for the current SCC
				gsccAdj[dstId] = srcId
			}
		}

		// Attempting to capitalize on the property that there will
		// never be any edges in the next iteration that arrive at
		// vertices contained in SCCs already examined, by shrinking
		// the slice being used to keep track of adjacencies.
		gsccAdj = gsccAdj[: srcId-1 : srcId]
	}

	return gscc, sccs, nil
}
