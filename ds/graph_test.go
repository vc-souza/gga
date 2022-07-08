package ds

import (
	"testing"

	ut "github.com/vc-souza/gga/internal/testutils"
)

type GraphGen func() *Graph[ut.ID]
type edgeList []GraphEdge[ut.ID]
type vertList []*ut.ID

var GraphGenFuncs = map[string]GraphGen{
	"graph":   NewUndirectedGraph[ut.ID],
	"digraph": NewDirectedGraph[ut.ID],
}

var vA = ut.ID("a")
var vB = ut.ID("b")
var vC = ut.ID("c")
var vD = ut.ID("d")

type CounterGraphVisitor struct {
	gCalls int
	vCalls int
	eCalls int
}

func (c *CounterGraphVisitor) VisitGraphStart(g *Graph[ut.ID]) {
	c.gCalls++
}

func (c *CounterGraphVisitor) VisitGraphEnd(g *Graph[ut.ID]) {
	c.gCalls++
}

func (c *CounterGraphVisitor) VisitVertex(v *GraphVertex[ut.ID]) {
	c.vCalls++
}

func (c *CounterGraphVisitor) VisitEdge(e *GraphEdge[ut.ID]) {
	c.eCalls++
}

func tag(gtype, desc string) string {
	return gtype + " " + desc
}

func edge(src, dst *ut.ID) GraphEdge[ut.ID] {
	return GraphEdge[ut.ID]{Src: src, Dst: dst}
}

func TestNewDirectedGraph(t *testing.T) {
	g := NewDirectedGraph[ut.ID]()

	if g == nil {
		t.Log("got a nil graph")
		t.FailNow()
	}

	if g.Undirected() {
		t.Log("asked for a directed graph, got an undirected one")
	}
}

func TestNewUndirectedGraph(t *testing.T) {
	g := NewDirectedGraph[ut.ID]()

	if g == nil {
		t.Log("got a nil graph")
		t.FailNow()
	}

	if g.Directed() {
		t.Log("asked for an undirected graph, got a directed one")
	}
}

func TestVertexCount(t *testing.T) {
	cases := []struct {
		desc   string
		verts  vertList
		expect int
	}{
		{
			desc:   "no vertices",
			verts:  vertList{},
			expect: 0,
		},
		{
			desc:   "one vertex",
			verts:  vertList{&vA},
			expect: 1,
		},
	}

	for _, tc := range cases {
		for gtype, f := range GraphGenFuncs {
			t.Run(tag(gtype, tc.desc), func(t *testing.T) {
				g := f()

				g.AddVertex(tc.verts...)

				ut.AssertEqual(t, tc.expect, g.VertexCount())
			})
		}
	}
}

func TestEdgeCount(t *testing.T) {
	cases := []struct {
		desc   string
		edges  edgeList
		expect int
	}{
		{
			desc:   "zero edges",
			edges:  edgeList{},
			expect: 0,
		},
		{
			desc:   "one edge",
			edges:  edgeList{edge(&vA, &vB)},
			expect: 1,
		},
	}

	for _, tc := range cases {
		for gtype, f := range GraphGenFuncs {
			t.Run(tag(gtype, tc.desc), func(t *testing.T) {
				g := f()

				for _, e := range tc.edges {
					g.AddWeightedEdge(e.Src, e.Dst, e.Wt)
				}

				ut.AssertEqual(t, tc.expect, g.EdgeCount())
			})
		}
	}
}

func TestVertexExists(t *testing.T) {
	cases := []struct {
		desc   string
		verts  vertList
		vert   *ut.ID
		expect bool
	}{
		{
			desc:   "exists",
			verts:  vertList{&vA},
			vert:   &vA,
			expect: true,
		},
		{
			desc:   "does not exist",
			verts:  vertList{&vA},
			vert:   &vB,
			expect: false,
		},
		{
			desc:   "nil vertex",
			verts:  vertList{&vA},
			vert:   nil,
			expect: false,
		},
	}

	for _, tc := range cases {
		for gtype, f := range GraphGenFuncs {
			t.Run(tag(gtype, tc.desc), func(t *testing.T) {
				g := f()

				g.AddVertex(tc.verts...)

				ut.AssertEqual(t, tc.expect, g.VertexExists(tc.vert))
			})
		}
	}
}

func TestGetEdge(t *testing.T) {
	cases := []struct {
		desc   string
		verts  vertList
		edges  edgeList
		edge   GraphEdge[ut.ID]
		expect bool
	}{
		{
			desc:   "exists",
			verts:  vertList{&vA, &vB},
			edges:  edgeList{edge(&vA, &vB)},
			edge:   edge(&vA, &vB),
			expect: true,
		},
		{
			desc:   "does not exist (src)",
			verts:  vertList{&vA, &vB},
			edges:  edgeList{edge(&vA, &vB)},
			edge:   edge(&vC, &vB),
			expect: false,
		},
		{
			desc:   "does not exist (nil src)",
			verts:  vertList{&vA, &vB},
			edges:  edgeList{edge(&vA, &vB)},
			edge:   edge(nil, &vB),
			expect: false,
		},
		{
			desc:   "does not exist (dst)",
			verts:  vertList{&vA, &vB},
			edges:  edgeList{edge(&vA, &vB)},
			edge:   edge(&vA, &vC),
			expect: false,
		},
		{
			desc:   "does not exist (nil dst)",
			verts:  vertList{&vA, &vB},
			edges:  edgeList{edge(&vA, &vB)},
			edge:   edge(&vA, nil),
			expect: false,
		},
	}

	for _, tc := range cases {
		for gtype, f := range GraphGenFuncs {
			t.Run(tag(gtype, tc.desc), func(t *testing.T) {
				g := f()

				g.AddVertex(tc.verts...)

				for _, e := range tc.edges {
					g.AddWeightedEdge(e.Src, e.Dst, e.Wt)
				}

				_, ok := g.GetEdge(tc.edge.Src, tc.edge.Dst)

				ut.AssertEqual(t, tc.expect, ok)
			})
		}
	}
}

func TestGetVertex(t *testing.T) {
	cases := []struct {
		desc   string
		verts  vertList
		vert   *ut.ID
		expect bool
	}{
		{
			desc:   "exists",
			verts:  vertList{&vA},
			vert:   &vA,
			expect: true,
		},
		{
			desc:   "does not exist",
			verts:  vertList{},
			vert:   &vA,
			expect: false,
		},
	}

	for _, tc := range cases {
		for gtype, f := range GraphGenFuncs {
			t.Run(tag(gtype, tc.desc), func(t *testing.T) {
				g := f()

				g.AddVertex(tc.verts...)

				vert, ok := g.GetVertex(tc.vert)

				ut.AssertEqual(t, tc.expect, ok)

				if vert != nil {
					ut.AssertEqual(t, tc.vert, vert.Sat)
				}
			})
		}
	}
}

func TestAddVertex(t *testing.T) {
	cases := []struct {
		desc   string
		verts  vertList
		expect int
	}{
		{
			desc:   "no vertices",
			verts:  vertList{},
			expect: 0,
		},
		{
			desc:   "nil vertex",
			verts:  vertList{nil},
			expect: 0,
		},
		{
			desc:   "unique calls",
			verts:  vertList{&vA, &vB, &vC},
			expect: 3,
		},
		{
			desc:   "duplicated calls",
			verts:  vertList{&vA, &vA, &vB, &vB, &vB},
			expect: 2,
		},
	}

	for _, tc := range cases {
		for gtype, f := range GraphGenFuncs {
			t.Run(tag(gtype, tc.desc), func(t *testing.T) {
				g := f()

				g.AddVertex(tc.verts...)

				ut.AssertEqual(t, tc.expect, g.VertexCount())
			})
		}
	}
}

func TestAddWeightedEdge(t *testing.T) {
	cases := []struct {
		desc        string
		verts       vertList
		edges       edgeList
		edge        GraphEdge[ut.ID]
		expectEdges bool
		expectCount int
	}{
		{
			desc:        "new edge",
			verts:       vertList{&vA, &vB},
			edges:       edgeList{},
			edge:        edge(&vA, &vB),
			expectEdges: true,
			expectCount: 1,
		},
		{
			desc:        "existing edge",
			verts:       vertList{&vA, &vB},
			edges:       edgeList{edge(&vA, &vB)},
			edge:        edge(&vA, &vB),
			expectEdges: true,
			expectCount: 1,
		},
		{
			desc:        "nil src",
			verts:       vertList{&vA, &vB},
			edges:       edgeList{},
			edge:        edge(nil, &vB),
			expectEdges: false,
			expectCount: 0,
		},
		{
			desc:        "nil dst",
			verts:       vertList{&vA, &vB},
			edges:       edgeList{},
			edge:        edge(&vA, nil),
			expectEdges: false,
			expectCount: 0,
		},
	}

	for _, tc := range cases {
		for gtype, f := range GraphGenFuncs {
			t.Run(tag(gtype, tc.desc), func(t *testing.T) {
				g := f()

				g.AddVertex(tc.verts...)

				for _, e := range tc.edges {
					g.AddWeightedEdge(e.Src, e.Dst, e.Wt)
				}

				g.AddWeightedEdge(tc.edge.Src, tc.edge.Dst, tc.edge.Wt)

				if tc.expectEdges {
					_, ok := g.GetEdge(tc.edge.Src, tc.edge.Dst)

					ut.AssertEqual(t, true, ok)

					if g.Undirected() {
						_, ok := g.GetEdge(tc.edge.Dst, tc.edge.Src)

						ut.AssertEqual(t, true, ok)
					}
				}

				ut.AssertEqual(t, tc.expectCount, g.EdgeCount())
			})
		}
	}
}

func TestAddEdge(t *testing.T) {
	src := &vA
	dst := &vB

	for gtype, f := range GraphGenFuncs {
		t.Run(tag(gtype, "0 wt edge created"), func(t *testing.T) {
			g := f()

			g.AddEdge(src, dst)

			e, ok := g.GetEdge(src, dst)

			ut.AssertEqual(t, true, ok)
			ut.AssertEqual(t, 0, e.Wt)

			if g.Undirected() {
				e, ok := g.GetEdge(dst, src)

				ut.AssertEqual(t, true, ok)
				ut.AssertEqual(t, 0, e.Wt)
			}
		})
	}
}

func TestGraphVisitor(t *testing.T) {
	cases := []struct {
		desc    string
		verts   vertList
		edges   edgeList
		expectG int
		expectV int
		expectE int
	}{
		{
			desc:    "empty",
			verts:   vertList{},
			edges:   edgeList{},
			expectG: 2,
			expectV: 0,
			expectE: 0,
		},
		{
			desc:  "equal counts",
			verts: vertList{&vA, &vB, &vC, &vD},
			edges: edgeList{
				edge(&vA, &vB),
				edge(&vB, &vC),
				edge(&vC, &vD),
				edge(&vD, &vA),
			},
			expectG: 2,
			expectV: 4,
			expectE: 4,
		},
	}

	for _, tc := range cases {
		for gtype, f := range GraphGenFuncs {
			t.Run(tag(gtype, tc.desc), func(t *testing.T) {
				g := f()

				g.AddVertex(tc.verts...)

				for _, e := range tc.edges {
					g.AddWeightedEdge(e.Src, e.Dst, e.Wt)
				}

				v := CounterGraphVisitor{}

				g.Accept(&v)

				ut.AssertEqual(t, tc.expectG, v.gCalls)
				ut.AssertEqual(t, tc.expectV, v.vCalls)
				ut.AssertEqual(t, tc.expectE, v.eCalls)
			})
		}
	}
}
