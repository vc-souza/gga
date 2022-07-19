package algo

import (
	"github.com/vc-souza/gga/ds"
)

/*
A DFNode node represents a node in a Depth-First tree in a Depth-First forest, holding
the attributes produced by a DFS, for a particular vertex. At the end of the DFS, every
vertex is part of one of the DF trees in the DF forest produced by the algorithm.
*/
type DFNode[V ds.Item] struct {
	iDFS

	// Discovery records when the vertex was marked as discovered.
	Discovery int

	// Finish records when the vertex was marked as fully explored.
	Finish int

	/*
		Parent holds the vertex that discovered this vertex, with the edge (v.Parent, v) being called a tree edge.
		This is how the DF tree that this vertex is a part of is encoded: by following the parent pointers from
		this vertex, one can get to the root of the DF tree.

		After a DFS, every root of a DF tree in the DF forest will have a nil Parent.
	*/
	Parent *V
}

/*
A DFForest is one of the results of a DFS, representing a forest of DF trees (connected acyclic subgraphs),
with each DF tree being composed of all (at the time) unvisited vertices that were reachable from the root
of the DF tree, during the execution of the DFS.

Slightly different trees can be generated for the same graph, if the visiting order for either vertices
or edges is changed.

The gga graph implementation guarantees both vertex and edge traversal in insertion order,
so repeated DFS calls always produce the same DF forest.
*/
type DFForest[V ds.Item] map[*V]*DFNode[V]

func classifyDirectedEdge[V ds.Item](fst DFForest[V], tps *EdgeTypes[V], e *ds.GE[V]) {
	// the vertex being reached (Dst) was discovered before
	// the vertex being explored (Src), so Dst is either
	// an ancestor of Src, or they do not have a direct
	// ancestor/descendant relationship
	if fst[e.Src].Discovery >= fst[e.Dst].Discovery {
		// ancestor/descendant relationship,
		// self-loops included here
		if fst[e.Dst].Finish == 0 {
			tps.Back = append(tps.Back, e)
		} else {
			tps.Cross = append(tps.Cross, e)
		}
	} else {
		// Src is an ancestor of Dst, and since Dst has
		// been discovered before, this is a Forward edge
		tps.Forward = append(tps.Forward, e)
	}
}

func classifyUndirectedEdge[V ds.Item](fst DFForest[V], tps *EdgeTypes[V], e *ds.GE[V]) {
	// due to how adjacency lists work, undirected
	// graphs represent the same edge twice, so
	// if we're dealing with the reverse of a tree
	// edge, then do not flag the reverse edge as
	// being a back edge
	if fst[e.Src].Parent == e.Dst {
		return
	}

	// undirected graphs only have tree and back edges
	// even if this looks like a forward edge from one
	// side, it will be classified as a back edge
	// when the reverse edge gets classified
	tps.Back = append(tps.Back, e)
}

/*
DFS implements the Depth-First Search (DFS) algorithm.

Given a graph, DFS eventually explores every vertex, building a DF forest that holds DF trees.
The search is conducted using a depth-first approach: given a vertex, one adjacent vertex
and all of its descendants need to be fully explored before the next adjacent vertex
(and all of its descendants) can be explored. After all of its adjacent vertices are
explored, then a vertex can be marked as fully explored.

During its execution, DFS will record and assign both discovery and finish times to each vertex,
which can later be used to infer ancestor/descendant relationships between vertices in the forest.

Certain implementations of DFS (like this one) can also classify edges (in addition to tree edges)
according to the DF forest generated by the algorithm:
  - Forward Edges: Connect an ancestor to a descendant in the same tree.
  - Back Edges: Connect a descendant to an ancestor in the same tree.
  - Cross Edges: Either connect vertices in different trees,
  	or non ancestor/descendant vertices in the same tree.

Expectations:
	- The graph is correctly built.

Complexity:
	- Time:  Θ(V + E)
	- Space (without edge classification): Θ(V)
	- Space (wit edge classification): Θ(V) + O(E)
*/
func DFS[V ds.Item](g *ds.G[V], classify bool) (DFForest[V], *EdgeTypes[V], error) {
	var visit func(*V)

	fst := DFForest[V]{}
	tps := &EdgeTypes[V]{}
	t := 0

	for v := range g.E {
		fst[v] = &DFNode[V]{}
	}

	visit = func(vtx *V) {
		t++

		fst[vtx].Discovery = t
		fst[vtx].visited = true

		for _, e := range g.E[vtx] {
			if fst[e.Dst].visited {
				if !classify {
					continue
				}

				if g.Directed() {
					classifyDirectedEdge(fst, tps, e)
				} else {
					classifyUndirectedEdge(fst, tps, e)
				}
			} else {
				fst[e.Dst].Parent = vtx
				visit(e.Dst)
			}
		}

		t++

		fst[vtx].Finish = t
	}

	for _, vert := range g.V {
		if !fst[vert.Ptr].visited {
			visit(vert.Ptr)
		}
	}

	return fst, tps, nil
}
