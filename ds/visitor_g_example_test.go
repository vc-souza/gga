package ds

import "fmt"

type person struct {
	Name string
}

func (p person) Label() string {
	return p.Name
}

type ConsoleVisitor struct{}

func (cv *ConsoleVisitor) VisitGraphStart(G) {
	fmt.Println("graph start")
}

func (cv *ConsoleVisitor) VisitGraphEnd(G) {
	fmt.Println("graph end")
}

func (cv *ConsoleVisitor) VisitVertex(g G, v GV) {
	fmt.Println("vertex", v.Label())
}

func (cv *ConsoleVisitor) VisitEdge(g G, e GE) {
	fmt.Println(
		"edge,",
		g.V[e.Src].Label(),
		"to",
		g.V[e.Dst].Label(),
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

	g.Accept(&ConsoleVisitor{})

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
