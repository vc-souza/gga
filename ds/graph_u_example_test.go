package ds

import "fmt"

func ExampleG_undirected() {
	wt := Text("Whiterun")
	dt := Text("Dawnstar")
	mt := Text("Markarth")
	rt := Text("Riften")

	g := NewUndirectedGraph[Text]()

	// explicitly adding a vertex that does not participate in any edges
	g.AddVertex(&rt)

	// vertices that are part of an edge do not have to be added explicitly
	// since this is an undirected graph, reverse edges have to be added
	g.AddUnweightedEdge(&wt, &dt)
	g.AddUnweightedEdge(&dt, &wt)

	g.AddUnweightedEdge(&wt, &mt)
	g.AddUnweightedEdge(&mt, &wt)

	g.AddUnweightedEdge(&dt, &mt)
	g.AddUnweightedEdge(&mt, &dt)

	fmt.Println(g.VertexCount())

	// undirected graphs report half the numbers of actual edges, since reverse edges are excluded
	fmt.Println(g.EdgeCount())

	// Output:
	// 4
	// 6
}
