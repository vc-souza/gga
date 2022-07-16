package algo

import (
	"github.com/vc-souza/gga/ds"
)

/*
SCCAlgorithm describes the signature of an algorithm that can discover all
strongly connected components in a graph. If a particular algorithm can
only work on a particular type of graph, then undefined behavior is
indicated by the ds.ErrUndefOp error being returned.
*/
type SCCAlgorithm func(*ds.Graph[ds.Text]) ([]SCC[ds.Text], error)

// An SCC holds the vertices in a strongly connected component of a graph.
type SCC[V ds.Item] []*V

/*
SCCKosaraju implements Kosaraju's algorithm for finding the strongly connected
components of a directed graph. A strongly connected component of a graph being
a subgraph where every vertex is reachable from every other vertex. Such
a subgraph is maximal: no other vertex or edge from the graph can be added
to the subgraph without breaking its property of being strongly connected.

Given a directed graph, SCCKosaraju will obtain an ordering of the vertices,
in decreasing order of finish time in a DFS. This is implemented as a call
to TSort (Topological Sort), even if the final sorting might not be an
actual topological sorting (undefined for cyclic graphs), it will still
return an ordering of vertices in decreasing order of finish time in a DFS.

A transpose of the original graph is then calculated (same graph with the direction
of every edge reversed), and a second DFS is executed on it (TSort being a DFS itself).
The second DFS uses the ordering obtained from the Topological Sort to calculate
the DF forest of the transpose (the main loop of the DFS will visit the vertices in
that order), and each DF tree in the forest will correspond to an SCC of the transpose.
Since a graph and its transpose share the same SCCs, after the second DFS, the
algorith will have found the SCCs of the original graph.

Link: https://en.wikipedia.org/wiki/Kosaraju%27s_algorithm

Expectations:
	- The graph is correctly built.
	- The graph is directed.

Complexity:
	- Time:  Θ(V + E)
	- Space: Θ(V)
*/
func SCCKosaraju[V ds.Item](g *ds.Graph[V]) ([]SCC[V], error) {
	if g.Undirected() {
		return nil, ds.ErrUndefOp
	}

	stk := ds.NewStack[*V]()
	sccs := []SCC[V]{}

	visited := map[*V]bool{}
	next := map[*V]int{}

	// Θ(V + E)
	ord, err := TSort(g)

	if err != nil {
		return nil, err
	}

	// Θ(V + E)
	tg, err := g.Transpose()

	if err != nil {
		return nil, err
	}

	for e := ord.Front(); e != nil; e = e.Next() {
		v, ok := e.Value.(*V)

		if !ok {
			return nil, ds.ErrInvalidType
		}

		if visited[v] {
			continue
		}

		scc := SCC[V]{}

		stk.Push(v)

		for !stk.Empty() {
			vtx, _ := stk.Peek()
			visited[vtx] = true

			if next[vtx] >= len(tg.Adj[vtx]) {
				stk.Pop()
				scc = append(scc, vtx)
				continue
			}

			for i := next[vtx]; i < len(tg.Adj[vtx]); i++ {
				e := tg.Adj[vtx][i]
				next[vtx]++

				if !visited[e.Dst] {
					stk.Push(e.Dst)
					break
				}
			}
		}

		sccs = append(sccs, scc)
	}

	return sccs, nil
}
