package ds

import "fmt"

func ExampleG_undirected() {
	wt := Text("Whiterun")
	dt := Text("Dawnstar")
	mt := Text("Markarth")
	rt := Text("Riften")

	g := NewGraph()

	g.AddVertex(&wt)
	g.AddVertex(&dt)
	g.AddVertex(&mt)
	g.AddVertex(&rt)

	g.AddEdge(&wt, &dt, 0)
	g.AddEdge(&dt, &wt, 0)

	g.AddEdge(&wt, &mt, 0)
	g.AddEdge(&mt, &wt, 0)

	g.AddEdge(&dt, &mt, 0)
	g.AddEdge(&mt, &dt, 0)

	fmt.Println(g.VertexCount())

	fmt.Println(g.EdgeCount())

	// Output:
	// 4
	// 6
}
