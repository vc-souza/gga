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

func (cv *ConsoleVisitor) VisitGraphStart(g *G) {
	fmt.Println("graph start")
}

func (cv *ConsoleVisitor) VisitGraphEnd(*G) {
	fmt.Println("graph end")
}

func (cv *ConsoleVisitor) VisitVertex(v *GV) {
	fmt.Println("vertex", v.Item.Label())
}

func (cv *ConsoleVisitor) VisitEdge(e *GE) {
	fmt.Println("edge,", cv.g.V[e.Src].Item.Label(), "to", cv.g.V[e.Dst].Item.Label())
}

func ExampleGraphVisitor() {
	john := &person{"John"}
	jane := &person{"Jane"}
	jonas := &person{"Jonas"}

	g := NewDigraph()

	g.AddEdge(john, jane, 0)
	g.AddEdge(jane, john, 0)
	g.AddEdge(jane, jane, 0)

	g.AddVertex(jonas)

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
