package ds

import "fmt"

func ExampleG_directed() {
	si := Text("initialization")
	sm := Text("maintenance")
	st := Text("termination")
	u1 := Text("unrelated1")
	u2 := Text("unrelated2")

	g := NewDirectedGraph[Text]()

	g.AddVertex(&u1)
	g.AddVertex(&u2)

	g.AddUnweightedEdge(&si, &sm)
	g.AddUnweightedEdge(&sm, &sm)
	g.AddUnweightedEdge(&sm, &st)

	fmt.Println(g.VertexCount())

	// directed graphs report the number of actual edges
	fmt.Println(g.EdgeCount())

	// Output:
	// 5
	// 3
}
