package viz

import (
	"testing"

	"github.com/vc-souza/gga/algo"
	"github.com/vc-souza/gga/ds"
	ut "github.com/vc-souza/gga/internal/testutils"
)

func TestTSortViz(t *testing.T) {
	g, _, err := ds.NewTextParser().Parse(ut.UDGDress)

	ut.Equal(t, true, err == nil)

	ord, err := algo.TSort(g)

	ut.Equal(t, true, err == nil)

	vi := NewTSortViz(g, ord, nil)

	eNotExistsCount := 0
	eExitsCount := 0
	eCount := 0
	vCount := 0

	vi.OnVertexRank = func(*ds.GV[ds.Text], int) {
		vCount++
	}

	vi.OnOrderEdge = func(d *ds.GE[ds.Text], b bool) {
		if b {
			eExitsCount++
		} else {
			eNotExistsCount++
		}

		eCount++
	}

	ExportViz[ds.Text](vi, ut.DummyWriter{})

	ut.Equal(t, g.VertexCount(), vCount)
	ut.Equal(t, g.VertexCount()-1, eCount)
	ut.Equal(t, 4, eExitsCount)
	ut.Equal(t, 4, eNotExistsCount)
}
