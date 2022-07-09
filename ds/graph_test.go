package ds

import (
	"testing"

	ut "github.com/vc-souza/gga/internal/testutils"
)

type GraphGen func() *Graph[ut.ID]
type edgeList []GraphEdge[ut.ID]
type vertList []*ut.ID

const (
	UndirectedGraphKey = "graph"
	DirectedGraphKey   = "digraph"
)

var GraphGenFuncs = map[string]GraphGen{
	UndirectedGraphKey: NewUndirectedGraph[ut.ID],
	DirectedGraphKey:   NewDirectedGraph[ut.ID],
}

var vA = ut.ID("a")
var vB = ut.ID("b")
var vC = ut.ID("c")
var vD = ut.ID("d")
var vE = ut.ID("e")

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

func assertEdge(t *testing.T, g *Graph[ut.ID], src, dst *ut.ID, wt float64) {
	e, _, ok := g.GetEdge(src, dst)
	ut.AssertEqual(t, true, ok)
	ut.AssertEqual(t, wt, e.Wt)
}

func TestNewDirectedGraph(t *testing.T) {
	g := NewDirectedGraph[ut.ID]()

	ut.AssertEqual(t, true, g != nil)
	ut.AssertEqual(t, true, g.Directed())
	ut.AssertEqual(t, false, g.Undirected())
}

func TestNewUndirectedGraph(t *testing.T) {
	g := NewUndirectedGraph[ut.ID]()

	ut.AssertEqual(t, true, g != nil)
	ut.AssertEqual(t, true, g.Undirected())
	ut.AssertEqual(t, false, g.Directed())
}

func TestEmptyCopy(t *testing.T) {
	for gtype, f := range GraphGenFuncs {
		t.Run(gtype, func(t *testing.T) {
			g := f()
			cp := g.EmptyCopy()

			ut.AssertEqual(t, g.Directed(), cp.Directed())
			ut.AssertEqual(t, g.Undirected(), cp.Undirected())
			ut.AssertEqual(t, 0, cp.VertexCount())
			ut.AssertEqual(t, 0, cp.EdgeCount())
		})
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

				_, _, ok := g.GetEdge(tc.edge.Src, tc.edge.Dst)

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

				vert, _, ok := g.GetVertex(tc.vert)

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
		skipDir     bool
		skipUndir   bool
		expectErr   bool
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
			expectErr:   true,
			expectEdges: false,
			expectCount: 0,
		},
		{
			desc:        "nil dst",
			verts:       vertList{&vA, &vB},
			edges:       edgeList{},
			edge:        edge(&vA, nil),
			expectErr:   true,
			expectEdges: false,
			expectCount: 0,
		},
		{
			desc:        "self-loop",
			verts:       vertList{&vA},
			edges:       edgeList{},
			edge:        edge(&vA, &vA),
			skipDir:     true,
			expectErr:   true,
			expectEdges: false,
			expectCount: 0,
		},
		{
			desc:        "self-loop",
			verts:       vertList{&vA},
			edges:       edgeList{},
			edge:        edge(&vA, &vA),
			skipUndir:   true,
			expectEdges: false,
			expectCount: 1,
		},
	}

	for _, tc := range cases {
		for gtype, f := range GraphGenFuncs {
			if tc.skipDir && gtype == DirectedGraphKey {
				continue
			}

			if tc.skipUndir && gtype == UndirectedGraphKey {
				continue
			}

			t.Run(tag(gtype, tc.desc), func(t *testing.T) {
				g := f()

				g.AddVertex(tc.verts...)

				for _, e := range tc.edges {
					g.AddWeightedEdge(e.Src, e.Dst, e.Wt)
				}

				err := g.AddWeightedEdge(tc.edge.Src, tc.edge.Dst, tc.edge.Wt)

				ut.AssertEqual(t, tc.expectErr, err != nil)
				ut.AssertEqual(t, tc.expectCount, g.EdgeCount())

				if !tc.expectEdges {
					return
				}

				_, _, ok := g.GetEdge(tc.edge.Src, tc.edge.Dst)

				ut.AssertEqual(t, true, ok)

				if g.Undirected() {
					_, _, ok := g.GetEdge(tc.edge.Dst, tc.edge.Src)

					ut.AssertEqual(t, true, ok)
				}
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

			err := g.AddEdge(src, dst)

			ut.AssertEqual(t, true, err == nil)

			e, _, ok := g.GetEdge(src, dst)

			ut.AssertEqual(t, true, ok)
			ut.AssertEqual(t, 0, e.Wt)

			if g.Undirected() {
				e, _, ok := g.GetEdge(dst, src)

				ut.AssertEqual(t, true, ok)
				ut.AssertEqual(t, 0, e.Wt)
			}
		})
	}
}

func TestRemoveVertex(t *testing.T) {
	cases := []struct {
		desc        string
		verts       vertList
		edges       edgeList
		vert        *ut.ID
		expectError bool
		expectVerts vertList
		expectEdges edgeList
	}{
		{
			desc:        "does not exist",
			verts:       vertList{&vA, &vC},
			edges:       edgeList{edge(&vA, &vC)},
			vert:        &vB,
			expectError: true,
			expectVerts: vertList{&vA, &vC},
			expectEdges: edgeList{edge(&vA, &vC)},
		},
		{
			desc:  "vertex at the start",
			verts: vertList{&vA, &vB, &vC},
			edges: edgeList{
				edge(&vA, &vC),
				edge(&vB, &vA),
				edge(&vB, &vC),
			},
			vert:        &vA,
			expectVerts: vertList{&vB, &vC},
			expectEdges: edgeList{edge(&vB, &vC)},
		},
		{
			desc:  "vertex at the middle",
			verts: vertList{&vA, &vB, &vC},
			edges: edgeList{
				edge(&vA, &vC),
				edge(&vB, &vA),
				edge(&vC, &vB),
			},
			vert:        &vB,
			expectVerts: vertList{&vA, &vC},
			expectEdges: edgeList{edge(&vA, &vC)},
		},
		{
			desc:  "vertex at the end",
			verts: vertList{&vA, &vB, &vC},
			edges: edgeList{
				edge(&vA, &vB),
				edge(&vB, &vC),
				edge(&vC, &vA),
			},
			vert:        &vC,
			expectVerts: vertList{&vA, &vB},
			expectEdges: edgeList{edge(&vA, &vB)},
		},
	}

	for _, tc := range cases {
		for gtype, f := range GraphGenFuncs {
			t.Run(tag(gtype, tc.desc), func(t *testing.T) {
				var ok bool

				g := f()

				g.AddVertex(tc.verts...)

				for _, e := range tc.edges {
					g.AddWeightedEdge(e.Src, e.Dst, e.Wt)
				}

				err := g.RemoveVertex(tc.vert)

				ut.AssertEqual(t, tc.expectError, err != nil)

				// adjacency list removed
				_, ok = g.Adj[tc.vert]
				ut.AssertEqual(t, false, ok)

				// vert mapping removed
				_, ok = g.VertMap[tc.vert]
				ut.AssertEqual(t, false, ok)

				ut.AssertEqual(t, len(tc.expectVerts), g.VertexCount())
				ut.AssertEqual(t, len(tc.expectEdges), g.EdgeCount())

				// vertices correctly rearranged, indexes updated
				for i := 0; i < len(tc.expectVerts); i++ {
					expected := tc.expectVerts[i]
					actual := g.Verts[i]

					// correct item at the correct position
					ut.AssertEqual(t, expected, actual.Sat)

					// correct mapping for the item
					ut.AssertEqual(t, i, g.VertMap[actual.Sat])
				}

				// correct edges still in place
				for _, e := range tc.expectEdges {
					_, _, ok := g.GetEdge(e.Src, e.Dst)
					ut.AssertEqual(t, true, ok)
				}
			})
		}
	}
}

func TestRemoveEdge(t *testing.T) {
	cases := []struct {
		desc        string
		verts       vertList
		edges       edgeList
		edge        GraphEdge[ut.ID]
		exists      bool
		expectError bool
		expectCount int
	}{
		{
			desc:        "does not exist",
			verts:       vertList{&vA, &vB},
			edges:       edgeList{},
			edge:        edge(&vA, &vB),
			expectError: true,
			expectCount: 0,
		},
		{
			desc:        "last edge",
			verts:       vertList{&vA, &vB},
			edges:       edgeList{edge(&vA, &vB)},
			edge:        edge(&vA, &vB),
			exists:      true,
			expectError: false,
			expectCount: 0,
		},
		{
			desc:        "common edge",
			verts:       vertList{&vA, &vB, &vC},
			edges:       edgeList{edge(&vA, &vB), edge(&vA, &vC), edge(&vB, &vC)},
			edge:        edge(&vA, &vB),
			exists:      true,
			expectError: false,
			expectCount: 2,
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

				if tc.exists {
					_, _, ok := g.GetEdge(tc.edge.Src, tc.edge.Dst)
					ut.AssertEqual(t, true, ok)
				}

				err := g.RemoveEdge(tc.edge.Src, tc.edge.Dst)

				ut.AssertEqual(t, tc.expectError, err != nil)
				ut.AssertEqual(t, tc.expectCount, g.EdgeCount())

				_, _, ok := g.GetEdge(tc.edge.Src, tc.edge.Dst)
				ut.AssertEqual(t, false, ok)

				if g.Directed() {
					return
				}

				_, _, ok = g.GetEdge(tc.edge.Dst, tc.edge.Src)
				ut.AssertEqual(t, false, ok)
			})
		}
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
			},
			expectG: 2,
			expectV: 4,
			expectE: 3,
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

				if g.Directed() {
					ut.AssertEqual(t, tc.expectE, v.eCalls)
				} else {
					ut.AssertEqual(t, tc.expectE*2, v.eCalls)
				}
			})
		}
	}
}

func TestTranspose_directed(t *testing.T) {
	//                     (loop)
	// B -> C -> D -> E -> A
	//      |---------^
	g := NewDirectedGraph[ut.ID]()

	g.AddWeightedEdge(&vA, &vA, 1)
	g.AddWeightedEdge(&vB, &vC, 2)
	g.AddWeightedEdge(&vC, &vD, 3)
	g.AddWeightedEdge(&vC, &vE, 4)
	g.AddWeightedEdge(&vD, &vE, 5)
	g.AddWeightedEdge(&vE, &vA, 6)

	//                     (loop)
	// B <- C <- D <- E <- A
	//      ^---------|
	tp := g.Transpose()

	ut.AssertEqual(t, g.VertexCount(), tp.VertexCount())
	ut.AssertEqual(t, g.EdgeCount(), tp.EdgeCount())

	ut.AssertEqual(t, true, tp.VertexExists(&vA))
	ut.AssertEqual(t, true, tp.VertexExists(&vB))
	ut.AssertEqual(t, true, tp.VertexExists(&vC))
	ut.AssertEqual(t, true, tp.VertexExists(&vD))
	ut.AssertEqual(t, true, tp.VertexExists(&vE))

	assertEdge(t, tp, &vA, &vA, 1)
	assertEdge(t, tp, &vC, &vB, 2)
	assertEdge(t, tp, &vD, &vC, 3)
	assertEdge(t, tp, &vE, &vC, 4)
	assertEdge(t, tp, &vE, &vD, 5)
	assertEdge(t, tp, &vA, &vE, 6)
}

func TestTranspose_undirected(t *testing.T) {
	g := NewUndirectedGraph[ut.ID]()

	g.AddWeightedEdge(&vA, &vB, 1)
	g.AddWeightedEdge(&vB, &vC, 2)
	g.AddWeightedEdge(&vC, &vD, 3)
	g.AddWeightedEdge(&vD, &vE, 4)
	g.AddWeightedEdge(&vE, &vA, 5)

	tp := g.Transpose()

	ut.AssertEqual(t, g.VertexCount(), tp.VertexCount())
	ut.AssertEqual(t, g.EdgeCount(), tp.EdgeCount())

	assertEdge(t, tp, &vA, &vB, 1)
	assertEdge(t, tp, &vB, &vA, 1)
	assertEdge(t, tp, &vB, &vC, 2)
	assertEdge(t, tp, &vC, &vB, 2)
	assertEdge(t, tp, &vC, &vD, 3)
	assertEdge(t, tp, &vD, &vC, 3)
	assertEdge(t, tp, &vD, &vE, 4)
	assertEdge(t, tp, &vE, &vD, 4)
	assertEdge(t, tp, &vE, &vA, 5)
	assertEdge(t, tp, &vA, &vE, 5)
}
