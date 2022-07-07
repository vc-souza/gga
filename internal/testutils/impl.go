package testutils

// ID is a test struct that implements the graph.Vertex interface.
type ID string

func (i ID) Label() string {
	return string(i)
}
