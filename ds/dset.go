package ds

/*
DSet implementations are able to behave like a Disjoint-Set data structure,
storing and operating on a collection of disjoint sets.
*/
type DSet[T any] interface {
	/*
		MakeSet creates a new disjoint set containing a single element.
	*/
	MakeSet(x *T)

	/*
		FindSet, given an element, finds the representative element of the disjoint set
		that contains the original element; a behavior that can be used to test if two
		elements belong to the same set: if so, they should have the same representative.
	*/
	FindSet(x *T) *T

	/*
		Union merges the disjoint sets containing the given elements, creating a new set.
	*/
	Union(x, y *T)
}
