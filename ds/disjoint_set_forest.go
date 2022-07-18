package ds

import (
	"fmt"
	"strings"
)

type dsfNode[T any] struct {
	// TODO: docs
	Parent *dsfNode[T]

	// TODO: docs
	Rank int

	// TODO: docs
	Ptr *T
}

func (n *dsfNode[T]) String() string {
	var arrow string

	if n.Parent == n {
		arrow = "↺"
	} else {
		arrow = "→ "
	}

	return fmt.Sprintf("%v[r=%d] %s", *(n.Ptr), n.Rank, arrow)
}

// TODO: docs
type DSForest[T any] map[*T]*dsfNode[T]

// TODO: docs
func NewDisjointSet[T any]() DisjointSet[T] {
	return DisjointSet[T](&DSForest[T]{})
}

func (f DSForest[T]) MakeSet(x *T) {
	node := &dsfNode[T]{Ptr: x}
	node.Parent = node
	f[x] = node
}

func (f DSForest[T]) FindSet(x *T) *T {
	node := f[x]

	if node.Parent == node {
		return node.Ptr
	}

	f.compress(node)

	return node.Parent.Ptr
}

func (f DSForest[T]) compress(node *dsfNode[T]) {
	root := node

	for root.Parent != root {
		root = root.Parent
	}

	curr := node

	for curr.Parent != root {
		parent := curr.Parent
		curr.Parent = root
		curr = parent
	}
}

func (f DSForest[T]) Union(x, y *T) {
	f.Link(f.FindSet(x), f.FindSet(y))
}

func (f DSForest[T]) Link(x, y *T) {
	parent := f[y]
	child := f[x]

	if child.Rank > parent.Rank {
		parent, child = child, parent
	}

	child.Parent = parent

	if parent.Rank == child.Rank {
		parent.Rank++
	}
}

func (f DSForest[T]) String() string {
	s := strings.Builder{}

	for _, node := range f {
		curr := node

		for {
			s.WriteString(curr.String())

			if curr.Parent == curr {
				break
			}

			curr = curr.Parent
		}

		s.WriteString("\n")
	}

	return s.String()
}
