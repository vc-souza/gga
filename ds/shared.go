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
type Group struct {
	Items []Item
	Id    int
}

func (z Group) Label() string {
	return strconv.Itoa(z.Id)
}

// TODO: docs, based on:
/*
?????????????? removes the element at a given index, from a slice of pointers.
If an element from a slice of pointers is removed using the usual way of deleting an
element from a slice:

	s = append(s[:idx], s[id+1:]...)

then we risk a memory leak, from the now unreachable reference that sits in
the underlying array used by the slice, preventing garbage collection.

Source: https://github.com/golang/go/wiki/SliceTricks

???????????????/
*/
func Cut[T any](s *[]T, idx int) {
	if idx < 0 || idx >= len(*s) {
		return
	}

	copy((*s)[idx:], (*s)[idx+1:])

	// avoiding memory leak by assigning the
	// zero value to the duplicated position
	var zero T
	(*s)[len(*s)-1] = zero

	*s = (*s)[:len(*s)-1]
}
