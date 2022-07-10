package ds

import "fmt"

type state struct {
	Name string
}

func (c state) Label() string {
	return c.Name
}

func ExampleGraph_directed() {
	si := &state{"initialization"}
	sm := &state{"maintenance"}
	st := &state{"termination"}
	u1 := &state{"unrelated1"}
	u2 := &state{"unrelated2"}

	g := NewDirectedGraph[state]()

	g.AddVertex(u1)
	g.AddVertex(u2)

	g.AddEdge(si, sm)
	g.AddEdge(sm, sm)
	g.AddEdge(sm, st)

	fmt.Println(g.VertexCount())

	// directed graphs report the number of actual edges
	fmt.Println(g.EdgeCount())

	// Output:
	// 5
	// 3
}
