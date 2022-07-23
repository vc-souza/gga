package ds

import "fmt"

type person struct {
	Name string
}

func (p person) Label() string {
	return p.Name
}

type ConsoleVisitor struct {
	g *G
}

func (cv *ConsoleVisitor) VisitGraphStart(G) {
	fmt.Println("graph start")
}

func (cv *ConsoleVisitor) VisitGraphEnd(G) {
	fmt.Println("graph end")
}

func (cv *ConsoleVisitor) VisitVertex(g G, v int) {
	fmt.Println("vertex", g.V[v].Item.Label())
}

func (cv *ConsoleVisitor) VisitEdge(g G, v int, e int) {
	fmt.Println(
		"edge,",
		g.V[cv.g.V[v].E[e].Src].Item.Label(),
		"to",
		g.V[cv.g.V[v].E[e].Dst].Item.Label(),
	)
}

func ExampleGraphVisitor() {
	john := &person{"John"}
	jane := &person{"Jane"}
	jonas := &person{"Jonas"}

	g := NewDigraph()

	g.AddVertex(john)
	g.AddVertex(jane)
	g.AddVertex(jonas)

	g.AddEdge(john, jane, 0)
	g.AddEdge(jane, john, 0)
	g.AddEdge(jane, jane, 0)

	g.Accept(&ConsoleVisitor{g})

	// Output:
	// graph start
	// vertex John
	// edge, John to Jane
	// vertex Jane
	// edge, Jane to John
	// edge, Jane to Jane
	// vertex Jonas
	// graph end
}
