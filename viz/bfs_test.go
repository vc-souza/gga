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
	g, vars, err := new(ds.TextParser).Parse(`
	digraph
	1#2,4
	2#5
	3#5,6
	4#2
	5#4
	6#6
	`)

	ut.AssertEqual(t, true, err == nil)

	src := vars["3"]

	tree, err := algo.BFS(g, src)

	ut.AssertEqual(t, true, err == nil)

	uCount := 0
	tCount := 0
	sCount := 0
	eCount := 0

	vi := NewBFSViz(g, tree, src)

	vi.OnUnVertex = func(d *ds.GraphVertex[ds.Text], a *algo.BFSNode[ds.Text]) {
		uCount++
	}

	vi.OnTreeVertex = func(d *ds.GraphVertex[ds.Text], a *algo.BFSNode[ds.Text]) {
		tCount++
	}

	vi.OnSourceVertex = func(d *ds.GraphVertex[ds.Text], a *algo.BFSNode[ds.Text]) {
		sCount++
	}

	vi.OnTreeEdge = func(d *ds.GraphEdge[ds.Text]) {
		eCount++
	}

	vi.Export(dummyWriter{})

	ut.AssertEqual(t, 1, uCount)
	ut.AssertEqual(t, 5, tCount)
	ut.AssertEqual(t, 1, sCount)
	ut.AssertEqual(t, 4, eCount)
}

func TestBFSViz_undirected(t *testing.T) {
	g := ds.NewUndirectedGraph[ds.Text]()

	vR := ds.Text("r")
	vS := ds.Text("s")
	vT := ds.Text("t")
	vU := ds.Text("u")
	vV := ds.Text("v")
	vW := ds.Text("w")
	vX := ds.Text("x")
	vY := ds.Text("y")
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

	vi.OnUnVertex = func(d *ds.GraphVertex[ds.Text], a *algo.BFSNode[ds.Text]) {
		uCount++
	}

	vi.OnTreeVertex = func(d *ds.GraphVertex[ds.Text], a *algo.BFSNode[ds.Text]) {
		tCount++
	}

	vi.OnSourceVertex = func(d *ds.GraphVertex[ds.Text], a *algo.BFSNode[ds.Text]) {
		sCount++
	}

	vi.OnTreeEdge = func(d *ds.GraphEdge[ds.Text]) {
		eCount++
	}

	vi.Export(dummyWriter{})

	ut.AssertEqual(t, 0, uCount)
	ut.AssertEqual(t, 8, tCount)
	ut.AssertEqual(t, 1, sCount)
	ut.AssertEqual(t, 14, eCount)
}
