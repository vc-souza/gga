package ds

import (
	"testing"

	ut "github.com/vc-souza/gga/internal/testutils"
)

type GraphGen func() *Graph[ut.ID]

var GraphGenFuncs = map[string]GraphGen{
	"graph":   NewUndirectedGraph[ut.ID],
	"digraph": NewDirectedGraph[ut.ID],
}

func tag(gtype, desc string) string {
	return gtype + " " + desc
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
	a := ut.ID("a")
	b := ut.ID("b")

	a2b := GraphEdge[ut.ID]{Src: &a, Dst: &b}

	cases := []struct {
		desc   string
		edges  []GraphEdge[ut.ID]
		expect int
	}{
		{
			desc:   "zero edges",
			edges:  []GraphEdge[ut.ID]{},
			expect: 0,
		},
		{
			desc:   "one edge",
			edges:  []GraphEdge[ut.ID]{a2b},
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
	a := ut.ID("a")
	b := ut.ID("b")
	c := ut.ID("c")

	a2b := GraphEdge[ut.ID]{Src: &a, Dst: &b}

	cases := []struct {
		desc   string
		verts  []*ut.ID
		edges  []GraphEdge[ut.ID]
		edge   GraphEdge[ut.ID]
		expect bool
	}{
		{
			desc:   "exists",
			verts:  []*ut.ID{&a, &b},
			edges:  []GraphEdge[ut.ID]{a2b},
			edge:   GraphEdge[ut.ID]{Src: &a, Dst: &b},
			expect: true,
		},
		{
			desc:   "does not exist (src)",
			verts:  []*ut.ID{&a, &b},
			edges:  []GraphEdge[ut.ID]{a2b},
			edge:   GraphEdge[ut.ID]{Src: &c, Dst: &b},
			expect: false,
		},
		{
			desc:   "does not exist (nil src)",
			verts:  []*ut.ID{&a, &b},
			edges:  []GraphEdge[ut.ID]{a2b},
			edge:   GraphEdge[ut.ID]{Src: nil, Dst: &b},
			expect: false,
		},
		{
			desc:   "does not exist (dst)",
			verts:  []*ut.ID{&a, &b},
			edges:  []GraphEdge[ut.ID]{a2b},
			edge:   GraphEdge[ut.ID]{Src: &a, Dst: &c},
			expect: false,
		},
		{
			desc:   "does not exist (nil dst)",
			verts:  []*ut.ID{&a, &b},
			edges:  []GraphEdge[ut.ID]{a2b},
			edge:   GraphEdge[ut.ID]{Src: &a, Dst: nil},
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
	a := ut.ID("a")

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
			verts:  []*ut.ID{},
			vert:   &a,
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
					ut.AssertEqual(t, tc.vert, vert.Satellite)
				}
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
	a := ut.ID("a")
	b := ut.ID("b")

	a2b := GraphEdge[ut.ID]{Src: &a, Dst: &b}

	cases := []struct {
		desc        string
		verts       []*ut.ID
		edges       []GraphEdge[ut.ID]
		edge        GraphEdge[ut.ID]
		expectEdges bool
		expectCount int
	}{
		{
			desc:        "new edge",
			verts:       []*ut.ID{&a, &b},
			edges:       []GraphEdge[ut.ID]{},
			edge:        GraphEdge[ut.ID]{Src: &a, Dst: &b},
			expectEdges: true,
			expectCount: 1,
		},
		{
			desc:        "existing edge",
			verts:       []*ut.ID{&a, &b},
			edges:       []GraphEdge[ut.ID]{a2b},
			edge:        GraphEdge[ut.ID]{Src: &a, Dst: &b},
			expectEdges: true,
			expectCount: 1,
		},
		{
			desc:        "nil src",
			verts:       []*ut.ID{&a, &b},
			edges:       []GraphEdge[ut.ID]{},
			edge:        GraphEdge[ut.ID]{Src: nil, Dst: &b},
			expectEdges: false,
			expectCount: 0,
		},
		{
			desc:        "nil dst",
			verts:       []*ut.ID{&a, &b},
			edges:       []GraphEdge[ut.ID]{},
			edge:        GraphEdge[ut.ID]{Src: &a, Dst: nil},
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
	a := ut.ID("a")
	b := ut.ID("b")

	src := &a
	dst := &b

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
