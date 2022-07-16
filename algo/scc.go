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

	calls := ds.NewStack[*V]()
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

		calls.Push(v)

		for !calls.Empty() {
			vtx, _ := calls.Peek()
			visited[vtx] = true

			if next[vtx] >= len(tg.Adj[vtx]) {
				calls.Pop()
				scc = append(scc, vtx)
				continue
			}

			for i := next[vtx]; i < len(tg.Adj[vtx]); i++ {
				e := tg.Adj[vtx][i]
				next[vtx]++

				if !visited[e.Dst] {
					calls.Push(e.Dst)
					break
				}
			}
		}

		sccs = append(sccs, scc)
	}

	return sccs, nil
}

// TODO: docs
func SCCTarjan[V ds.Item](g *ds.Graph[V]) ([]SCC[V], error) {
	if g.Undirected() {
		return nil, ds.ErrUndefOp
	}

	calls := ds.NewStack[*V]()
	stack := ds.NewStack[*V]()
	sccs := []SCC[V]{}

	index := map[*V]int{}
	low := map[*V]int{}
	next := map[*V]int{}
	wait := map[*V]int{}
	on := map[*V]bool{}

	i := 1

	visit := func(root *V) {
		calls.Push(root)

		for !calls.Empty() {
			vtx, _ := calls.Peek()

			// vertex is being discovered
			if index[vtx] == 0 {
				index[vtx] = i
				low[vtx] = i

				stack.Push(vtx)
				on[vtx] = true

				i++
			}

			// looking at the low value that the previous child computed
			// if next is not 0, then an unvisited child started
			// calculating its low value, and the current vertex needs it
			if next[vtx] != 0 {
				// adj list for the current vertex
				adj := g.Adj[vtx]

				// index of the pending child
				idx := next[vtx] - 1

				// pending child
				child := adj[idx].Dst

				low[vtx] = Min(low[vtx], low[child])
			}

			// finished exploring adj
			if next[vtx] >= len(g.Adj[vtx]) {
				calls.Pop()

				// root of a SCC
				if low[vtx] == index[vtx] {
					scc := SCC[V]{}

					for !stack.Empty() {
						w, _ := stack.Pop()
						on[w] = false

						scc = append(scc, w)

						if w == vtx {
							break
						}
					}

					sccs = append(sccs, scc)
				}

				continue
			}

			// visit adjacent vertices
			for i := next[vtx]; i < len(g.Adj[vtx]); i++ {
				e := g.Adj[vtx][i]
				next[vtx]++

				// will need to wait for the adjancent
				// node to have its low value calculated,
				// then it can be used to update vtx's
				if index[e.Dst] == 0 {
					calls.Push(e.Dst)
					break
				} else if on[e.Dst] {
					low[vtx] = Min(low[vtx], index[e.Dst])
				}
			}
		}
	}

	for _, vert := range g.Verts {
		root := vert.Val

		if index[root] != 0 {
			continue
		}

		visit(root)
	}

	return sccs, nil
}
