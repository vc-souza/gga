package viz

import (
	"testing"

	"github.com/vc-souza/gga/algo"
	"github.com/vc-souza/gga/ds"
	ut "github.com/vc-souza/gga/internal/testutils"
)

type dummyWriter struct{}

func (d dummyWriter) Write(p []byte) (int, error) { return 0, nil }

func TestBFSViz_directed(t *testing.T) {
	g := ds.NewDirectedGraph[ut.ID]()

	v1 := ut.ID("1")
	v2 := ut.ID("2")
	v3 := ut.ID("3")
	v4 := ut.ID("4")
	v5 := ut.ID("5")
	v6 := ut.ID("6")
	src := &v3

	g.AddVertex(&v1)
	g.AddVertex(&v2)
	g.AddVertex(&v3)
	g.AddVertex(&v4)
	g.AddVertex(&v5)
	g.AddVertex(&v6)

	g.AddUnweightedEdge(&v1, &v2)
	g.AddUnweightedEdge(&v1, &v4)
	g.AddUnweightedEdge(&v2, &v5)
	g.AddUnweightedEdge(&v3, &v5)
	g.AddUnweightedEdge(&v3, &v6)
	g.AddUnweightedEdge(&v4, &v2)
	g.AddUnweightedEdge(&v5, &v4)
	g.AddUnweightedEdge(&v6, &v6)

	tree, err := algo.BFS(g, src)

	ut.AssertEqual(t, true, err == nil)

	uCount := 0
	tCount := 0
	sCount := 0
	eCount := 0

	vi := NewBFSViz(g, tree, src)

	vi.OnUnVertex = func(d *ds.GraphVertex[ut.ID], a *algo.BFSNode[ut.ID]) {
		uCount++
	}

	vi.OnTreeVertex = func(d *ds.GraphVertex[ut.ID], a *algo.BFSNode[ut.ID]) {
		tCount++
	}

	vi.OnSourceVertex = func(d *ds.GraphVertex[ut.ID], a *algo.BFSNode[ut.ID]) {
		sCount++
	}

	vi.OnTreeEdge = func(d *ds.GraphEdge[ut.ID]) {
		eCount++
	}

	vi.Export(dummyWriter{})

	ut.AssertEqual(t, 1, uCount)
	ut.AssertEqual(t, 5, tCount)
	ut.AssertEqual(t, 1, sCount)
	ut.AssertEqual(t, 4, eCount)
}

func TestBFSViz_undirected(t *testing.T) {
	g := ds.NewUndirectedGraph[ut.ID]()

	vR := ut.ID("r")
	vS := ut.ID("s")
	vT := ut.ID("t")
	vU := ut.ID("u")
	vV := ut.ID("v")
	vW := ut.ID("w")
	vX := ut.ID("x")
	vY := ut.ID("y")
	src := &vU

	g.AddVertex(&vR)
	g.AddVertex(&vS)
	g.AddVertex(&vT)
	g.AddVertex(&vU)
	g.AddVertex(&vV)
	g.AddVertex(&vW)
	g.AddVertex(&vX)
	g.AddVertex(&vY)

	g.AddUnweightedEdge(&vR, &vS)
	g.AddUnweightedEdge(&vR, &vV)

	g.AddUnweightedEdge(&vS, &vR)
	g.AddUnweightedEdge(&vS, &vW)

	g.AddUnweightedEdge(&vT, &vU)
	g.AddUnweightedEdge(&vT, &vW)
	g.AddUnweightedEdge(&vT, &vX)

	g.AddUnweightedEdge(&vU, &vT)
	g.AddUnweightedEdge(&vU, &vX)
	g.AddUnweightedEdge(&vU, &vY)

	g.AddUnweightedEdge(&vV, &vR)

	g.AddUnweightedEdge(&vW, &vS)
	g.AddUnweightedEdge(&vW, &vT)
	g.AddUnweightedEdge(&vW, &vX)

	g.AddUnweightedEdge(&vX, &vT)
	g.AddUnweightedEdge(&vX, &vU)
	g.AddUnweightedEdge(&vX, &vW)
	g.AddUnweightedEdge(&vX, &vY)

	g.AddUnweightedEdge(&vY, &vU)
	g.AddUnweightedEdge(&vY, &vX)

	tree, err := algo.BFS(g, src)

	ut.AssertEqual(t, true, err == nil)

	uCount := 0
	tCount := 0
	sCount := 0
	eCount := 0

	vi := NewBFSViz(g, tree, src)

	vi.OnUnVertex = func(d *ds.GraphVertex[ut.ID], a *algo.BFSNode[ut.ID]) {
		uCount++
	}

	vi.OnTreeVertex = func(d *ds.GraphVertex[ut.ID], a *algo.BFSNode[ut.ID]) {
		tCount++
	}

	vi.OnSourceVertex = func(d *ds.GraphVertex[ut.ID], a *algo.BFSNode[ut.ID]) {
		sCount++
	}

	vi.OnTreeEdge = func(d *ds.GraphEdge[ut.ID]) {
		eCount++
	}

	vi.Export(dummyWriter{})

	ut.AssertEqual(t, 0, uCount)
	ut.AssertEqual(t, 8, tCount)
	ut.AssertEqual(t, 1, sCount)
	ut.AssertEqual(t, 14, eCount)
}
