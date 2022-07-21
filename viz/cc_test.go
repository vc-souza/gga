package viz

import (
	"testing"

	"github.com/vc-souza/gga/algo"
	"github.com/vc-souza/gga/ds"
	ut "github.com/vc-souza/gga/internal/testutils"
)

func TestCCViz(t *testing.T) {
	g, _, err := ds.Parse(ut.UUGDisc)

	ut.Equal(t, true, err == nil)

	ccs, err := algo.CCDFS(g)

	ut.Equal(t, true, err == nil)

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

	ut.Equal(t, g.VertexCount(), vCount)
	ut.Equal(t, g.EdgeCount(), eCount)
}
