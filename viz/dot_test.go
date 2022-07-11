package viz

import (
	"bytes"
	"testing"

	"github.com/vc-souza/gga/ds"
	ut "github.com/vc-souza/gga/internal/testutils"
)

var ExpectedDirectedDOT = `strict digraph {
graph [ label="A Test" ]
edge [ arrowhead="vee" ]
"Jonas"
"John" [ shape="hexagon" ]
"John" -> "Jane"
"Jane"
"Jane" -> "John"
"Jane" -> "Jane"
}
`

var ExpectedUndirectedDOT = `strict graph {
graph [ label="A Test" ]
edge [ arrowhead="vee" ]
"Jonas"
"John" [ shape="hexagon" ]
"John" -- "Jane"
"Jane"
"Jane" -- "John"
}
`

type Person struct {
	Name string
}

func (p Person) Label() string {
	return p.Name
}

func TestGraphVisitor(t *testing.T) {
	cases := []struct {
		desc   string
		gen    func() *ds.Graph[Person]
		expect string
	}{
		{
			desc:   "graph",
			gen:    ds.NewUndirectedGraph[Person],
			expect: ExpectedUndirectedDOT,
		},
		{
			desc:   "digraph",
			gen:    ds.NewDirectedGraph[Person],
			expect: ExpectedDirectedDOT,
		},
	}

	for _, tc := range cases {
		t.Run(tc.desc, func(t *testing.T) {
			g := tc.gen()
			e := NewDotExporter(g)

			john := &Person{"John"}
			jane := &Person{"Jane"}
			jonas := &Person{"Jonas"}

			e.DefaultGraphFmt = ds.FmtAttrs{
				"label": "A Test",
			}

			e.DefaultEdgeFmt = ds.FmtAttrs{
				"arrowhead": "vee",
			}

			g.AddVertex(jonas)

			vJohn, err := g.AddVertex(john)

			ut.AssertEqual(t, true, err == nil)

			vJohn.SetFmtAttr("shape", "hexagon")

			g.AddUnweightedEdge(john, jane)
			g.AddUnweightedEdge(jane, john)
			g.AddUnweightedEdge(jane, jane)

			g.Accept(e)

			buf := bytes.Buffer{}

			e.Export(&buf)

			ut.AssertEqual(t, tc.expect, buf.String())
		})
	}
}

func TestDotAttrs(t *testing.T) {
	cases := []struct {
		desc   string
		attrs  ds.FmtAttrs
		expect string
	}{
		{
			desc:   "empty",
			attrs:  ds.FmtAttrs{},
			expect: "",
		},
		{
			desc:   "single pair",
			attrs:  ds.FmtAttrs{"a": "b"},
			expect: `[ a="b" ]`,
		},
	}

	for _, tc := range cases {
		t.Run(tc.desc, func(t *testing.T) {
			ut.AssertEqual(t, tc.expect, DotAttrs(tc.attrs))
		})
	}
}
