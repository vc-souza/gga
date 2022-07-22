package viz

import (
	"testing"

	"github.com/vc-souza/gga/algo"
	"github.com/vc-souza/gga/ds"
	ut "github.com/vc-souza/gga/internal/testutils"
)

func TestGSCCViz(t *testing.T) {
	g, _, err := ds.Parse(ut.UDGDeps)

	ut.AssertNil(t, err)

	gscc, _, err := algo.GSCC(g)

	ut.AssertNil(t, err)

	vi := NewGSCCViz(gscc, nil)

	vCount := 0

	vi.OnGSCCVertex = func(*ds.GV[ds.Group[ds.Text]]) {
		vCount++
	}

	ExportViz[ds.Group[ds.Text]](vi, ut.DummyWriter{})

	ut.AssertEqual(t, gscc.VertexCount(), vCount)
}
