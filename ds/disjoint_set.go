package ds

import "fmt"

// TODO: docs
type DisjointSet[T any] interface {
	fmt.Stringer

	// TODO: docs
	MakeSet(x *T)

	// TODO: docs
	FindSet(x *T) *T

	// TODO: docs
	Union(x, y *T)

	// TODO: docs
	Link(x, y *T)
}
