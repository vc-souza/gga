package viz

import (
	"testing"

	"github.com/vc-souza/gga/algo"
	"github.com/vc-souza/gga/ds"
	ut "github.com/vc-souza/gga/internal/testutils"
)

func TestSCCViz(t *testing.T) {
	g, _, err := ds.Parse(ut.UDGDeps)

	ut.Equal(t, true, err == nil)

	sccs, err := algo.SCCKosaraju(g)

	ut.Equal(t, true, err == nil)

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

	ut.Equal(t, g.VertexCount(), vCount)
	ut.Equal(t, 8, sECount)
	ut.Equal(t, 6, cECount)
}
