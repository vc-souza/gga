package ds

import "fmt"

func ExampleG_directed() {
	si := Text("initialization")
	sm := Text("maintenance")
	st := Text("termination")
	u1 := Text("unrelated1")
	u2 := Text("unrelated2")

	g := NewDigraph()

	g.AddVertex(&si)
	g.AddVertex(&sm)
	g.AddVertex(&st)
	g.AddVertex(&u1)
	g.AddVertex(&u2)

	g.AddEdge(&si, &sm, 0)
	g.AddEdge(&sm, &sm, 0)
	g.AddEdge(&sm, &st, 0)

	fmt.Println(g.VertexCount())

	fmt.Println(g.EdgeCount())

	// Output:
	// 5
	// 3
}
