package ds

import "fmt"

type person struct {
	Name string
}

func (p person) Label() string {
	return p.Name
}

type ConsoleVisitor struct{}

func (cv *ConsoleVisitor) VisitGraphStart(g *G[person]) {
	fmt.Println("graph start")
}

func (cv *ConsoleVisitor) VisitGraphEnd(*G[person]) {
	fmt.Println("graph end")
}

func (cv *ConsoleVisitor) VisitVertex(v *GV[person]) {
	fmt.Println("vertex", v.Label())
}

func (cv *ConsoleVisitor) VisitEdge(e *GE[person]) {
	fmt.Println("edge,", e.Src.Label(), "to", e.Dst.Label())
}

func ExampleGraphVisitor() {
	john := &person{"John"}
	jane := &person{"Jane"}
	jonas := &person{"Jonas"}

	g := NewDirectedGraph[person]()

	g.AddUnweightedEdge(john, jane)
	g.AddUnweightedEdge(jane, john)
	g.AddUnweightedEdge(jane, jane)

	g.AddVertex(jonas)

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
