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

				tu.AssertEqual(t, tc.expect, len(g.Adj))
			})
		}
	}
}

// func TestAddWeightedEdge(t *testing.T) {
// 	a := tu.ID("a")
// 	b := tu.ID("b")
// 	c := tu.ID("c")

// 	cases := []struct {
// 		desc   string
// 		verts  []*tu.ID
// 		edges  []Edge[tu.ID]
// 		expect int
// 	}{
// 		{
// 			desc:   "???",
// 			verts:  []*tu.ID{},
// 			edges:  []Edge[tu.ID]{},
// 			expect: 0,
// 		},
// 	}
// }
