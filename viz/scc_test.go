package viz

import (
	"testing"

	"github.com/vc-souza/gga/algo"
	"github.com/vc-souza/gga/ds"
	ut "github.com/vc-souza/gga/internal/testutils"
)

func TestSCCViz(t *testing.T) {
	g, _, err := ds.NewTextParser().Parse(ut.UDGDeps)

	ut.Equal(t, true, err == nil)

	sccs, err := algo.SCCKosaraju(g)

	ut.Equal(t, true, err == nil)

	vi := NewSCCViz(g, sccs, nil)

	vCount := 0
	sECount := 0
	cECount := 0

	vi.OnSCCVertex = func(v *ds.GV[ds.Text], c int) {
		vCount++
	}

	vi.OnSCCEdge = func(e *ds.GE[ds.Text], c int) {
		sECount++
	}

	vi.OnCrossSCCEdge = func(e *ds.GE[ds.Text], cSrc, cDst int) {
		cECount++
	}

	ExportViz[ds.Text](vi, ut.DummyWriter{})

	ut.Equal(t, g.VertexCount(), vCount)
	ut.Equal(t, 8, sECount)
	ut.Equal(t, 6, cECount)
}
