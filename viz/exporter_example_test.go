package viz

import (
	"fmt"
	"os"

	"github.com/vc-souza/gga/ds"
)

type pokemon struct {
	Type string
	Name string
}

func (c pokemon) Label() string {
	return fmt.Sprintf("%s (%s)", c.Name, c.Type)
}

func ExampleExporter() {
	pwag := &pokemon{Type: "water", Name: "Poliwag"}
	pwl := &pokemon{Type: "water", Name: "Poliwhirl"}
	pot := &pokemon{Type: "water", Name: "Politoed"}
	pow := &pokemon{Type: "water/fighting", Name: "Poliwrath"}

	g := ds.NewDigraph()
	e := NewExporter()

	g.AddVertex(pwag)
	g.AddVertex(pwl)
	g.AddVertex(pot)
	g.AddVertex(pow)

	g.AddEdge(pwag, pwl, 0)
	g.AddEdge(pwl, pot, 0)
	g.AddEdge(pwl, pow, 0)

	e.DefaultGraphFmt = ds.FAttrs{
		"rankdir": "LR",
	}

	e.DefaultVertexFmt = ds.FAttrs{
		"style": "filled",
	}

	e.DefaultEdgeFmt = ds.FAttrs{
		"arrowhead": "open",
	}

	if i, ok := g.GetVertexIndex(pwag); ok {
		g.V[i].SetFmtAttr("shape", "square")
	}

	if v, e, ok := g.GetEdgeIndex(pwag, pwl); ok {
		g.V[v].E[e].SetFmtAttr("label", "Level 25")
	}

	if v, e, ok := g.GetEdgeIndex(pwl, pot); ok {
		g.V[v].E[e].SetFmtAttr("label", "Trade holding King's Rock")
	}

	if v, e, ok := g.GetEdgeIndex(pwl, pow); ok {
		g.V[v].E[e].SetFmtAttr("label", "Water Stone")
	}

	e.Export(g, os.Stdout)

	// Output:
	// strict digraph {
	// graph [ rankdir="LR" ]
	// node [ style="filled" ]
	// edge [ arrowhead="open" ]
	// "Poliwag (water)" [ shape="square" ]
	// "Poliwag (water)" -> "Poliwhirl (water)" [ label="Level 25" ]
	// "Poliwhirl (water)"
	// "Poliwhirl (water)" -> "Politoed (water)" [ label="Trade holding King's Rock" ]
	// "Poliwhirl (water)" -> "Poliwrath (water/fighting)" [ label="Water Stone" ]
	// "Politoed (water)"
	// "Poliwrath (water/fighting)"
	// }
	//
}
