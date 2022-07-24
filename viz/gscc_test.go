package viz

import (
	"testing"

	"github.com/vc-souza/gga/algo"
	"github.com/vc-souza/gga/ds"
	ut "github.com/vc-souza/gga/internal/testutils"
)

func TestGSCCViz(t *testing.T) {
	g, _, err := ds.Parse(ut.UDGDeps)

	ut.Nil(t, err)

	gscc, _, err := algo.GSCC(g)

	ut.Nil(t, err)

	vi := NewGSCCViz(gscc, nil)

	vCount := 0

	vi.OnGSCCVertex = func(int) {
		vCount++
	}

	ExportViz(vi, ut.DummyWriter{})

	ut.Equal(t, gscc.VertexCount(), vCount)
}
