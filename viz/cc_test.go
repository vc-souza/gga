package viz

import (
	"testing"

	"github.com/vc-souza/gga/algo"
	"github.com/vc-souza/gga/ds"
	ut "github.com/vc-souza/gga/internal/testutils"
)

func TestCCViz(t *testing.T) {
	g, _, err := ds.Parse(ut.UUGDisc)

	ut.Nil(t, err)

	ccs, err := algo.CCDFS(g)

	ut.Nil(t, err)

	vi := NewCCViz(g, ccs, nil)

	vCount := 0
	eCount := 0

	vi.OnCCVertex = func(int, int) {
		vCount++
	}

	vi.OnCCEdge = func(int, int, int) {
		eCount++
	}

	ExportViz(vi, ut.DummyWriter{})

	ut.Equal(t, g.VertexCount(), vCount)
	ut.Equal(t, g.EdgeCount(), eCount)
}
