package algo

import (
	"github.com/vc-souza/gga/ds"
)

/*
SCCAlgo describes the signature of an algorithm that can discover all
strongly connected components in a directed graph. If such an algorithm is
called on an undirected graph, the ds.ErrUndefOp error is returned.
*/
type SCCAlgo func(*ds.G) ([]SCC, error)

// An SCC holds the vertices in a strongly connected component of a directed graph.
type SCC []int

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
func SCCKosaraju(g *ds.G) ([]SCC, error) {
	if g.Undirected() {
		return nil, ds.ErrUndirected
	}

	var visit func(int)
	var scc *SCC

	visited := make([]bool, g.VertexCount())
	sccs := []SCC{}

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

	visit = func(v int) {
		visited[v] = true

		for _, e := range tg.V[v].E {
			if !visited[e.Dst] {
				visit(e.Dst)
			}
		}

		*scc = append(*scc, v)
	}

	for _, v := range ord {
		if visited[v] {
			continue
		}

		scc = &SCC{}

		visit(v)

		sccs = append(sccs, *scc)
	}

	return sccs, nil
}

/*
tjSCCAttrs is an auxiliary type used only by SCCTarjan to keep
track of extra data needed by the algorithm, per vertex.
*/
type tjSCCAttrs struct {
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
func SCCTarjan(g *ds.G) ([]SCC, error) {
	if g.Undirected() {
		return nil, ds.ErrUndirected
	}

	var visit func(int)

	att := make([]tjSCCAttrs, g.VertexCount())
	stack := ds.NewStack[int]()
	sccs := []SCC{}

	// using 1 as the starting point so that the zero-value
	// of tjAttrs.index (0) can indicate an unvisited vertex
	i := 1

	visit = func(v int) {
		att[v].index = i
		att[v].lowIndex = i

		stack.Push(v)
		att[v].onStack = true

		i++

		for _, e := range g.V[v].E {
			if att[e.Dst].index == 0 {
				visit(e.Dst)

				att[v].lowIndex = min(
					att[v].lowIndex,
					att[e.Dst].lowIndex,
				)
			} else if att[e.Dst].onStack {
				// can't use the low index of e.Dst since it is on the stack,
				// and as such, not in vtx's subtree: using the index
				// is the best we can do since we know vtx can reach e.Dst
				att[v].lowIndex = min(
					att[v].lowIndex,
					att[e.Dst].index,
				)
			}
		}

		// root of an SCC, otherwise do not pop anything
		if att[v].lowIndex == att[v].index {
			scc := SCC{}

			// every vertex that is currently on the stack
			// is a part of the SCC where vtx is the root,
			// so we pop until we find vtx
			for !stack.Empty() {
				w, _ := stack.Pop()
				att[w].onStack = false

				scc = append(scc, w)

				if w == v {
					break
				}
			}

			sccs = append(sccs, scc)
		}
	}

	for v := range g.V {
		if att[v].index == 0 {
			visit(v)
		}
	}

	return sccs, nil
}
