package viz

import (
	"testing"

	"github.com/vc-souza/gga/algo"
	"github.com/vc-souza/gga/ds"
	ut "github.com/vc-souza/gga/internal/testutils"
)

func TestMSTViz(t *testing.T) {
	g, _, err := ds.Parse(ut.WUGSimple)

	ut.Nil(t, err)

	mst, err := algo.MSTPrim(g)

	ut.Nil(t, err)

	vi := NewMSTViz(g, mst, nil)

	eCount := 0

	vi.OnMSTEdge = func(int, int) {
		eCount++
	}

	ExportViz(vi, ut.DummyWriter{})

	ut.Equal(t, 2*(g.VertexCount()-1), eCount)
}
