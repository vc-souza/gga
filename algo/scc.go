package algo

import (
	"github.com/vc-souza/gga/ds"
)

/*
SCCAlgo describes the signature of an algorithm that can discover all
strongly connected components in a directed graph. If such an algorithm is
called on an undirected graph, the ds.ErrUndefOp error is returned.
*/
type SCCAlgo[T ds.Item] func(*ds.G[T]) ([]SCC[T], error)

// An SCC holds the vertices in a strongly connected component of a directed graph.
type SCC[T ds.Item] []*T

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

Expectations:
	- The graph is correctly built.
	- The graph is directed.

Complexity:
	- Time:  Θ(V + E)
	- Space: Θ(V)
*/
func SCCKosaraju[T ds.Item](g *ds.G[T]) ([]SCC[T], error) {
	if g.Undirected() {
		return nil, ds.ErrUndefOp
	}

	var visit func(*T)
	var scc *SCC[T]

	sccs := []SCC[T]{}
	visited := map[*T]bool{}

	for v := range g.E {
		visited[v] = false
	}

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

	visit = func(vtx *T) {
		visited[vtx] = true

		for _, e := range tg.E[vtx] {
			if !visited[e.Dst] {
				visit(e.Dst)
			}
		}

		*scc = append(*scc, vtx)
	}

	for _, v := range ord {
		if visited[v] {
			continue
		}

		scc = &SCC[T]{}

		visit(v)

		sccs = append(sccs, *scc)
	}

	return sccs, nil
}

// tjSCC is an auxiliary type used only by SCCTarjan.
type tjSCC struct {
	// index represents when the vertex was first discovered.
	index int

	/*
		lowIndex represents the smallest index of any vertex currently on the stack
		known to be reachable from v through v's DFS subtree, including v itself.
	*/
	lowIndex int

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

  A vertex remains on the stack after being explored IFF there exists a path from
  the vertex to some other vertex earlier on the stack: meaning that a vertex is
  only removed from the stack after alls of its connected paths have been traversed.

If after exploring a vertex and all of its descendants, the vertex still has no
path to earlier vertices on the stack, then every vertex on the stack is popped
until the current vertex is reached (it is included): this set of vertices
is an SCC rooted at the vertex.

An important property of Tarjan's algorithm is that the SCCs are discovered
in reverse topological order of the condensation graph of the input, which
is a DAG obtained by contracting every vertex in a SCC into a single vertex.

Expectations:
	- The graph is correctly built.
	- The graph is directed.

Complexity:
	- Time:  Θ(V + E)
	- Space: Θ(V)
*/
func SCCTarjan[T ds.Item](g *ds.G[T]) ([]SCC[T], error) {
	if g.Undirected() {
		return nil, ds.ErrUndefOp
	}

	var visit func(*T)

	stack := ds.NewStack[*T]()
	att := map[*T]*tjSCC{}
	sccs := []SCC[T]{}

	for v := range g.E {
		att[v] = &tjSCC{}
	}

	// using 1 as the starting point so that the zero-value
	// of tjAttrs.index (0) can indicate an unvisited vertex
	i := 1

	visit = func(vtx *T) {
		att[vtx].index = i
		att[vtx].lowIndex = i

		stack.Push(vtx)
		att[vtx].onStack = true

		i++

		for _, e := range g.E[vtx] {
			if att[e.Dst].index == 0 {
				visit(e.Dst)

				att[vtx].lowIndex = min(
					att[vtx].lowIndex,
					att[e.Dst].lowIndex,
				)
			} else if att[e.Dst].onStack {
				// can't use the low index of e.Dst since it is on the stack,
				// and as such, not in vtx's subtree: using the index
				// is the best we can do since we know vtx can reach e.Dst
				att[vtx].lowIndex = min(
					att[vtx].lowIndex,
					att[e.Dst].index,
				)
			}
		}

		// root of an SCC, otherwise do not pop anything
		if att[vtx].lowIndex == att[vtx].index {
			scc := SCC[T]{}

			// every vertex that is currently on the stack
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
	}

	for _, vert := range g.V {
		if att[vert.Ptr].index == 0 {
			visit(vert.Ptr)
		}
	}

	return sccs, nil
}
