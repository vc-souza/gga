package viz

import (
	"testing"

	"github.com/vc-souza/gga/algo"
	"github.com/vc-souza/gga/ds"
	ut "github.com/vc-souza/gga/internal/testutils"
)

func TestCCViz(t *testing.T) {
	g, _, err := ds.Parse(ut.UUGDisc)

	ut.AssertNil(t, err)

	ccs, err := algo.CCDFS(g)

	ut.AssertNil(t, err)

	vi := NewCCViz(g, ccs, nil)

	vCount := 0
	eCount := 0

	vi.OnCCVertex = func(*ds.GV[ds.Text], int) {
		vCount++
	}

	vi.OnCCEdge = func(*ds.GE[ds.Text], int) {
		eCount++
	}

	ExportViz[ds.Text](vi, ut.DummyWriter{})

	ut.AssertEqual(t, g.VertexCount(), vCount)
	ut.AssertEqual(t, g.EdgeCount(), eCount)
}
