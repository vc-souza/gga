package viz

import (
	"bytes"
	"sort"
	"strings"
	"testing"

	"github.com/vc-souza/gga/ds"
	ut "github.com/vc-souza/gga/internal/testutils"
)

var ExpectedDirectedDOT = `
strict digraph {
graph [ label="A Test" ]
node [ shape="folder" ]
edge [ arrowhead="vee" ]
Jonas
John [ shape="hexagon" ]
John -> Jane
Jane
Jane -> John
Jane -> Jane
}`

var ExpectedUndirectedDOT = `
strict graph {
graph [ label="A Test" ]
node [ shape="folder" ]
edge [ arrowhead="vee" ]
Jonas
John [ shape="hexagon" ]
John -- Jane
Jane
Jane -- John
Jane -- Jane
}`

type Person struct {
	Name string
}

func (p Person) Label() string {
	return p.Name
}

func sortedString(s string) string {
	ss := strings.Split(s, "")
	sort.Strings(ss)
	return strings.Join(ss, "")
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
			de := NewDotExporter(g)

			john := &Person{"John"}
			jane := &Person{"Jane"}
			jonas := &Person{"Jonas"}

			de.DefaultGraphFmt = ds.FmtAttrs{
				"label": "A Test",
			}

			de.DefaultVertexFmt = ds.FmtAttrs{
				"shape": "folder",
			}

			de.DefaultEdgeFmt = ds.FmtAttrs{
				"arrowhead": "vee",
			}

			g.AddVertex(jonas)
			g.AddEdge(john, jane)
			g.AddEdge(jane, john)
			g.AddEdge(jane, jane)

			johnV, ok := g.GetVertex(john)

			ut.AssertEqual(t, true, ok)

			johnV.SetFmtAttr("shape", "hexagon")

			g.Accept(de)

			buf := bytes.Buffer{}

			de.Export(&buf)

			ut.AssertEqual(t, sortedString(tc.expect), sortedString(buf.String()))
		})
	}
}
