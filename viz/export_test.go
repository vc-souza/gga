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

type person struct {
	Name string
}

func (p person) Label() string {
	return p.Name
}

func TestGraphVisitor(t *testing.T) {
	cases := []struct {
		desc   string
		gen    func() *ds.G
		expect string
	}{
		{
			desc:   "graph",
			gen:    ds.NewGraph,
			expect: ExpectedUndirectedDOT,
		},
		{
			desc:   "digraph",
			gen:    ds.NewDigraph,
			expect: ExpectedDirectedDOT,
		},
	}

	for _, tc := range cases {
		t.Run(tc.desc, func(t *testing.T) {
			g := tc.gen()
			e := NewExporter()

			john := &person{"John"}
			jane := &person{"Jane"}
			jonas := &person{"Jonas"}

			e.DefaultGraphFmt = ds.FAttrs{
				"label": "A Test",
			}

			e.DefaultEdgeFmt = ds.FAttrs{
				"arrowhead": "vee",
			}

			g.AddVertex(jonas)

			iJohn, err := g.AddVertex(john)

			g.AddVertex(jane)

			ut.Nil(t, err)

			g.V[iJohn].SetFmtAttr("shape", "hexagon")

			g.AddEdge(john, jane, 0)
			g.AddEdge(jane, john, 0)
			g.AddEdge(jane, jane, 0)

			buf := bytes.Buffer{}

			e.Export(g, &buf)

			ut.Equal(t, tc.expect, buf.String())
		})
	}
}

func TestDotAttrs(t *testing.T) {
	cases := []struct {
		desc   string
		attrs  ds.FAttrs
		expect string
	}{
		{
			desc:   "empty",
			attrs:  ds.FAttrs{},
			expect: "",
		},
		{
			desc:   "single pair",
			attrs:  ds.FAttrs{"a": "b"},
			expect: ` [ a="b" ]`,
		},
	}

	for _, tc := range cases {
		t.Run(tc.desc, func(t *testing.T) {
			ut.Equal(t, tc.expect, DotAttrs(tc.attrs))
		})
	}
}

func TestResetGraphFmt(t *testing.T) {
	isClear := func(g *ds.G) bool {
		for v := range g.V {
			if len(g.V[v].F) != 0 {
				return false
			}

			for j := range g.V[v].E {
				if len(g.V[v].E[j].F) != 0 {
					return false
				}
			}
		}

		return true
	}

	cases := []struct {
		desc   string
		gen    func() *ds.G
		expect string
	}{
		{
			desc:   "graph",
			gen:    ds.NewGraph,
			expect: ExpectedUndirectedDOT,
		},
		{
			desc:   "digraph",
			gen:    ds.NewDigraph,
			expect: ExpectedDirectedDOT,
		},
	}

	for _, tc := range cases {
		t.Run(tc.desc, func(t *testing.T) {
			g := tc.gen()

			ut.True(t, isClear(g))

			john := &person{"John"}
			jane := &person{"Jane"}

			iJohn, err := g.AddVertex(john)

			ut.Nil(t, err)

			g.V[iJohn].SetFmtAttr("label", "John is here")

			iJane, err := g.AddVertex(jane)

			ut.Nil(t, err)

			g.V[iJane].SetFmtAttr("label", "Jane is here")

			iVtx, iEdge, err := g.AddEdge(john, jane, 0)

			ut.Nil(t, err)

			g.V[iVtx].E[iEdge].SetFmtAttr("label", "Connection")

			ut.False(t, isClear(g))

			ResetGraphFmt(g)

			ut.True(t, isClear(g))
		})
	}
}

type exportTestTheme struct{}

func (t exportTestTheme) SetGraphFmt(attrs ds.FAttrs) {
	attrs["layout"] = "dot"
}

func (t exportTestTheme) SetVertexFmt(attrs ds.FAttrs) {
	attrs["style"] = "filled"
}

func (t exportTestTheme) SetEdgeFmt(attrs ds.FAttrs) {
	attrs["arrowhead"] = "vee"
}

var expectedSnapshot = `strict digraph {
graph [ layout="dot" ]
node [ style="filled" ]
edge [ arrowhead="vee" ]
"a"
"a" -> "b"
"a" -> "c"
"b"
"b" -> "d"
"c"
"d"
"d" -> "b"
}
`

func TestSnapshot(t *testing.T) {
	g, _, err := ds.Parse(`
	digraph
	a#b,c
	b#d
	c#
	d#b
	`)

	ut.Nil(t, err)

	buf := bytes.Buffer{}

	Snapshot(g, &buf, exportTestTheme{})

	ut.Equal(t, expectedSnapshot, buf.String())
}
