package viz

import (
	"testing"

	"github.com/vc-souza/gga/algo"
	"github.com/vc-souza/gga/ds"
	ut "github.com/vc-souza/gga/internal/testutils"
)

func TestMSTViz(t *testing.T) {
	g, _, err := ds.Parse(ut.WUGSimple)

	ut.AssertNil(t, err)

	mst, err := algo.MSTKruskal(g)

	ut.AssertNil(t, err)

	vi := NewMSTViz(g, mst, nil)

	eCount := 0

	vi.OnMSTEdge = func(*ds.GE[ds.Text]) {
		eCount++
	}

	ExportViz[ds.Text](vi, ut.DummyWriter{})

	ut.AssertEqual(t, 2*(g.VertexCount()-1), eCount)
}
