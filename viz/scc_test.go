package viz

import (
	"testing"

	"github.com/vc-souza/gga/algo"
	"github.com/vc-souza/gga/ds"
	ut "github.com/vc-souza/gga/internal/testutils"
)

func TestSCCViz(t *testing.T) {
	g, _, err := ds.Parse(ut.UDGDeps)

	ut.AssertNil(t, err)

	sccs, err := algo.SCCKosaraju(g)

	ut.AssertNil(t, err)

	vi := NewSCCViz(g, sccs, nil)

	vCount := 0
	sECount := 0
	cECount := 0

	vi.OnSCCVertex = func(*ds.GV[ds.Text], int) {
		vCount++
	}

	vi.OnSCCEdge = func(*ds.GE[ds.Text], int) {
		sECount++
	}

	vi.OnCrossSCCEdge = func(*ds.GE[ds.Text], int, int) {
		cECount++
	}

	ExportViz[ds.Text](vi, ut.DummyWriter{})

	ut.AssertEqual(t, g.VertexCount(), vCount)
	ut.AssertEqual(t, 8, sECount)
	ut.AssertEqual(t, 6, cECount)
}
