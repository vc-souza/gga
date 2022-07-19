package ds

import "strconv"

/*
An Item implementation can be used as satellite data for an item in a gga data structure.
The main feature of an Item is being able to provide a label for easy identification.
Some use cases would be logging and the generation of data visualizations using tools like Graphviz.
*/
type Item interface {
	Label() string
}

/*
Group groups items of a type that implements the Item interface, and also implements
the Item interface itself, using an id assigned during the creation of the Group,
so data structures that can hold Item implementations can also hold Group values.

Such a capability is useful for some algorithms that group items together and then
create a new data structure that holds the groups as new elements (e.g.: GSCC).
*/
type Group[V Item] struct {
	Items []*V
	Id    int
}

func (z Group[V]) Label() string {
	return strconv.Itoa(z.Id)
}

/*
RemoveFromPointersSlice removes the element at a given index, from a slice of pointers.
If an element from a slice of pointers is removed using the usual way of deleting an
element from a slice:

	s = append(s[:idx], s[id+1:]...)

then we risk a memory leak, from the now unreachable reference that sits in
the underlying array used by the slice, preventing garbage collection.

Source: https://github.com/golang/go/wiki/SliceTricks
*/
func RemoveFromPointersSlice[T any](s []*T, idx int) []*T {
	if idx < 0 || idx >= len(s) {
		return s
	}

	// Overwrite the element to be deleted by copying the remainder of
	// the slice over it. Now the last element of the slice is duplicated:
	// the same pointer exists both in the old position and in the position
	// that needed to be deleted.
	copy(s[idx:], s[idx+1:])

	// Delete the extra reference to the pointer by assigning nil to
	// its old position.
	s[len(s)-1] = nil

	// Shrink the slice by slicing it again and then return the new, shorter
	// slice to the caller. Just like with 'append', the caller needs to store
	// the new slice reference where its old slice used to be stored.
	return s[:len(s)-1]
}
