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
	e := NewExporter(g)

	g.AddUnweightedEdge(pwag, pwl)
	g.AddUnweightedEdge(pwl, pot)
	g.AddUnweightedEdge(pwl, pow)

	e.DefaultGraphFmt = ds.FAttrs{
		"rankdir": "LR",
	}

	e.DefaultVertexFmt = ds.FAttrs{
		"style": "filled",
	}

	e.DefaultEdgeFmt = ds.FAttrs{
		"arrowhead": "open",
	}

	if v, _, ok := g.GetVertex(pwag); ok {
		v.SetFmtAttr("shape", "square")
	}

	if e, _, ok := g.GetEdge(pwag, pwl); ok {
		e.SetFmtAttr("label", "Level 25")
	}

	if e, _, ok := g.GetEdge(pwl, pot); ok {
		e.SetFmtAttr("label", "Trade holding King's Rock")
	}

	if e, _, ok := g.GetEdge(pwl, pow); ok {
		e.SetFmtAttr("label", "Water Stone")
	}

	e.Export(os.Stdout)

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
