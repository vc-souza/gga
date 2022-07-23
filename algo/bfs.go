package algo

import (
	"math"

	"github.com/vc-souza/gga/ds"
)

/*
A BFNode represents a node in a Breadth-First tree, holding the attributes produced by a BFS,
for a particular vertex. At the end of the BFS, nodes with a distance < infinity are a part
of a BF tree, rooted at the source vertex.
*/
type BFNode struct {
	/*
		Distance is the length of the shortest path (edge count), from the source to this vertex.
		If the vertex is unreachable from the source, this value will be math.Inf(1).
	*/
	Distance float64

	/*
		Parent holds the vertex that discovered this vertex, with the edge (v.Parent, v) being called a tree edge.
		This is how the BF tree is encoded: by following the parent pointers from any reachable vertex back to
		the source, one can generate a shortest path from the source to the vertex.

		After a BFS, both the source and all unreachable vertices have a nil Parent.
	*/
	Parent int

	visited bool
}

/*
A BFTree (Breadth-First Tree) is the result of a BFS, representing a tree (connected acyclic subgraph)
rooted at the source, and containing every vertex that is reachable from the source. A BF tree encodes
both the length of the shortest path between the source and each reachable vertex (Distance)
and the path itself (Parent pointer).

Slightly different trees can be generated for the same graph and source, if the visiting order for
either vertices or edges is changed, but the optimal distances are guaranteed to remain the same.

The gga graph implementation guarantees both vertex and edge traversal in insertion order,
so repeated BFS calls always produce the same BF tree.
*/
type BFTree []BFNode

/*
BFS implements the Breadth-First Search (BFS) algorithm.

Given a graph and a source vertex, BFS explores all vertices that are reachable from the source, with the end
result being a Breadth-First tree rooted at the source. The search explores all non-explored vertices at a
certain distance d from the source before moving on to the vertices at distance d+1, working in "waves".

Expectations:
	- The graph is correctly built.
	- The source vertex exists.

Complexity:
	- Time:  Θ(V + E)
	- Space: Θ(V)
*/
func BFS(g *ds.G, src int) (BFTree, error) {
	tree := make(BFTree, g.VertexCount())
	queue := ds.NewQueue[int]()

	for i := range g.V {
		tree[i].Distance = math.Inf(1)
		tree[i].Parent = -1
	}

	tree[src].Distance = 0
	tree[src].visited = true

	queue.Enqueue(src)

	for !queue.Empty() {
		curr, _ := queue.Dequeue()

		for _, edge := range g.V[curr].E {
			if tree[edge.Dst].visited {
				continue
			}

			tree[edge.Dst].Distance = tree[curr].Distance + 1
			tree[edge.Dst].Parent = curr
			tree[edge.Dst].visited = true

			queue.Enqueue(edge.Dst)
		}
	}

	return tree, nil
}
