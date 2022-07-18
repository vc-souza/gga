package ds

// A dsfNode represents a node in a Disjoint-Set Forest.
type dsfNode[T any] struct {
	parent *dsfNode[T]
	rank   int
	ptr    *T
}

/*
DisjointSetForest is a Disjoint-Set Forest implementation for the DisjointSet interface,
using both the union-by-rank and path compression heuristics. Due to these heuristics,
a sequence of m operations on a DisjointSetForest - with n of them being calls
to MakeSet - has a total amortized running time of O(mα(n)).

The function α(n) is the inverse of the Ackermann function, and grows extremely
slowly, with α(n) == 4 for pretty much any practical value of n, which makes
O(mα(n)) asymptotically superlinear, ω(m), but very close to linear in practice.
*/
type DisjointSetForest[T any] map[*T]*dsfNode[T]

// NewDisjointSet returns a new DisjointSet, using a Disjoint-Set Forest implementation.
func NewDisjointSet[T any]() DisjointSet[T] {
	return DisjointSet[T](&DisjointSetForest[T]{})
}

func (f DisjointSetForest[T]) MakeSet(x *T) {
	node := &dsfNode[T]{ptr: x}
	node.parent = node
	f[x] = node
}

func (f DisjointSetForest[T]) FindSet(x *T) *T {
	node := f[x]

	if node.parent == node {
		return node.ptr
	}

	// path compression heuristic
	f.compress(node)

	return node.parent.ptr
}

func (f DisjointSetForest[T]) Union(x, y *T) {
	f.Link(f.FindSet(x), f.FindSet(y))
}

func (f DisjointSetForest[T]) Link(x, y *T) {
	parent := f[y]
	child := f[x]

	// union-by-rank heuristic
	if child.rank > parent.rank {
		parent, child = child, parent
	}

	child.parent = parent

	if parent.rank == child.rank {
		parent.rank++
	}
}

func (f DisjointSetForest[T]) compress(node *dsfNode[T]) {
	root := node

	for root.parent != root {
		root = root.parent
	}

	curr := node

	for curr.parent != root {
		parent := curr.parent
		curr.parent = root
		curr = parent
	}
}
