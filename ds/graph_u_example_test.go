package ds

import "fmt"

type city struct {
	Name string
}

func (c city) Label() string {
	return c.Name
}

func ExampleGraph_undirected() {
	wt := &city{"Whiterun"}
	dt := &city{"Dawnstar"}
	mt := &city{"Markarth"}
	rt := &city{"Riften"}

	g := NewUndirectedGraph[city]()

	// explicitly adding a vertex that does not participate in any edges
	g.AddVertex(rt)

	// vertices that are part of an edge do not have to be added explicitly
	// the reverse edge does not have to be added explicitly
	g.AddEdge(wt, dt)
	g.AddEdge(wt, mt)
	g.AddEdge(dt, mt)

	fmt.Println(g.VertexCount())

	// undirected graphs report half the numbers of actual edges, since reverse edges are excluded
	fmt.Println(g.EdgeCount())

	// Output:
	// 4
	// 3
}
