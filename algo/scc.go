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
algorithm will have found the SCCs of the original graph.

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

// tjAttrs is an auxiliary type used only by SCCTarjan.
type tjAttrs struct {
	// index represents when the vertex was first discovered.
	index int

	/*
		lowIndex represents the smallest index of any vertex currently on the stack
		known to be reachable from v through v's DFS subtree, including v itself.
	*/
	lowIndex int

	// next indicates the next adjacent vertex to explore.
	next int

	/*
		waiting flags that the vertex is waiting for one of its adjacent vertices
		to finish being explored, so that it can check their low index.
	*/
	waiting bool

	// onStack flags that a vertex is on the stack.
	onStack bool
}

/*
SCCTarjan implements Tarjan's algorithm for finding the strongly connected
components of a directed graph. A strongly connected component of a graph being
a subgraph where every vertex is reachable from every other vertex. Such
a subgraph is maximal: no other vertex or edge from the graph can be added
to the subgraph without breaking its property of being strongly connected.

Given a directed graph, SCCTarjan will keep an auxiliary stack where it pushes
vertices as soon as they are first visited during a modified DFS. The vertices
are not necessarily popped from the stack after being fully explored, though,
with the following invariant always holding:

  A vertex remains in the stack after being explored IFF there exists a path from
  the vertex to some other vertex earlier in the stack: meaning that a vertex is
  only removed from the stack after alls of its connected paths have been traversed.

If after exploring a vertex and all of its descendants, the vertex still has no
path to earlier vertices in the stack, then every vertex in the stack is popped
until the current vertex is reached (it is included): this set of vertices
is an SCC rooted at the vertex.

Link: https://en.wikipedia.org/wiki/Tarjan%27s_strongly_connected_components_algorithm

Expectations:
	- The graph is correctly built.
	- The graph is directed.

Complexity:
	- Time:  Θ(V + E)
	- Space: Θ(V)
*/
func SCCTarjan[V ds.Item](g *ds.Graph[V]) ([]SCC[V], error) {
	if g.Undirected() {
		return nil, ds.ErrUndefOp
	}

	// stack that simulates the call stack
	// necessary for the iterative version
	calls := ds.NewStack[*V]()

	// stack used by Tarjan's algorithm
	stack := ds.NewStack[*V]()

	sccs := []SCC[V]{}
	att := map[*V]*tjAttrs{}

	for v := range g.Adj {
		att[v] = &tjAttrs{}
	}

	// using 1 as the starting point so that the zero-value
	// of tjAttrs.index (0) can indicate an unvisited vertex
	i := 1

	visit := func(root *V) {
		calls.Push(root)

		for !calls.Empty() {
			vtx, _ := calls.Peek()

			// vertex is being discovered
			// assign a new index to it
			if att[vtx].index == 0 {
				att[vtx].index = i
				att[vtx].lowIndex = i

				stack.Push(vtx)
				att[vtx].onStack = true

				i++
			}

			// looking at the low value that the previous child
			// finished computed if vtx is waiting for a result
			if att[vtx].waiting {
				// adj list for the current vertex
				adj := g.Adj[vtx]

				// index of the pending child
				idx := att[vtx].next - 1

				// pending child
				child := adj[idx].Dst

				att[vtx].lowIndex = Min(att[vtx].lowIndex, att[child].lowIndex)
				att[vtx].waiting = false
			}

			// finished exploring adj
			if att[vtx].next >= len(g.Adj[vtx]) {
				calls.Pop()

				// root of an SCC, otherwise do not pop anything
				if att[vtx].lowIndex == att[vtx].index {
					scc := SCC[V]{}

					// every vertex that is currently in the stack
					// is a part of the SCC where vtx is the root,
					// so we pop until we find vtx
					for !stack.Empty() {
						w, _ := stack.Pop()
						att[w].onStack = false

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
			for i := att[vtx].next; i < len(g.Adj[vtx]); i++ {
				e := g.Adj[vtx][i]
				att[vtx].next++

				// will need to wait for the adjancent
				// vertex to have its low value calculated,
				// then it can be used to update vtx's
				if att[e.Dst].index == 0 {
					calls.Push(e.Dst)
					att[vtx].waiting = true
					break
				} else if att[e.Dst].onStack {
					// can't use the lowIndex of e.Dst since it is in the stack,
					// and as such, not in vtx's subtree: using the index
					// is the best we can do since we know vtx can reach e.Dst
					att[vtx].lowIndex = Min(att[vtx].lowIndex, att[e.Dst].index)
				}
			}
		}
	}

	for _, vert := range g.Verts {
		root := vert.Val

		if att[root].index != 0 {
			continue
		}

		visit(root)
	}

	return sccs, nil
}
