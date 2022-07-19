package ds

// A dsfNode represents a node in a Disjoint-Set Forest.
type dsfNode[T any] struct {
	parent *dsfNode[T]
	rank   int
	ptr    *T
}

/*
DSetForest is a Disjoint-Set Forest implementation for the DSet interface,
using both the union-by-rank and path compression heuristics. Due to these
heuristics, a sequence of m operations on a DSetForest - with n of them
being calls to MakeSet - has a total amortized running time of O(mα(n)).

The function α(n) is the inverse of the Ackermann function, and grows extremely
slowly, with α(n) == 4 for pretty much any practical value of n, which makes
O(mα(n)) asymptotically superlinear, ω(m), but close to linear in practice.
*/
type DSetForest[T any] map[*T]*dsfNode[T]

// NewDSet returns a new DSet, using a Disjoint-Set Forest implementation.
func NewDSet[T any]() DSet[T] {
	return DSet[T](&DSetForest[T]{})
}

func (f DSetForest[T]) MakeSet(x *T) {
	node := &dsfNode[T]{ptr: x}
	node.parent = node
	f[x] = node
}

func (f DSetForest[T]) FindSet(x *T) *T {
	node := f[x]

	if node.parent == node {
		return node.ptr
	}

	f.compress(node)

	return node.parent.ptr
}

func (f DSetForest[T]) Union(x, y *T) {
	f.link(f.FindSet(x), f.FindSet(y))
}

// link implements the union-by-rank heuristic.
func (f DSetForest[T]) link(x, y *T) {
	parent := f[y]
	child := f[x]

	if child.rank > parent.rank {
		parent, child = child, parent
	}

	child.parent = parent

	if parent.rank == child.rank {
		parent.rank++
	}
}

// compress implements the path compression heuristic.
func (f DSetForest[T]) compress(node *dsfNode[T]) {
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
