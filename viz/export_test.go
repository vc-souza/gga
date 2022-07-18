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
		gen    func() *ds.G[person]
		expect string
	}{
		{
			desc:   "graph",
			gen:    ds.NewUndirectedGraph[person],
			expect: ExpectedUndirectedDOT,
		},
		{
			desc:   "digraph",
			gen:    ds.NewDirectedGraph[person],
			expect: ExpectedDirectedDOT,
		},
	}

	for _, tc := range cases {
		t.Run(tc.desc, func(t *testing.T) {
			g := tc.gen()
			e := NewExporter(g)

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

			vJohn, err := g.AddVertex(john)

			ut.Equal(t, true, err == nil)

			vJohn.SetFmtAttr("shape", "hexagon")

			g.AddUnweightedEdge(john, jane)
			g.AddUnweightedEdge(jane, john)
			g.AddUnweightedEdge(jane, jane)

			buf := bytes.Buffer{}

			e.Export(&buf)

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
			expect: `[ a="b" ]`,
		},
	}

	for _, tc := range cases {
		t.Run(tc.desc, func(t *testing.T) {
			ut.Equal(t, tc.expect, DotAttrs(tc.attrs))
		})
	}
}

func TestResetGraphFmt(t *testing.T) {
	isClear := func(g *ds.G[person]) bool {
		for _, vtx := range g.V {
			if len(vtx.F) != 0 {
				return false
			}
		}

		for _, es := range g.E {
			for _, e := range es {
				if len(e.F) != 0 {
					return false
				}
			}
		}

		return true
	}

	cases := []struct {
		desc   string
		gen    func() *ds.G[person]
		expect string
	}{
		{
			desc:   "graph",
			gen:    ds.NewUndirectedGraph[person],
			expect: ExpectedUndirectedDOT,
		},
		{
			desc:   "digraph",
			gen:    ds.NewDirectedGraph[person],
			expect: ExpectedDirectedDOT,
		},
	}

	for _, tc := range cases {
		t.Run(tc.desc, func(t *testing.T) {
			g := tc.gen()

			ut.Equal(t, true, isClear(g))

			john := &person{"John"}
			jane := &person{"Jane"}

			vJohn, err := g.AddVertex(john)

			ut.Equal(t, true, err == nil)

			vJohn.SetFmtAttr("label", "John is here")

			vJane, err := g.AddVertex(jane)

			ut.Equal(t, true, err == nil)

			vJane.SetFmtAttr("label", "Jane is here")

			edg, err := g.AddUnweightedEdge(john, jane)

			ut.Equal(t, true, err == nil)

			edg.SetFmtAttr("label", "Connection")

			ut.Equal(t, false, isClear(g))

			ResetGraphFmt(g)

			ut.Equal(t, true, isClear(g))
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
	g, _, err := ds.NewTextParser().Parse(`
	digraph
	a#b,c
	b#d
	c#
	d#b
	`)

	ut.Equal(t, true, err == nil)

	buf := bytes.Buffer{}

	Snapshot(g, &buf, exportTestTheme{})

	ut.Equal(t, expectedSnapshot, buf.String())
}
