package ds

import (
	"errors"
	"testing"

	ut "github.com/vc-souza/gga/internal/testutils"
)

type graphGen func() *G[Text]
type edgeList []GE[Text]
type vertList []*Text

var graphGenFuncs = map[string]graphGen{
	undirectedGraphKey: NewUndirectedGraph[Text],
	directedGraphKey:   NewDirectedGraph[Text],
}

var vA = Text("a")
var vB = Text("b")
var vC = Text("c")
var vD = Text("d")
var vE = Text("e")

type counterGraphVisitor struct {
	gCalls int
	vCalls int
	eCalls int
}

func (c *counterGraphVisitor) VisitGraphStart(g *G[Text]) {
	c.gCalls++
}

func (c *counterGraphVisitor) VisitGraphEnd(g *G[Text]) {
	c.gCalls++
}

func (c *counterGraphVisitor) VisitVertex(v *GV[Text]) {
	c.vCalls++
}

func (c *counterGraphVisitor) VisitEdge(e *GE[Text]) {
	c.eCalls++
}

func tagGraphTest(gType, desc string) string {
	return gType + " " + desc
}

func edge(src, dst *Text) GE[Text] {
	return GE[Text]{Src: src, Dst: dst}
}

func addVerts(g *G[Text], verts vertList) {
	for _, v := range verts {
		g.AddVertex(v)
	}
}

func addEdges(g *G[Text], edges edgeList) {
	for _, e := range edges {
		g.AddWeightedEdge(e.Src, e.Dst, e.Wt)

		if g.Directed() {
			continue
		}

		g.AddWeightedEdge(e.Dst, e.Src, e.Wt)
	}
}

func assertEdge(t *testing.T, g *G[Text], src, dst *Text, wt float64) {
	e, _, ok := g.GetEdge(src, dst)
	ut.True(t, ok)
	ut.Equal(t, wt, e.Wt)
}

func TestNewDirectedGraph(t *testing.T) {
	g := NewDirectedGraph[Text]()

	ut.NotNil(t, g)
	ut.True(t, g.Directed())
	ut.False(t, g.Undirected())
}

func TestNewUndirectedGraph(t *testing.T) {
	g := NewUndirectedGraph[Text]()

	ut.NotNil(t, g)
	ut.True(t, g.Undirected())
	ut.False(t, g.Directed())
}

func TestGEmptyCopy(t *testing.T) {
	for gType, f := range graphGenFuncs {
		t.Run(gType, func(t *testing.T) {
			g := f()
			cp := g.EmptyCopy()

			ut.Equal(t, g.Directed(), cp.Directed())
			ut.Equal(t, g.Undirected(), cp.Undirected())
			ut.Equal(t, 0, cp.VertexCount())
			ut.Equal(t, 0, cp.EdgeCount())
		})
	}
}

func TestGVertexCount(t *testing.T) {
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
		for gType, f := range graphGenFuncs {
			t.Run(tagGraphTest(gType, tc.desc), func(t *testing.T) {
				g := f()

				addVerts(g, tc.verts)

				ut.Equal(t, tc.expect, g.VertexCount())
			})
		}
	}
}

func TestGEdgeCount(t *testing.T) {
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
		for gType, f := range graphGenFuncs {
			t.Run(tagGraphTest(gType, tc.desc), func(t *testing.T) {
				g := f()

				addEdges(g, tc.edges)

				if g.Directed() {
					ut.Equal(t, tc.expect, g.EdgeCount())
				} else {
					ut.Equal(t, tc.expect*2, g.EdgeCount())
				}
			})
		}
	}
}

func TestGVertexExists(t *testing.T) {
	cases := []struct {
		desc   string
		verts  vertList
		vert   *Text
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
		for gType, f := range graphGenFuncs {
			t.Run(tagGraphTest(gType, tc.desc), func(t *testing.T) {
				g := f()

				addVerts(g, tc.verts)

				ut.Equal(t, tc.expect, g.VertexExists(tc.vert))
			})
		}
	}
}

func TestGGetEdge(t *testing.T) {
	cases := []struct {
		desc   string
		verts  vertList
		edges  edgeList
		edge   GE[Text]
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
		for gType, f := range graphGenFuncs {
			t.Run(tagGraphTest(gType, tc.desc), func(t *testing.T) {
				g := f()

				addVerts(g, tc.verts)
				addEdges(g, tc.edges)

				_, _, ok := g.GetEdge(tc.edge.Src, tc.edge.Dst)

				ut.Equal(t, tc.expect, ok)
			})
		}
	}
}

func TestGGetVertex(t *testing.T) {
	cases := []struct {
		desc   string
		verts  vertList
		vert   *Text
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
		for gType, f := range graphGenFuncs {
			t.Run(tagGraphTest(gType, tc.desc), func(t *testing.T) {
				g := f()

				addVerts(g, tc.verts)

				vert, _, ok := g.GetVertex(tc.vert)

				ut.Equal(t, tc.expect, ok)

				if vert != nil {
					ut.Equal(t, tc.vert, vert.Ptr)
				}
			})
		}
	}
}

func TestGAddVertex(t *testing.T) {
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
		for gType, f := range graphGenFuncs {
			t.Run(tagGraphTest(gType, tc.desc), func(t *testing.T) {
				g := f()

				addVerts(g, tc.verts)

				ut.Equal(t, tc.expect, g.VertexCount())
			})
		}
	}
}

func TestGAddWeightedEdge(t *testing.T) {
	cases := []struct {
		desc        string
		verts       vertList
		edges       edgeList
		edge        GE[Text]
		skipDir     bool
		skipUndir   bool
		expectErr   error
		expectEdges bool
		expectCount int
	}{
		{
			desc:        "new edge",
			verts:       vertList{&vA, &vB},
			edges:       edgeList{},
			edge:        edge(&vA, &vB),
			expectEdges: true,
		},
		{
			desc:        "existing edge",
			verts:       vertList{&vA, &vB},
			edges:       edgeList{edge(&vA, &vB)},
			edge:        edge(&vA, &vB),
			expectErr:   ErrExists,
			expectEdges: true,
		},
		{
			desc:        "nil src",
			verts:       vertList{&vA, &vB},
			edges:       edgeList{},
			edge:        edge(nil, &vB),
			expectErr:   ErrNilArg,
			expectEdges: false,
		},
		{
			desc:        "nil dst",
			verts:       vertList{&vA, &vB},
			edges:       edgeList{},
			edge:        edge(&vA, nil),
			expectErr:   ErrNilArg,
			expectEdges: false,
		},
		{
			desc:        "self-loop",
			verts:       vertList{&vA},
			edges:       edgeList{},
			edge:        edge(&vA, &vA),
			skipDir:     true,
			expectErr:   ErrInvLoop,
			expectEdges: false,
		},
		{
			desc:        "self-loop",
			verts:       vertList{&vA},
			edges:       edgeList{},
			edge:        edge(&vA, &vA),
			skipUndir:   true,
			expectEdges: false,
		},
	}

	for _, tc := range cases {
		for gType, f := range graphGenFuncs {
			if tc.skipDir && gType == directedGraphKey {
				continue
			}

			if tc.skipUndir && gType == undirectedGraphKey {
				continue
			}

			t.Run(tagGraphTest(gType, tc.desc), func(t *testing.T) {
				g := f()

				addVerts(g, tc.verts)
				addEdges(g, tc.edges)

				_, err := g.AddWeightedEdge(tc.edge.Src, tc.edge.Dst, tc.edge.Wt)

				ut.Equal(t, tc.expectErr == nil, err == nil)

				if tc.expectErr != nil {
					ut.True(t, errors.Is(err, tc.expectErr))
				}

				if !tc.expectEdges {
					return
				}

				assertEdge(t, g, tc.edge.Src, tc.edge.Dst, tc.edge.Wt)
			})
		}
	}
}

func TestGAddUnweightedEdge(t *testing.T) {
	src := &vA
	dst := &vB

	for gType, f := range graphGenFuncs {
		t.Run(tagGraphTest(gType, "0 wt edge created"), func(t *testing.T) {
			g := f()

			_, err := g.AddUnweightedEdge(src, dst)

			ut.Nil(t, err)

			assertEdge(t, g, src, dst, 0)
		})
	}
}

func TestGRemoveVertex(t *testing.T) {
	cases := []struct {
		desc        string
		verts       vertList
		edges       edgeList
		vert        *Text
		err         error
		expectVerts vertList
		expectEdges edgeList
	}{
		{
			desc:  "does not exist",
			verts: vertList{&vA, &vC},
			edges: edgeList{
				edge(&vA, &vC),
				edge(&vC, &vA),
			},
			vert:        &vB,
			err:         ErrDoesNotExist,
			expectVerts: vertList{&vA, &vC},
			expectEdges: edgeList{
				edge(&vA, &vC),
				edge(&vC, &vA),
			},
		},
		{
			desc:  "vertex at the start",
			verts: vertList{&vA, &vB, &vC},
			edges: edgeList{
				edge(&vA, &vB),
				edge(&vA, &vC),
				edge(&vB, &vA),
				edge(&vB, &vC),
				edge(&vC, &vA),
				edge(&vC, &vB),
			},
			vert:        &vA,
			expectVerts: vertList{&vB, &vC},
			expectEdges: edgeList{
				edge(&vB, &vC),
				edge(&vC, &vB),
			},
		},
		{
			desc:  "vertex at the middle",
			verts: vertList{&vA, &vB, &vC},
			edges: edgeList{
				edge(&vA, &vB),
				edge(&vA, &vC),
				edge(&vB, &vA),
				edge(&vB, &vC),
				edge(&vC, &vA),
				edge(&vC, &vB),
			},
			vert:        &vB,
			expectVerts: vertList{&vA, &vC},
			expectEdges: edgeList{
				edge(&vA, &vC),
				edge(&vC, &vA),
			},
		},
		{
			desc:  "vertex at the end",
			verts: vertList{&vA, &vB, &vC},
			edges: edgeList{
				edge(&vA, &vB),
				edge(&vA, &vC),
				edge(&vB, &vA),
				edge(&vB, &vC),
				edge(&vC, &vA),
				edge(&vC, &vB),
			},
			vert:        &vC,
			expectVerts: vertList{&vA, &vB},
			expectEdges: edgeList{
				edge(&vA, &vB),
				edge(&vB, &vA),
			},
		},
	}

	for _, tc := range cases {
		for gType, f := range graphGenFuncs {
			t.Run(tagGraphTest(gType, tc.desc), func(t *testing.T) {
				var ok bool

				g := f()

				addVerts(g, tc.verts)
				addEdges(g, tc.edges)

				err := g.RemoveVertex(tc.vert)

				ut.Equal(t, tc.err == nil, err == nil)

				if tc.err != nil {
					ut.True(t, errors.Is(err, tc.err))
				}

				// adjacency list removed
				_, ok = g.E[tc.vert]
				ut.False(t, ok)

				// vert mapping removed
				_, ok = g.vMap[tc.vert]
				ut.False(t, ok)

				ut.Equal(t, len(tc.expectVerts), g.VertexCount())
				ut.Equal(t, len(tc.expectEdges), g.EdgeCount())

				// vertices correctly rearranged, indexes updated
				for i := 0; i < len(tc.expectVerts); i++ {
					expected := tc.expectVerts[i]
					actual := g.V[i]

					// correct item at the correct position
					ut.Equal(t, expected, actual.Ptr)

					// correct mapping for the item
					ut.Equal(t, i, g.vMap[actual.Ptr])
				}

				// correct edges still in place
				for _, e := range tc.expectEdges {
					_, _, ok := g.GetEdge(e.Src, e.Dst)
					ut.True(t, ok)
				}
			})
		}
	}
}

func TestGRemoveEdge(t *testing.T) {
	cases := []struct {
		desc             string
		verts            vertList
		edges            edgeList
		edge             GE[Text]
		exists           bool
		err              bool
		expectDirCount   int
		expectUndirCount int
	}{
		{
			desc:             "does not exist",
			verts:            vertList{&vA, &vB},
			edges:            edgeList{},
			edge:             edge(&vA, &vB),
			err:              true,
			expectDirCount:   0,
			expectUndirCount: 0,
		},
		{
			desc:  "common edge",
			verts: vertList{&vA, &vB, &vC},
			edges: edgeList{
				edge(&vA, &vB),
				edge(&vA, &vC),
				edge(&vB, &vA),
				edge(&vB, &vC),
				edge(&vC, &vA),
				edge(&vC, &vB),
			},
			edge:             edge(&vA, &vB),
			exists:           true,
			err:              false,
			expectDirCount:   5,
			expectUndirCount: 4,
		},
	}

	for _, tc := range cases {
		for gType, f := range graphGenFuncs {
			t.Run(tagGraphTest(gType, tc.desc), func(t *testing.T) {
				g := f()

				addVerts(g, tc.verts)
				addEdges(g, tc.edges)

				if tc.exists {
					_, _, ok := g.GetEdge(tc.edge.Src, tc.edge.Dst)
					ut.True(t, ok)
				}

				err := g.RemoveEdge(tc.edge.Src, tc.edge.Dst)

				ut.Equal(t, tc.err, err != nil)

				if g.Directed() {
					ut.Equal(t, tc.expectDirCount, g.EdgeCount())
				} else {
					ut.Equal(t, tc.expectUndirCount, g.EdgeCount())
				}

				_, _, ok := g.GetEdge(tc.edge.Src, tc.edge.Dst)
				ut.False(t, ok)

				if g.Directed() {
					return
				}

				_, _, ok = g.GetEdge(tc.edge.Dst, tc.edge.Src)
				ut.False(t, ok)
			})
		}
	}
}

func TestGVisitor(t *testing.T) {
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
		for gType, f := range graphGenFuncs {
			t.Run(tagGraphTest(gType, tc.desc), func(t *testing.T) {
				g := f()

				addVerts(g, tc.verts)
				addEdges(g, tc.edges)

				v := counterGraphVisitor{}

				g.Accept(&v)

				ut.Equal(t, tc.expectG, v.gCalls)
				ut.Equal(t, tc.expectV, v.vCalls)

				if g.Directed() {
					ut.Equal(t, tc.expectE, v.eCalls)
				} else {
					ut.Equal(t, tc.expectE*2, v.eCalls)
				}
			})
		}
	}
}

func TestGTranspose_directed(t *testing.T) {
	//                     (loop)
	// B -> C -> D -> E -> A
	//      |---------^
	g := NewDirectedGraph[Text]()

	g.AddWeightedEdge(&vA, &vA, 1)
	g.AddWeightedEdge(&vB, &vC, 2)
	g.AddWeightedEdge(&vC, &vD, 3)
	g.AddWeightedEdge(&vC, &vE, 4)
	g.AddWeightedEdge(&vD, &vE, 5)
	g.AddWeightedEdge(&vE, &vA, 6)

	//                     (loop)
	// B <- C <- D <- E <- A
	//      ^---------|
	tp, err := g.Transpose()

	ut.Nil(t, err)
	ut.Equal(t, g.VertexCount(), tp.VertexCount())
	ut.Equal(t, g.EdgeCount(), tp.EdgeCount())

	ut.True(t, tp.VertexExists(&vA))
	ut.True(t, tp.VertexExists(&vB))
	ut.True(t, tp.VertexExists(&vC))
	ut.True(t, tp.VertexExists(&vD))
	ut.True(t, tp.VertexExists(&vE))

	assertEdge(t, tp, &vA, &vA, 1)
	assertEdge(t, tp, &vC, &vB, 2)
	assertEdge(t, tp, &vD, &vC, 3)
	assertEdge(t, tp, &vE, &vC, 4)
	assertEdge(t, tp, &vE, &vD, 5)
	assertEdge(t, tp, &vA, &vE, 6)
}

func TestGTranspose_undirected(t *testing.T) {
	g := NewUndirectedGraph[Text]()
	_, err := g.Transpose()

	ut.NotNil(t, err)
	ut.True(t, errors.Is(err, ErrUndefOp))
}
