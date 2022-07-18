package ds

/*
DisjointSet implementations are able to operate on a collection of disjoint sets.
*/
type DisjointSet[T any] interface {
	// MakeSet creates a new disjoint set containing a single element.
	MakeSet(x *T)

	/*
		FindSet, given an element, finds the representative element of the disjoint set
		that contains the original element. This behavior can be used to test if two
		elements belong to the same set: if so, their representatives should be the same.
	*/
	FindSet(x *T) *T

	// Union merges the sets that contain the two given elements, creating a new set.
	Union(x, y *T)
}
