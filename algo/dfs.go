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
	// Discovery records when the vertex was marked as discovered.
	Discovery int

	// Discovery records when the vertex was marked as fully explored.
	Finish int

	/*
		Color holds the current color of the vertex.
			- White vertices are undiscovered, but at the end of the DFS, no vertex will remain so.
			- Gray vertices are discovered and either being explored or fully explored.
	*/
	Color int

	/*
		Parent holds the vertex that discovered this vertex, with the edge (v.Parent, v) being called a tree edge.
		This is how the DF tree that this vertex is a part of is encoded: by following the parent pointers from
		this vertex, one can get to the root of the DF tree.

		After a DFS, every root of a DF tree in the DF forest will have a nil Parent.
	*/
	Parent *V

	/*
		next is used to keep track of which one of the vertices that are adjacent to this vertex
		should be processed next, whenever this vertex gets picked up again for processing.
	*/
	next int
}

/*
A DFForest is one of the results of a DFS, representing a forest of DF trees (connected acyclic subgraphs),
with each DF tree being composed of all (at the time) white vertices that were reachable from the root
of the DF tree, during the execution of the DFS.

Slightly different trees can be generated for the same graph, if the visiting order for either vertices
or edges is changed.

The gga graph implementation guarantees both vertex and edge traversal in insertion order,
so repeated DFS calls always produce the same DF forest.
*/
type DFForest[V ds.Item] map[*V]*DFNode[V]

func classifyDirectedEdge[V ds.Item](fst DFForest[V], tps *EdgeTypes[V], e *ds.GraphEdge[V]) {
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

func classifyUndirectedEdge[V ds.Item](fst DFForest[V], tps *EdgeTypes[V], e *ds.GraphEdge[V]) {
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

Link: https://en.wikipedia.org/wiki/Depth-first_search

Expectations:
	- The graph is correctly built.

Complexity:
	- Time:  Θ(V + E)
	- Space (without edge classification): Θ(V)
	- Space (wit edge classification): O(V + E)
*/
func DFS[V ds.Item](g *ds.Graph[V], classify bool) (DFForest[V], *EdgeTypes[V], error) {
	fst := DFForest[V]{}
	tps := &EdgeTypes[V]{}
	t := 0

	for v := range g.Adj {
		fst[v] = &DFNode[V]{
			Color: ColorWhite,
		}
	}

	// build a DF tree rooted at the given vertex;
	// the tree will be a part of the DF forest
	tree := func(root *V) {
		// only using the ds.Stack interface
		stk := ds.Stack[*V](new(ds.Deque[*V]))

		stk.Push(root)

		for !stk.Empty() {
			vtx, _ := stk.Peek()

			// vertex is being discovered
			if fst[vtx].Color == ColorWhite {
				t++
				fst[vtx].Discovery = t
				fst[vtx].Color = ColorGray
			}

			// vertex has exhausted its adjacency list:
			// all of its descendants have been
			// discovered and fully explored
			if fst[vtx].next >= len(g.Adj[vtx]) {
				stk.Pop()
				t++
				fst[vtx].Finish = t

				continue
			}

			// explore what remains of the adjacency list of the vertex:
			// new nodes will be pushed to the stack and old ones will
			// trigger the classification of the edge that connects them
			for i := fst[vtx].next; i < len(g.Adj[vtx]); i++ {
				e := g.Adj[vtx][i]

				if fst[e.Dst].Color != ColorWhite {
					fst[vtx].next++

					if !classify {
						continue
					}

					if g.Directed() {
						classifyDirectedEdge(fst, tps, e)
					} else {
						classifyUndirectedEdge(fst, tps, e)
					}

					continue
				}

				// found a tree edge
				fst[e.Dst].Parent = vtx
				stk.Push(e.Dst)
				fst[vtx].next++

				// depth-first means that a descendant needs to be fully
				// explored before the next adjacent vertex is considered;
				// whenever we run out of descendants to explore, the value
				// of fst[vtx].next will give us the next adjacent node
				// to fully explore.
				break
			}
		}
	}

	// if a vertex is not included in a tree during a call to the 'tree'
	// function, then it could be picked as the root of the next tree:
	// by iterating over all white vertices, we assure that no vertex
	// will be left without being assign to a DF tree, even if its
	// tree ends up only containing the vertex itself.
	for _, vert := range g.Verts {
		root := vert.Val

		// skip: already part of another tree
		if fst[root].Color != ColorWhite {
			continue
		}

		tree(root)
	}

	return fst, tps, nil
}
