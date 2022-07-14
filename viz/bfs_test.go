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
	g, vars, err := ds.NewTextParser().Parse(ut.BasicUDG)

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
	g, vars, err := ds.NewTextParser().Parse(ut.BasicUUG)

	ut.AssertEqual(t, true, err == nil)

	src := vars["u"]

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
	ut.AssertEqual(t, 7, eCount)
}
