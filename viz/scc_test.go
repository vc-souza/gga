package viz

import (
	"testing"

	"github.com/vc-souza/gga/algo"
	"github.com/vc-souza/gga/ds"
	ut "github.com/vc-souza/gga/internal/testutils"
)

func TestSCCViz(t *testing.T) {
	g, _, err := ds.NewTextParser().Parse(ut.UDGDeps)

	ut.AssertEqual(t, true, err == nil)

	sccs, err := algo.SCCKosaraju(g)

	ut.AssertEqual(t, true, err == nil)

	vi := NewSCCViz(g, sccs, nil)

	vCount := 0
	sECount := 0
	cECount := 0

	vi.OnSCCVertex = func(v *ds.GraphVertex[ds.Text], c int) {
		vCount++
	}

	vi.OnSCCEdge = func(e *ds.GraphEdge[ds.Text], c int) {
		sECount++
	}

	vi.OnCrossSCCEdge = func(e *ds.GraphEdge[ds.Text], cSrc, cDst int) {
		cECount++
	}

	ExportViz[ds.Text](vi, ut.DummyWriter{})

	ut.AssertEqual(t, g.VertexCount(), vCount)
	ut.AssertEqual(t, 8, sECount)
	ut.AssertEqual(t, 6, cECount)
}
