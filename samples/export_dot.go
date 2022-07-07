package main

import "github.com/vc-souza/gga/ds"

type Person struct {
	Name string
	Age  int
}

// Necessary to be used as satellite data of a graph vertex
func (p Person) Label() string {
	return p.Name
}

func main() {
	jimmy := Person{"Jimmy", 30}
	joe := Person{"Joe", 20}
	mary := Person{"Mary", 25}

	g := ds.NewDirectedGraph[Person]()

	g.AddVertex(&mary)

	g.AddEdge(&jimmy, &joe)
	g.AddEdge(&joe, &jimmy)

	// TODO: dot stuff
}
