package graph

import (
	"testing"

	tu "github.com/vc-souza/gga/internal/testutils"
)

type GraphGen func() *Graph[tu.ID]

var GraphGenFuncs = []GraphGen{
	NewDirectedGraph[tu.ID],
	NewUndirectedGraph[tu.ID],
}

func TestNewDirectedGraph(t *testing.T) {
	g := NewDirectedGraph[tu.ID]()

	if g == nil {
		t.Log("got a nil graph")
		t.FailNow()
	}

	if !g.Directed() {
		t.Log("asked for a directed graph, got an undirected one")
	}
}

func TestNewUndirectedGraph(t *testing.T) {
	g := NewDirectedGraph[tu.ID]()

	if g == nil {
		t.Log("got a nil graph")
		t.FailNow()
	}

	if g.Directed() {
		t.Log("asked for an undirected graph, got a directed one")
	}
}

func TestVertexExists(t *testing.T) {
	a := tu.ID("a")
	b := tu.ID("b")

	cases := []struct {
		desc   string
		verts  []*tu.ID
		vert   *tu.ID
		expect bool
	}{
		{
			desc:   "exists",
			verts:  []*tu.ID{&a},
			vert:   &a,
			expect: true,
		},
		{
			desc:   "does not exist",
			verts:  []*tu.ID{&a},
			vert:   &b,
			expect: false,
		},
		{
			desc:   "nil vertex",
			verts:  []*tu.ID{&a},
			vert:   nil,
			expect: false,
		},
	}

	for _, tc := range cases {
		for _, f := range GraphGenFuncs {
			t.Run(tc.desc, func(t *testing.T) {
				g := f()

				g.AddVertex(tc.verts...)

				tu.AssertEqual(t, tc.expect, g.VertexExists(tc.vert))
			})
		}
	}
}

func TestAddVertex(t *testing.T) {
	a := tu.ID("a")
	b := tu.ID("b")
	c := tu.ID("c")

	cases := []struct {
		desc   string
		verts  []*tu.ID
		expect int
	}{
		{
			desc:   "no vertices",
			verts:  []*tu.ID{},
			expect: 0,
		},
		{
			desc:   "nil vertex",
			verts:  []*tu.ID{nil},
			expect: 0,
		},
		{
			desc:   "unique calls",
			verts:  []*tu.ID{&a, &b, &c},
			expect: 3,
		},
		{
			desc:   "duplicated calls",
			verts:  []*tu.ID{&a, &a, &b, &b, &b},
			expect: 2,
		},
	}

	for _, tc := range cases {
		for _, f := range GraphGenFuncs {
			t.Run(tc.desc, func(t *testing.T) {
				g := f()

				g.AddVertex(tc.verts...)

				tu.AssertEqual(t, tc.expect, g.CountVertices())
			})
		}
	}
}

func TestEdgeExists(t *testing.T) {
	a := tu.ID("a")
	b := tu.ID("b")
	c := tu.ID("c")

	wt := 1.
	aToB := Edge[tu.ID]{&a, &b, wt}

	cases := []struct {
		desc   string
		verts  []*tu.ID
		edges  []Edge[tu.ID]
		edge   Edge[tu.ID]
		expect bool
	}{
		{
			desc:   "exists",
			verts:  []*tu.ID{&a, &b},
			edges:  []Edge[tu.ID]{aToB},
			edge:   Edge[tu.ID]{&a, &b, wt},
			expect: true,
		},
		{
			desc:   "does not exist (src)",
			verts:  []*tu.ID{&a, &b},
			edges:  []Edge[tu.ID]{aToB},
			edge:   Edge[tu.ID]{&c, &b, wt},
			expect: false,
		},
		{
			desc:   "does not exist (nil src)",
			verts:  []*tu.ID{&a, &b},
			edges:  []Edge[tu.ID]{aToB},
			edge:   Edge[tu.ID]{nil, &b, wt},
			expect: false,
		},
		{
			desc:   "does not exist (dst)",
			verts:  []*tu.ID{&a, &b},
			edges:  []Edge[tu.ID]{aToB},
			edge:   Edge[tu.ID]{&a, &c, wt},
			expect: false,
		},
		{
			desc:   "does not exist (nil dst)",
			verts:  []*tu.ID{&a, &b},
			edges:  []Edge[tu.ID]{aToB},
			edge:   Edge[tu.ID]{&a, nil, wt},
			expect: false,
		},
		{
			desc:   "does not exist (wt)",
			verts:  []*tu.ID{&a, &b},
			edges:  []Edge[tu.ID]{aToB},
			edge:   Edge[tu.ID]{&a, &b, wt + 1},
			expect: false,
		},
	}

	for _, tc := range cases {
		for _, f := range GraphGenFuncs {
			t.Run(tc.desc, func(t *testing.T) {
				g := f()

				g.AddVertex(tc.verts...)

				for _, e := range tc.edges {
					g.AddWeightedEdge(e.Src, e.Dst, e.Wt)
				}

				tu.AssertEqual(t, tc.expect, g.EdgeExists(tc.edge.Src, tc.edge.Dst, tc.edge.Wt))
			})
		}
	}
}

func TestAddWeightedEdge(t *testing.T) {
	a := tu.ID("a")
	b := tu.ID("b")

	wt := 1.
	aToB := Edge[tu.ID]{&a, &b, wt}

	cases := []struct {
		desc        string
		verts       []*tu.ID
		edges       []Edge[tu.ID]
		edge        Edge[tu.ID]
		expectEdges bool
		expectCount int
	}{
		{
			desc:        "new edge",
			verts:       []*tu.ID{&a, &b},
			edges:       []Edge[tu.ID]{},
			edge:        Edge[tu.ID]{&a, &b, wt},
			expectEdges: true,
			expectCount: 1,
		},
		{
			desc:        "existing edge, same wt",
			verts:       []*tu.ID{&a, &b},
			edges:       []Edge[tu.ID]{aToB},
			edge:        Edge[tu.ID]{&a, &b, wt},
			expectEdges: true,
			expectCount: 1,
		},
		{
			desc:        "existing edge, different wt",
			verts:       []*tu.ID{&a, &b},
			edges:       []Edge[tu.ID]{aToB},
			edge:        Edge[tu.ID]{&a, &b, wt + 1},
			expectEdges: true,
			expectCount: 2,
		},
		{
			desc:        "nil src",
			verts:       []*tu.ID{&a, &b},
			edges:       []Edge[tu.ID]{},
			edge:        Edge[tu.ID]{nil, &b, wt},
			expectEdges: false,
			expectCount: 0,
		},
		{
			desc:        "nil dst",
			verts:       []*tu.ID{&a, &b},
			edges:       []Edge[tu.ID]{},
			edge:        Edge[tu.ID]{&a, nil, wt},
			expectEdges: false,
			expectCount: 0,
		},
	}

	for _, tc := range cases {
		for _, f := range GraphGenFuncs {
			t.Run(tc.desc, func(t *testing.T) {
				g := f()

				g.AddVertex(tc.verts...)

				for _, e := range tc.edges {
					g.AddWeightedEdge(e.Src, e.Dst, e.Wt)
				}

				g.AddWeightedEdge(tc.edge.Src, tc.edge.Dst, tc.edge.Wt)

				if tc.expectEdges {
					tu.AssertEqual(t, true, g.EdgeExists(tc.edge.Src, tc.edge.Dst, tc.edge.Wt))

					if !g.Directed() {
						tu.AssertEqual(t, true, g.EdgeExists(tc.edge.Dst, tc.edge.Src, tc.edge.Wt))
					}
				}

				tu.AssertEqual(t, tc.expectCount, g.CountEdges())
			})
		}
	}
}

func TestAddEdge_directed(t *testing.T) {
	a := tu.ID("a")
	b := tu.ID("b")

	src := &a
	dst := &b

	g := NewDirectedGraph[tu.ID]()

	g.AddEdge(src, dst)

	tu.AssertEqual(t, true, g.EdgeExists(src, dst, 0))
}

func TestAddEdge_undirected(t *testing.T) {
	a := tu.ID("a")
	b := tu.ID("b")

	src := &a
	dst := &b

	g := NewUndirectedGraph[tu.ID]()

	g.AddEdge(src, dst)

	tu.AssertEqual(t, true, g.EdgeExists(src, dst, 0))
	tu.AssertEqual(t, true, g.EdgeExists(dst, src, 0))
}
