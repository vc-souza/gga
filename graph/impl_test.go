package graph

import (
	"testing"

	ut "github.com/vc-souza/gga/internal/testutils"
)

type GraphGen func() *Graph[ut.ID]

var GraphGenFuncs = []GraphGen{
	NewDirectedGraph[ut.ID],
	NewUndirectedGraph[ut.ID],
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
	a := ut.ID("a")

	cases := []struct {
		desc   string
		verts  []*ut.ID
		expect int
	}{
		{
			desc:   "no vertices",
			verts:  []*ut.ID{},
			expect: 0,
		},
		{
			desc:   "one vertex",
			verts:  []*ut.ID{&a},
			expect: 1,
		},
	}

	for _, tc := range cases {
		for _, f := range GraphGenFuncs {
			t.Run(tc.desc, func(t *testing.T) {
				g := f()

				g.AddVertex(tc.verts...)

				ut.AssertEqual(t, tc.expect, g.VertexCount())
			})
		}
	}
}

func TestEdgeCount(t *testing.T) {
	a := ut.ID("a")
	b := ut.ID("b")

	a2b := Edge[ut.ID]{Src: &a, Dst: &b}

	cases := []struct {
		desc   string
		edges  []Edge[ut.ID]
		expect int
	}{
		{
			desc:   "zero edges",
			edges:  []Edge[ut.ID]{},
			expect: 0,
		},
		{
			desc:   "one edge",
			edges:  []Edge[ut.ID]{a2b},
			expect: 1,
		},
	}

	for _, tc := range cases {
		for _, f := range GraphGenFuncs {
			t.Run(tc.desc, func(t *testing.T) {
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
	a := ut.ID("a")
	b := ut.ID("b")

	cases := []struct {
		desc   string
		verts  []*ut.ID
		vert   *ut.ID
		expect bool
	}{
		{
			desc:   "exists",
			verts:  []*ut.ID{&a},
			vert:   &a,
			expect: true,
		},
		{
			desc:   "does not exist",
			verts:  []*ut.ID{&a},
			vert:   &b,
			expect: false,
		},
		{
			desc:   "nil vertex",
			verts:  []*ut.ID{&a},
			vert:   nil,
			expect: false,
		},
	}

	for _, tc := range cases {
		for _, f := range GraphGenFuncs {
			t.Run(tc.desc, func(t *testing.T) {
				g := f()

				g.AddVertex(tc.verts...)

				ut.AssertEqual(t, tc.expect, g.VertexExists(tc.vert))
			})
		}
	}
}

func TestEdgeExists(t *testing.T) {
	a := ut.ID("a")
	b := ut.ID("b")
	c := ut.ID("c")

	wt := 1.
	a2b := Edge[ut.ID]{&a, &b, wt}

	cases := []struct {
		desc   string
		verts  []*ut.ID
		edges  []Edge[ut.ID]
		edge   Edge[ut.ID]
		expect bool
	}{
		{
			desc:   "exists",
			verts:  []*ut.ID{&a, &b},
			edges:  []Edge[ut.ID]{a2b},
			edge:   Edge[ut.ID]{&a, &b, wt},
			expect: true,
		},
		{
			desc:   "does not exist (src)",
			verts:  []*ut.ID{&a, &b},
			edges:  []Edge[ut.ID]{a2b},
			edge:   Edge[ut.ID]{&c, &b, wt},
			expect: false,
		},
		{
			desc:   "does not exist (nil src)",
			verts:  []*ut.ID{&a, &b},
			edges:  []Edge[ut.ID]{a2b},
			edge:   Edge[ut.ID]{nil, &b, wt},
			expect: false,
		},
		{
			desc:   "does not exist (dst)",
			verts:  []*ut.ID{&a, &b},
			edges:  []Edge[ut.ID]{a2b},
			edge:   Edge[ut.ID]{&a, &c, wt},
			expect: false,
		},
		{
			desc:   "does not exist (nil dst)",
			verts:  []*ut.ID{&a, &b},
			edges:  []Edge[ut.ID]{a2b},
			edge:   Edge[ut.ID]{&a, nil, wt},
			expect: false,
		},
		{
			desc:   "does not exist (wt)",
			verts:  []*ut.ID{&a, &b},
			edges:  []Edge[ut.ID]{a2b},
			edge:   Edge[ut.ID]{&a, &b, wt + 1},
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

				ut.AssertEqual(t, tc.expect, g.EdgeExists(tc.edge.Src, tc.edge.Dst, tc.edge.Wt))
			})
		}
	}
}

func TestAddVertex(t *testing.T) {
	a := ut.ID("a")
	b := ut.ID("b")
	c := ut.ID("c")

	cases := []struct {
		desc   string
		verts  []*ut.ID
		expect int
	}{
		{
			desc:   "no vertices",
			verts:  []*ut.ID{},
			expect: 0,
		},
		{
			desc:   "nil vertex",
			verts:  []*ut.ID{nil},
			expect: 0,
		},
		{
			desc:   "unique calls",
			verts:  []*ut.ID{&a, &b, &c},
			expect: 3,
		},
		{
			desc:   "duplicated calls",
			verts:  []*ut.ID{&a, &a, &b, &b, &b},
			expect: 2,
		},
	}

	for _, tc := range cases {
		for _, f := range GraphGenFuncs {
			t.Run(tc.desc, func(t *testing.T) {
				g := f()

				g.AddVertex(tc.verts...)

				ut.AssertEqual(t, tc.expect, g.VertexCount())
			})
		}
	}
}

func TestAddWeightedEdge(t *testing.T) {
	a := ut.ID("a")
	b := ut.ID("b")

	wt := 1.
	a2b := Edge[ut.ID]{&a, &b, wt}

	cases := []struct {
		desc        string
		verts       []*ut.ID
		edges       []Edge[ut.ID]
		edge        Edge[ut.ID]
		expectEdges bool
		expectCount int
	}{
		{
			desc:        "new edge",
			verts:       []*ut.ID{&a, &b},
			edges:       []Edge[ut.ID]{},
			edge:        Edge[ut.ID]{&a, &b, wt},
			expectEdges: true,
			expectCount: 1,
		},
		{
			desc:        "existing edge, same wt",
			verts:       []*ut.ID{&a, &b},
			edges:       []Edge[ut.ID]{a2b},
			edge:        Edge[ut.ID]{&a, &b, wt},
			expectEdges: true,
			expectCount: 1,
		},
		{
			desc:        "existing edge, different wt",
			verts:       []*ut.ID{&a, &b},
			edges:       []Edge[ut.ID]{a2b},
			edge:        Edge[ut.ID]{&a, &b, wt + 1},
			expectEdges: true,
			expectCount: 2,
		},
		{
			desc:        "nil src",
			verts:       []*ut.ID{&a, &b},
			edges:       []Edge[ut.ID]{},
			edge:        Edge[ut.ID]{nil, &b, wt},
			expectEdges: false,
			expectCount: 0,
		},
		{
			desc:        "nil dst",
			verts:       []*ut.ID{&a, &b},
			edges:       []Edge[ut.ID]{},
			edge:        Edge[ut.ID]{&a, nil, wt},
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
					ut.AssertEqual(t, true, g.EdgeExists(tc.edge.Src, tc.edge.Dst, tc.edge.Wt))

					if g.Undirected() {
						ut.AssertEqual(t, true, g.EdgeExists(tc.edge.Dst, tc.edge.Src, tc.edge.Wt))
					}
				}

				ut.AssertEqual(t, tc.expectCount, g.EdgeCount())
			})
		}
	}
}

func TestAddEdge_directed(t *testing.T) {
	a := ut.ID("a")
	b := ut.ID("b")

	src := &a
	dst := &b

	g := NewDirectedGraph[ut.ID]()

	g.AddEdge(src, dst)

	ut.AssertEqual(t, true, g.EdgeExists(src, dst, 0))
}

func TestAddEdge_undirected(t *testing.T) {
	a := ut.ID("a")
	b := ut.ID("b")

	src := &a
	dst := &b

	g := NewUndirectedGraph[ut.ID]()

	g.AddEdge(src, dst)

	ut.AssertEqual(t, true, g.EdgeExists(src, dst, 0))
	ut.AssertEqual(t, true, g.EdgeExists(dst, src, 0))
}
