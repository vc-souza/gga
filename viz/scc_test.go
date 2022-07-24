package viz

import (
	"testing"

	"github.com/vc-souza/gga/algo"
	"github.com/vc-souza/gga/ds"
	ut "github.com/vc-souza/gga/internal/testutils"
)

func TestSCCViz(t *testing.T) {
	g, _, err := ds.Parse(ut.UDGDeps)

	ut.Nil(t, err)

	sccs, err := algo.SCCTarjan(g)

	ut.Nil(t, err)

	vi := NewSCCViz(g, sccs, nil)

	vCount := 0
	sECount := 0
	cECount := 0

	vi.OnSCCVertex = func(int, int) {
		vCount++
	}

	vi.OnSCCEdge = func(int, int, int) {
		sECount++
	}

	vi.OnCrossSCCEdge = func(int, int, int, int) {
		cECount++
	}

	ExportViz(vi, ut.DummyWriter{})

	ut.Equal(t, g.VertexCount(), vCount)
	ut.Equal(t, 8, sECount)
	ut.Equal(t, 6, cECount)
}
