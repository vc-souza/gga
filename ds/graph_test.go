package ds

import (
	"errors"
	"testing"

	ut "github.com/vc-souza/gga/internal/testutils"
)

var graphGen = map[string]func() *G{
	undirectedGraphKey: NewGraph,
	directedGraphKey:   NewDigraph,
}

var vA = Text("a")
var vB = Text("b")
var vC = Text("c")
var vD = Text("d")
var vE = Text("e")

type edge struct {
	src Item
	dst Item
	wt  float64
}

type counterGraphVisitor struct{ gCalls, vCalls, eCalls int }

func (c *counterGraphVisitor) VisitGraphStart(G) { c.gCalls++ }
func (c *counterGraphVisitor) VisitGraphEnd(G)   { c.gCalls++ }
func (c *counterGraphVisitor) VisitVertex(G, GV) { c.vCalls++ }
func (c *counterGraphVisitor) VisitEdge(G, GE)   { c.eCalls++ }

func tagGraphTest(gType, desc string) string {
	return gType + " " + desc
}

func addVerts(t *testing.T, g *G, is ...Item) {
	for _, item := range is {
		_, err := g.AddVertex(item)
		ut.Nil(t, err)
	}
}

func addEdges(t *testing.T, g *G, es ...edge) {
	var err error

	for _, e := range es {
		_, _, err = g.AddEdge(e.src, e.dst, e.wt)
		ut.Nil(t, err)

		if g.Directed() {
			continue
		}

		_, _, err = g.AddEdge(e.dst, e.src, e.wt)
		ut.Nil(t, err)
	}
}

func assertEdge(t *testing.T, g *G, e edge) {
	iV, iE, ok := g.EdgeIndex(e.src, e.dst)
	ut.True(t, ok)
	ut.Equal(t, e.wt, g.V[iV].E[iE].Wt)
}

func TestGNewGraph(t *testing.T) {
	g := NewGraph()
	ut.True(t, g.Undirected())
	ut.False(t, g.Directed())
}

func TestGNewDigraph(t *testing.T) {
	g := NewDigraph()
	ut.True(t, g.Directed())
	ut.False(t, g.Undirected())
}

func TestGDirected(t *testing.T) {
	cases := []struct {
		desc   string
		gen    func() *G
		expect bool
	}{
		{
			desc:   undirectedGraphKey,
			gen:    NewGraph,
			expect: false,
		},
		{
			desc:   directedGraphKey,
			gen:    NewDigraph,
			expect: true,
		},
	}

	for _, tc := range cases {
		t.Run(tc.desc, func(t *testing.T) {
			g := tc.gen()
			ut.Equal(t, tc.expect, g.Directed())
		})
	}
}

func TestGUndirected(t *testing.T) {
	cases := []struct {
		desc   string
		gen    func() *G
		expect bool
	}{
		{
			desc:   undirectedGraphKey,
			gen:    NewGraph,
			expect: true,
		},
		{
			desc:   directedGraphKey,
			gen:    NewDigraph,
			expect: false,
		},
	}

	for _, tc := range cases {
		t.Run(tc.desc, func(t *testing.T) {
			g := tc.gen()
			ut.Equal(t, tc.expect, g.Undirected())
		})
	}
}

func TestGVertexCount(t *testing.T) {
	cases := []struct {
		desc   string
		verts  []Item
		expect int
	}{
		{
			desc:   "zero vertices",
			expect: 0,
		},
		{
			desc:   "one vertex",
			verts:  []Item{vA},
			expect: 1,
		},
		{
			desc:   "multiple vertices",
			verts:  []Item{vA, vB, vC},
			expect: 3,
		},
	}

	for _, tc := range cases {
		for gType, gen := range graphGen {
			t.Run(tagGraphTest(gType, tc.desc), func(t *testing.T) {
				g := gen()

				addVerts(t, g, tc.verts...)

				ut.Equal(t, tc.expect, g.VertexCount())
			})
		}
	}
}

func TestGEdgeCount(t *testing.T) {
	cases := []struct {
		desc   string
		verts  []Item
		edges  []edge
		expect int
	}{
		{
			desc:   "zero edges",
			expect: 0,
		},
		{
			desc:   "one edge",
			verts:  []Item{vA, vB},
			edges:  []edge{{vA, vB, 0}},
			expect: 1,
		},
		{
			desc:   "multiple edges",
			verts:  []Item{vA, vB, vC},
			edges:  []edge{{vA, vB, 0}, {vA, vC, 0}, {vB, vC, 0}},
			expect: 3,
		},
	}

	for _, tc := range cases {
		for gType, gen := range graphGen {
			t.Run(tagGraphTest(gType, tc.desc), func(t *testing.T) {
				g := gen()

				addVerts(t, g, tc.verts...)
				addEdges(t, g, tc.edges...)

				if g.Undirected() {
					ut.Equal(t, tc.expect*2, g.EdgeCount())
				} else {
					ut.Equal(t, tc.expect, g.EdgeCount())
				}
			})
		}
	}
}

func TestGVertexIndex(t *testing.T) {
	cases := []struct {
		desc     string
		verts    []Item
		item     Item
		expectV  int
		expectOK bool
	}{
		{
			desc: "zero vertices",
			item: vA,
		},
		{
			desc:     "one vertex",
			verts:    []Item{vA},
			item:     vA,
			expectV:  0,
			expectOK: true,
		},
		{
			desc:     "multiple vertices",
			verts:    []Item{vA, vB, vC},
			item:     vB,
			expectV:  1,
			expectOK: true,
		},
	}

	for _, tc := range cases {
		for gType, gen := range graphGen {
			t.Run(tagGraphTest(gType, tc.desc), func(t *testing.T) {
				g := gen()

				addVerts(t, g, tc.verts...)

				v, ok := g.VertexIndex(tc.item)

				ut.Equal(t, tc.expectOK, ok)
				ut.Equal(t, tc.expectV, v)
			})
		}
	}
}

func TestGEdgeIndex(t *testing.T) {
	cases := []struct {
		desc     string
		verts    []Item
		edges    []edge
		input    edge
		expectV  int
		expectE  int
		expectOK bool
	}{
		{
			desc:  "zero edges",
			input: edge{vA, vB, 0},
		},
		{
			desc:     "one edge",
			verts:    []Item{vA, vB},
			edges:    []edge{{vA, vB, 0}},
			input:    edge{vA, vB, 0},
			expectV:  0,
			expectE:  0,
			expectOK: true,
		},
		{
			desc:     "multiple edges",
			verts:    []Item{vA, vB, vC},
			edges:    []edge{{vB, vA, 0}, {vB, vC, 0}},
			input:    edge{vB, vC, 0},
			expectV:  1,
			expectE:  1,
			expectOK: true,
		},
	}

	for _, tc := range cases {
		for gType, gen := range graphGen {
			t.Run(tagGraphTest(gType, tc.desc), func(t *testing.T) {
				g := gen()

				addVerts(t, g, tc.verts...)
				addEdges(t, g, tc.edges...)

				v, e, ok := g.EdgeIndex(tc.input.src, tc.input.dst)

				ut.Equal(t, tc.expectOK, ok)
				ut.Equal(t, tc.expectV, v)
				ut.Equal(t, tc.expectE, e)
			})
		}
	}
}

func TestGAddVertex(t *testing.T) {
	cases := []struct {
		desc      string
		verts     []Item
		item      Item
		expectV   int
		expectErr error
	}{
		{
			desc: "new vertex",
			item: vA,
		},
		{
			desc:      "duplicated vertex",
			verts:     []Item{vA},
			item:      vA,
			expectV:   0,
			expectErr: ErrExists,
		},
	}

	for _, tc := range cases {
		for gType, gen := range graphGen {
			t.Run(tagGraphTest(gType, tc.desc), func(t *testing.T) {
				g := gen()

				addVerts(t, g, tc.verts...)

				v, err := g.AddVertex(tc.item)

				if tc.expectErr == nil {
					ut.Nil(t, err)
				} else {
					ut.True(t, errors.Is(err, tc.expectErr))
				}

				ut.Equal(t, tc.expectV, v)
			})
		}
	}
}

func TestGAddEdge(t *testing.T) {
	cases := []struct {
		desc        string
		digraphOnly bool
		graphOnly   bool
		verts       []Item
		edges       []edge
		input       edge
		expectV     int
		expectE     int
		expectErr   error
	}{
		{
			desc:  "zero edges",
			verts: []Item{vA, vB},
			input: edge{vA, vB, 0},
		},
		{
			desc:    "one more edge",
			verts:   []Item{vA, vB, vC},
			edges:   []edge{{vA, vB, 0}},
			input:   edge{vA, vC, 0},
			expectV: 0,
			expectE: 1,
		},
		{
			desc:      "src does not exist",
			verts:     []Item{vB},
			input:     edge{vA, vB, 0},
			expectV:   0,
			expectE:   0,
			expectErr: ErrNoVtx,
		},
		{
			desc:      "dst does not exist",
			verts:     []Item{vA},
			input:     edge{vA, vB, 0},
			expectV:   0,
			expectE:   0,
			expectErr: ErrNoVtx,
		},
		{
			desc:      "invalid loop",
			graphOnly: true,
			verts:     []Item{vA},
			input:     edge{vA, vA, 0},
			expectV:   0,
			expectE:   0,
			expectErr: ErrInvLoop,
		},
	}

	for _, tc := range cases {
		for gType, gen := range graphGen {
			if tc.digraphOnly && gType == undirectedGraphKey {
				continue
			}

			if tc.graphOnly && gType == directedGraphKey {
				continue
			}

			t.Run(tagGraphTest(gType, tc.desc), func(t *testing.T) {
				g := gen()

				addVerts(t, g, tc.verts...)
				addEdges(t, g, tc.edges...)

				v, e, err := g.AddEdge(tc.input.src, tc.input.dst, tc.input.wt)

				if tc.expectErr == nil {
					ut.Nil(t, err)
				} else {
					ut.True(t, errors.Is(err, tc.expectErr))
				}

				ut.Equal(t, tc.expectV, v)
				ut.Equal(t, tc.expectE, e)
			})
		}
	}
}
